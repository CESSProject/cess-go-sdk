package client

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/CESSProject/sdk-go/core/erasure"
	"github.com/CESSProject/sdk-go/core/rule"
	"github.com/CESSProject/sdk-go/core/utils"
)

type SegmentInfo struct {
	SegmentHash  string
	FragmentHash []string
}

func (c *Cli) PutFile(owner []byte, path, filename, bucketname string) error {
	var err error
	var ok bool
	// var roothash string
	// var segment = make([]SegmentInfo, 0)

	//
	if ok = utils.CheckBucketName(bucketname); !ok {
		return errors.New("Invalid bucketname")
	}

	//
	ok, err = c.Chain.IsGrantor(owner)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("Unauthorized")
	}

	//
	c.ProcessingData(path)
	return err
}

func (c *Cli) ProcessingData(path string) ([]SegmentInfo, error) {
	var (
		err          error
		f            *os.File
		fstat        fs.FileInfo
		segmentCount int64
		num          int
	)

	fstat, err = os.Stat(path)
	if err != nil {
		return nil, err
	}
	if fstat.IsDir() {
		return nil, errors.New("Not a file")
	}

	segmentCount = fstat.Size() / rule.SegmentSize
	if fstat.Size()%int64(rule.SegmentSize) != 0 {
		segmentCount++
	}

	segment := make([]SegmentInfo, segmentCount)

	buf := make([]byte, rule.SegmentSize)

	f, err = os.Open(path)
	if err != nil {
		return segment, err
	}
	defer f.Close()

	for i := int64(0); i < segmentCount; i++ {
		f.Seek(rule.SegmentSize*i, 0)
		num, err = f.Read(buf)
		if err != nil && err != io.EOF {
			return segment, err
		}
		if num == 0 {
			break
		}
		if num < rule.SegmentSize {
			if i+1 != segmentCount {
				return segment, fmt.Errorf("Error reading %s", path)
			}
			utils.CalcSHA256(buf)
		}
		hash, err := utils.CalcSHA256(buf)
		if err != nil {
			return segment, err
		}

		segmentPath := filepath.Join(c.Workspace(), rule.TempDir, hash)
		_, err = os.Stat(segmentPath)
		if err != nil {
			fsegment, err := os.Create(segmentPath)
			if err != nil {
				return segment, err
			}
			_, err = fsegment.Write(buf[:num])
			if err != nil {
				fsegment.Close()
				return segment, err
			}
			err = fsegment.Sync()
			if err != nil {
				fsegment.Close()
				return segment, err
			}
			fsegment.Close()
		}

		segment[i].SegmentHash = segmentPath
		segment[i].FragmentHash, err = erasure.ReedSolomon(segmentPath)
		if err != nil {
			return segment, err
		}
	}

	return segment, err
}
