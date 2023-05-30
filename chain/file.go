package chain

import (
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/CESSProject/sdk-go/core/erasure"
	"github.com/CESSProject/sdk-go/core/hashtree"
	"github.com/CESSProject/sdk-go/core/pattern"
	"github.com/CESSProject/sdk-go/core/rule"
	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

func (c *ChainSDK) GetFile(roothash, dir string) (string, error) {
	var (
		segmentspath = make([]string, 0)
	)
	userfile := filepath.Join(dir, roothash)
	_, err := os.Stat(userfile)
	if err == nil {
		return userfile, nil
	}
	os.MkdirAll(dir, rule.DirMode)
	f, err := os.Create(userfile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fmeta, err := c.QueryFileMetadata(roothash)
	if err != nil {
		return "", err
	}

	defer func(basedir string) {
		for _, segment := range fmeta.SegmentList {
			os.Remove(filepath.Join(basedir, string(segment.Hash[:])))
			for _, fragment := range segment.FragmentList {
				os.Remove(filepath.Join(basedir, string(fragment.Hash[:])))
			}
		}
	}(dir)

	for _, segment := range fmeta.SegmentList {
		fragmentpaths := make([]string, 0)
		for _, fragment := range segment.FragmentList {
			//miner, err := c.QueryStorageMiner(fragment.Miner[:])
			if err != nil {
				return "", err
			}
			// peerid, err := c.AddMultiaddrToPearstore(string(miner.PeerId[:]), time.Hour)
			// if err != nil {
			// 	return "", err
			// }
			fragmentpath := filepath.Join(dir, string(fragment.Hash[:]))
			// err = c.Protocol.ReadFileAction(peerid, roothash, string(fragment.Hash[:]), fragmentpath, rule.FragmentSize)
			// if err != nil {
			// 	continue
			// }
			fragmentpaths = append(fragmentpaths, fragmentpath)
			segmentpath := filepath.Join(dir, string(segment.Hash[:]))
			if len(fragmentpaths) >= rule.DataShards {
				err = erasure.ReedSolomon_Restore(segmentpath, fragmentpaths)
				if err != nil {
					return "", err
				}
				segmentspath = append(segmentspath, segmentpath)
				break
			}
		}
	}

	if len(segmentspath) != len(fmeta.SegmentList) {
		return "", fmt.Errorf("Download failed")
	}
	var writecount = 0
	for i := 0; i < len(fmeta.SegmentList); i++ {
		for j := 0; j < len(segmentspath); j++ {
			if string(fmeta.SegmentList[i].Hash[:]) == filepath.Base(segmentspath[j]) {
				buf, err := os.ReadFile(segmentspath[j])
				if err != nil {
					return "", err
				}
				f.Write(buf)
				writecount++
				break
			}
		}
	}
	if writecount != len(fmeta.SegmentList) {
		return "", fmt.Errorf("Write failed")
	}
	return userfile, nil
}

func (c *ChainSDK) PutFile(owner []byte, segmentInfo []pattern.SegmentDataInfo, roothash, filename, bucketname string) (uint8, error) {
	var err error
	var storageOrder pattern.StorageOrder

	_, err = c.QueryFileMetadata(roothash)
	if err == nil {
		return 0, nil
	}

	for i := 0; i < 3; i++ {
		storageOrder, err = c.QueryStorageOrder(roothash)
		if err != nil {
			if err.Error() == pattern.ERR_Empty {
				err = c.GenerateStorageOrder(roothash, segmentInfo, owner, filename, bucketname)
				if err != nil {
					return 0, err
				}
			}
			time.Sleep(rule.BlockInterval)
			continue
		}
		break
	}
	if err != nil {
		return 0, err
	}

	// store fragment to storage
	err = c.StorageData(roothash, segmentInfo, storageOrder.AssignedMiner)
	if err != nil {
		return 0, err
	}
	return uint8(storageOrder.Count), nil
}

func (c *ChainSDK) ProcessingData(path string) ([]pattern.SegmentDataInfo, string, error) {
	var (
		err          error
		f            *os.File
		fstat        fs.FileInfo
		segmentCount int64
		num          int
	)

	fstat, err = os.Stat(path)
	if err != nil {
		return nil, "", err
	}
	if fstat.IsDir() {
		return nil, "", errors.New("Not a file")
	}

	baseDir := filepath.Dir(path)
	segmentCount = fstat.Size() / rule.SegmentSize
	if fstat.Size()%int64(rule.SegmentSize) != 0 {
		segmentCount++
	}

	segment := make([]pattern.SegmentDataInfo, segmentCount)

	buf := make([]byte, rule.SegmentSize)

	f, err = os.Open(path)
	if err != nil {
		return segment, "", err
	}
	defer f.Close()

	for i := int64(0); i < segmentCount; i++ {
		f.Seek(rule.SegmentSize*i, 0)
		num, err = f.Read(buf)
		if err != nil && err != io.EOF {
			return segment, "", err
		}
		if num == 0 {
			break
		}
		if num < rule.SegmentSize {
			if i+1 != segmentCount {
				return segment, "", fmt.Errorf("Error reading %s", path)
			}
			remainbuf := make([]byte, rule.SegmentSize-num)
			copy(buf[num:], remainbuf)
		}

		hash, err := utils.CalcSHA256(buf)
		if err != nil {
			return segment, "", err
		}

		segmentPath := filepath.Join(baseDir, hash)
		_, err = os.Stat(segmentPath)
		if err != nil {
			fsegment, err := os.Create(segmentPath)
			if err != nil {
				return segment, "", err
			}
			_, err = fsegment.Write(buf)
			if err != nil {
				fsegment.Close()
				return segment, "", err
			}
			err = fsegment.Sync()
			if err != nil {
				fsegment.Close()
				return segment, "", err
			}
			fsegment.Close()
		}

		segment[i].SegmentHash = segmentPath
		segment[i].FragmentHash, err = erasure.ReedSolomon(segmentPath)
		if err != nil {
			return segment, "", err
		}
	}

	segmenthash := ExtractSegmenthash(segment)

	// Calculate merkle hash tree
	hTree, err := hashtree.NewHashTree(segmenthash)
	if err != nil {
		return segment, "", err
	}

	return segment, hex.EncodeToString(hTree.MerkleRoot()), err
}

func (c *ChainSDK) GenerateStorageOrder(roothash string, segment []pattern.SegmentDataInfo, owner []byte, filename, buckname string) error {
	var err error
	var segmentList = make([]pattern.SegmentList, len(segment))
	var user pattern.UserBrief
	for i := 0; i < len(segment); i++ {
		hash := filepath.Base(segment[i].SegmentHash)
		for k := 0; k < len(hash); k++ {
			segmentList[i].SegmentHash[k] = types.U8(hash[k])
		}
		segmentList[i].FragmentHash = make([]pattern.FileHash, len(segment[i].FragmentHash))
		for j := 0; j < len(segment[i].FragmentHash); j++ {
			hash := filepath.Base(segment[i].FragmentHash[j])
			for k := 0; k < len(hash); k++ {
				segmentList[i].FragmentHash[j][k] = types.U8(hash[k])
			}
		}
	}
	acc, err := types.NewAccountID(owner)
	if err != nil {
		return err
	}
	user.User = *acc
	user.BucketName = types.NewBytes([]byte(buckname))
	user.FileName = types.NewBytes([]byte(filename))
	_, err = c.UploadDeclaration(roothash, segmentList, user)
	return err
}

func ExtractSegmenthash(segment []pattern.SegmentDataInfo) []string {
	var segmenthash = make([]string, len(segment))
	for i := 0; i < len(segment); i++ {
		segmenthash[i] = segment[i].SegmentHash
	}
	return segmenthash
}

func (c *ChainSDK) StorageData(roothash string, segment []pattern.SegmentDataInfo, minerTaskList []pattern.MinerTaskList) error {
	// var err error

	// query all assigned miner multiaddr
	// multiaddrs, err := c.QueryAssignedMiner(minerTaskList)
	// if err != nil {
	// 	return err
	// }

	// basedir := filepath.Dir(segment[0].FragmentHash[0])
	// for i := 0; i < len(multiaddrs); i++ {
	// 	peerid, err := c.Protocol.AddMultiaddrToPearstore(multiaddrs[i], 0)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	for j := 0; j < len(minerTaskList[i].Hash); j++ {
	// 		err = c.Protocol.WriteFileAction(peerid, roothash, filepath.Join(basedir, string(minerTaskList[i].Hash[j][:])))
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	return nil
}

func (c *ChainSDK) QueryAssignedMiner(minerTaskList []pattern.MinerTaskList) ([]string, error) {
	var multiaddrs = make([]string, len(minerTaskList))
	for i := 0; i < len(minerTaskList); i++ {
		minerInfo, err := c.QueryStorageMiner(minerTaskList[i].Account[:])
		if err != nil {
			return multiaddrs, err
		}
		multiaddrs[i] = string(minerInfo.PeerId[:])
	}
	return multiaddrs, nil
}
