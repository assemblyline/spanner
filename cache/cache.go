package cache

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/assemblyline/spanner/logger"
	"github.com/docker/docker/pkg/archive"
	"io"
)

//The Store interface allows storeage to be pluged in to back a Cache
type Store interface {
	Writer(name string) (io.WriteCloser, error)
	Reader(name string) (io.ReadCloser, error)
}

//A Cache provides a mechinisim to store and restore expensive to generate files between build generations
type Cache struct {
	Store Store
	Hash  []byte
	log   logger.Logger
}

//New returns a new cache, for the resource identified by hash, backed by store.
func New(hash []byte, store Store) Cache {
	return Cache{
		Store: store,
		Hash:  hash,
		log:   logger.New(),
	}
}

//Save persits the contents of dir to the cache, task identifies the task (external) that manages the files there
func (c Cache) Save(dir string, task string) {
	tarball, err := archive.Tar(dir, archive.Gzip)
	checkerr(err)

	cacheWriter, err := c.Store.Writer(c.path(dir, task))
	checkerr(err)

	_, err = io.Copy(cacheWriter, tarball)
	checkerr(err)

	cacheWriter.Close()
	c.log.Info("Cache for", dir, "saved as", c.path(dir, task))
}

//Restore restores the contents of dir from the cache
func (c Cache) Restore(dir, task string) {
	path := c.path(dir, task)
	cacheReader, err := c.Store.Reader(c.path(dir, task))
	if cacheReader != nil {
		err = archive.Untar(cacheReader, dir, &archive.TarOptions{})
		c.log.Info("Restoring cache for", dir, "from", path)
		if err != nil {
			c.log.Error("Error restoring cache for", dir, "from", path, err.Error())
		}
	} else {
		c.log.Info("No Cache for", dir, "to restore")
	}
}

func (c Cache) hash(dir, task string) string {
	hasher := sha256.New()

	hasher.Write(c.Hash)
	hasher.Write([]byte(dir))
	hasher.Write([]byte(task))

	return hex.EncodeToString(hasher.Sum(nil))
}

func (c Cache) path(dir, task string) string {
	return c.hash(dir, task) + ".tar.gz"
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
