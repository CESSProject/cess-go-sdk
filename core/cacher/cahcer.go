package cacher

import (
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/muesli/cache2go"
	"github.com/pkg/errors"
)

const (
	DEFAULT_EXPIRATION      = 60 * time.Minute
	DEFAULT_MAX_CACHE_SPACE = 512 * 1024 * 1024 * 1024
	DEFAULT_CACHE_NAME      = "file_cache"
)

type FileCache interface {
	MoveFileToCache(fname, fpath string) error
	SaveDataToCache(fname string, data []byte) error
	AddCacheRecord(fname, fpath string) error
	GetCacheRecord(fname string) (string, error)
	RemoveCacheRecord(fname string) error
}

type Cacher struct {
	lock       *sync.RWMutex
	cacher     *cache2go.CacheTable
	cacheSpace int64
	usedSpace  int64
	exp        time.Duration
	CacheDir   string
}

func NewCacher(exp time.Duration, maxSpace int64, cacheDir string) FileCache {
	if exp <= 0 {
		exp = DEFAULT_EXPIRATION
	}
	if maxSpace <= 0 {
		maxSpace = DEFAULT_MAX_CACHE_SPACE
	}
	cacher := &Cacher{
		exp:        exp,
		lock:       &sync.RWMutex{},
		cacher:     cache2go.Cache(DEFAULT_CACHE_NAME),
		CacheDir:   cacheDir,
		cacheSpace: maxSpace,
	}
	cacher.cacher.SetAboutToDeleteItemCallback(func(ci *cache2go.CacheItem) {
		fpath, ok := ci.Data().(string)
		if !ok {
			return
		}
		cacher.removeFile(fpath)
	})

	// restore cache record
	filepath.WalkDir(cacheDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		if info.Size() <= 0 || info.IsDir() {
			return nil
		}
		cacher.cacher.Add(d.Name(), cacher.exp, path)
		return nil
	})
	return cacher
}

func (c *Cacher) MoveFileToCache(fname, fpath string) error {
	f, err := os.Stat(fpath)
	if err != nil {
		return errors.Wrap(err, "move file to cache error")
	}
	if f.IsDir() {
		return errors.Wrap(errors.New("not a file"), "move file to cache error")
	}
	cpath := path.Join(c.CacheDir, fname)
	f2, err := os.Stat(cpath)
	size := f.Size()
	if err == nil {
		if f2.Size() == size {
			c.cacher.Add(fname, c.exp, cpath)
			return nil
		}
		size -= f2.Size()
	}

	input, err := os.Open(fpath)
	if err != nil {
		return errors.Wrap(err, "move file to cache error")
	}
	defer input.Close()

	output, err := os.Create(cpath)
	if err != nil {
		return errors.Wrap(err, "move file to cache error")
	}
	defer output.Close()

	c.lock.Lock()
	defer c.lock.Unlock()
	free, err := utils.GetDirFreeSpace(cpath)
	if err != nil {
		return errors.Wrap(err, "move file to cache error")
	}
	if c.usedSpace+size > c.cacheSpace || int64(free) < size {
		return errors.Wrap(errors.New("not enough cache space"), "move file to cache error")
	}

	_, err = io.Copy(output, input)
	if err != nil {
		return errors.Wrap(err, "move file to cache error")
	}
	os.Remove(fpath)
	c.cacher.Add(fname, c.exp, cpath)
	c.usedSpace += size
	return nil
}

func (c *Cacher) SaveDataToCache(fname string, data []byte) error {
	cpath := path.Join(c.CacheDir, fname)
	f, err := os.Stat(cpath)
	size := f.Size()
	if err == nil {
		if size == int64(len(data)) {
			c.cacher.Add(fname, c.exp, cpath)
			return nil
		}
		size = int64(len(data)) - size
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	free, err := utils.GetDirFreeSpace(cpath)
	if err != nil {
		return errors.Wrap(err, "save file to cache error")
	}
	if c.usedSpace+size > c.cacheSpace || int64(free) < size {
		return errors.Wrap(errors.New("not enough cache space"), "save file to cache error")
	}
	file, err := os.Create(cpath)
	if err != nil {
		return errors.Wrap(err, "save file to cache error")
	}
	_, err = file.Write(data)
	if err != nil {
		return errors.Wrap(err, "save file to cache error")
	}
	c.cacher.Add(fname, c.exp, cpath)
	c.usedSpace += size
	return nil
}

func (c *Cacher) AddCacheRecord(fname, cpath string) error {
	if cpath == "" {
		cpath = path.Join(c.CacheDir, fname)
	}
	f, err := os.Stat(cpath)
	if err != nil {
		return errors.Wrap(err, "add cache record error")
	}
	if f.Size() <= 0 {
		return errors.Wrap(errors.New("invalid file"), "add cache record error")
	}
	c.cacher.Add(fname, c.exp, cpath)
	return nil
}

func (c *Cacher) GetCacheRecord(fname string) (string, error) {
	value, err := c.cacher.Value(fname)
	if err != nil {
		return "", errors.Wrap(err, "get cache record error")
	}
	fdir, ok := value.Data().(string)
	if !ok {
		return fdir, errors.Wrap(err, "get cache record error")
	}
	return fdir, nil
}

func (c *Cacher) RemoveCacheRecord(fname string) error {
	if !c.cacher.Exists(fname) {
		return nil
	}
	value, err := c.cacher.Delete(fname)
	if err != nil {
		return errors.Wrap(err, "remove cache record error")
	}

	fpath, ok := value.Data().(string)
	if !ok {
		return errors.Wrap(errors.New("bad cache record"), "remove cache record error")
	}
	err = c.removeFile(fpath)
	return errors.Wrap(err, "remove cache record error")
}

func (c *Cacher) removeFile(fpath string) error {
	f, err := os.Stat(fpath)
	if err != nil {
		return errors.Wrap(err, "remove cached file error")
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	if f.Size() > c.usedSpace {
		return errors.Wrap(errors.New("bad file size"), "remove cached file error")
	}
	c.usedSpace -= f.Size()

	err = os.Remove(fpath)
	if err != nil {
		return errors.Wrap(err, "remove cached file error")
	}
	return nil
}
