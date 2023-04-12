package erasure

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/CESSProject/sdk-go/core/rule"
	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/klauspost/reedsolomon"
)

// ReedSolomon uses reed-solomon algorithm to redundancy files
// Return:
//
//  1. All file blocks (sorted sequentially)
//  2. Error message
func ReedSolomon(path string) ([]string, error) {
	var shardspath = make([]string, 0)
	fstat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if fstat.IsDir() {
		return nil, errors.New("not a file")
	}
	if fstat.Size() != rule.SegmentSize {
		return nil, errors.New("invalid size")
	}

	datashards, parshards := rule.DataShards, rule.ParShards
	basedir := filepath.Dir(path)

	enc, err := reedsolomon.New(datashards, parshards)
	if err != nil {
		return shardspath, err
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return shardspath, err
	}

	// Split the file into equally sized shards.
	shards, err := enc.Split(b)
	if err != nil {
		return shardspath, err
	}
	// Encode parity
	err = enc.Encode(shards)
	if err != nil {
		return shardspath, err
	}
	// Write out the resulting files.
	for _, shard := range shards {
		hash, err := utils.CalcSHA256(shard)
		if err != nil {
			return shardspath, err
		}
		newpath := filepath.Join(basedir, hash)
		_, err = os.Stat(newpath)
		if err != nil {
			err = ioutil.WriteFile(newpath, shard, os.ModePerm)
			if err != nil {
				return shardspath, err
			}
		}
		shardspath = append(shardspath, newpath)
	}
	return shardspath, nil

	// Create encoding matrix.
	// enc, err := reedsolomon.NewStream(datashards, parshards)
	// if err != nil {
	// 	return shardspath, err
	// }

	// f, err := os.Open(path)
	// if err != nil {
	// 	return shardspath, err
	// }

	// instat, err := f.Stat()
	// if err != nil {
	// 	return shardspath, err
	// }

	// shards := datashards + parshards
	// out := make([]*os.File, shards)

	// // Create the resulting files.
	// dir, file := filepath.Split(path)

	// for i := range out {
	// 	var outfn string
	// 	if i < 10 {
	// 		outfn = fmt.Sprintf("%s.00%d", file, i)
	// 	} else {
	// 		outfn = fmt.Sprintf("%s.0%d", file, i)
	// 	}
	// 	out[i], err = os.Create(filepath.Join(dir, outfn))
	// 	if err != nil {
	// 		return shardspath, err
	// 	}
	// 	//out[i].Close()
	// 	shardspath = append(shardspath, filepath.Join(dir, outfn))
	// }

	// // Split into files.
	// data := make([]io.Writer, datashards)
	// for i := range data {
	// 	data[i] = out[i]
	// }
	// // Do the split
	// err = enc.Split(f, data, instat.Size())
	// if err != nil {
	// 	return shardspath, err
	// }

	// // Close and re-open the files.
	// input := make([]io.Reader, datashards)

	// for i := range data {
	// 	f, err := os.Open(out[i].Name())
	// 	if err != nil {
	// 		return shardspath, err
	// 	}
	// 	input[i] = f
	// 	defer f.Close()
	// }

	// // Create parity output writers
	// parity := make([]io.Writer, parshards)
	// for i := range parity {
	// 	parity[i] = out[datashards+i]
	// 	defer out[datashards+i].Close()
	// }

	// // Encode parity
	// err = enc.Encode(input, parity)
	// if err != nil {
	// 	return shardspath, err
	// }

	// return shardspath, nil
}

func ReedSolomon_Restore(outpath string, shardspath []string) error {
	_, err := os.Stat(outpath)
	if err == nil {
		return nil
	}

	datashards, parshards := rule.DataShards, rule.ParShards

	enc, err := reedsolomon.New(datashards, parshards)
	if err != nil {
		return err
	}
	shards := make([][]byte, datashards+parshards)
	for k, v := range shardspath {
		//infn := fmt.Sprintf("%s.00%d", outfn, i)
		shards[k], err = ioutil.ReadFile(v)
		if err != nil {
			shards[k] = nil
		}
	}

	// Verify the shards
	ok, _ := enc.Verify(shards)
	if !ok {
		err = enc.Reconstruct(shards)
		if err != nil {
			return err
		}
		ok, err = enc.Verify(shards)
		if !ok {
			return err
		}
	}
	f, err := os.Create(outpath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = enc.Join(f, shards, len(shards[0])*datashards)
	return err

}

func openInput(dataShards, parShards int, fname string) (r []io.Reader, size int64, err error) {
	shards := make([]io.Reader, dataShards+parShards)
	for i := range shards {
		var infn string
		if i < 10 {
			infn = fmt.Sprintf("%s.00%d", fname, i)
		} else {
			infn = fmt.Sprintf("%s.0%d", fname, i)
		}
		f, err := os.Open(infn)
		if err != nil {
			shards[i] = nil
			continue
		} else {
			shards[i] = f
		}
		stat, err := f.Stat()
		if err != nil {
			return nil, 0, err
		}
		if stat.Size() > 0 {
			size = stat.Size()
		} else {
			shards[i] = nil
		}
	}
	return shards, size, nil
}
