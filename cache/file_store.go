package cache

import (
	"errors"
	"io"
	"os"
)

// FileStore impliments the cache.Store interface and provides a way of storing
// cached content on the local filesystem
type FileStore struct {
	dir string
}

func NewFileStore(dir string) FileStore {
	return FileStore{
		dir: dir,
	}
}

func (f FileStore) Writer(name string) (io.WriteCloser, error) {
	return os.Create(f.dir + "/" + name)
}

func (f FileStore) Reader(name string) (io.ReadCloser, error) {
	path := f.dir + "/" + name
	if _, err := os.Stat(path); err == nil {
		return os.Open(path)
	}
	return nil, errors.New("Cache not found")
}
