package cache

import (
	"errors"
	"io"
	"os"
)

// FileStore impliments the cache.Store interface and provides a way of storing
// cached content on the local filesystem
type FileStore struct {
	Dir string
}

//Writer returns an io.WriteCloser to allow items to be saved to the named file
func (f FileStore) Writer(name string) (io.WriteCloser, error) {
	return os.Create(f.Dir + "/" + name)
}

//Reader returns an io.ReadCloser to allow items to be read to the named file
//returns an error if the named file is not present
func (f FileStore) Reader(name string) (io.ReadCloser, error) {
	path := f.Dir + "/" + name
	if _, err := os.Stat(path); err == nil {
		return os.Open(path)
	}
	return nil, errors.New("Cache not found")
}
