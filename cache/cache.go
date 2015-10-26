package cache

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/assemblyline/spanner/assemblyfile"
	"github.com/assemblyline/spanner/logger"
	"github.com/docker/docker/pkg/archive"
	"io"
	"os"
)

type CacheStore interface {
	WriteCloser(name string) (io.WriteCloser, error)
	ReadCloser(name string) (io.ReadCloser, error)
}

type FileStore struct {
	dir string
}

func (f FileStore) WriteCloser(name string) (io.WriteCloser, error) {
	return os.Create(f.dir + "/" + name)
}

func (f FileStore) ReadCloser(name string) (io.ReadCloser, error) {
	return os.Open(f.dir + "/" + name)
}

type Cache struct {
	CacheStore   CacheStore
	Dir          string
	Hash         string
	Assemblyfile assemblyfile.Config
	log          logger.Logger
}

func New(dir string, config assemblyfile.Config, store CacheStore) Cache {
	return Cache{
		CacheStore:   store,
		Dir:          dir,
		Hash:         hash(config),
		Assemblyfile: config,
		log:          logger.New(),
	}
}

func (c Cache) Save() {
	tarball, err := archive.Tar(c.Dir, archive.Gzip)
	checkerr(err)

	cacheWriter, err := c.CacheStore.WriteCloser(c.path())
	checkerr(err)

	_, err = io.Copy(cacheWriter, tarball)
	checkerr(err)

	cacheWriter.Close()
	c.log.Info("Cache for", c.Dir, "saved as", c.path())
}

func (c Cache) Restore() {
	path := c.path()
	cacheReader, err := c.CacheStore.ReadCloser(c.path())
	err = archive.Untar(cacheReader, c.Dir, &archive.TarOptions{})

	if err != nil {
		c.log.Info("No Cache for", c.Dir, "to restore")
	} else {
		c.log.Info("Cache for", c.Dir, "restored from", path)
	}
}

func hash(a assemblyfile.Config) string {
	hasher := sha256.New()

	hasher.Write([]byte(a.Application.Name))
	hasher.Write([]byte(a.Application.Repo))
	hasher.Write([]byte(a.Build.Builder))
	hasher.Write([]byte(a.Build.Version))

	return hex.EncodeToString(hasher.Sum(nil))
}

func (c Cache) path() string {
	return hash(c.Assemblyfile) + ".tar.gz"
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
