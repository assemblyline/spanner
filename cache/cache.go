package cache

import(
  "crypto/sha256"
  "encoding/hex"
  "github.com/assemblyline/spanner/assemblyfile"
  "github.com/assemblyline/spanner"
  "github.com/docker/docker/pkg/archive"
  "io"
  "os"
)

const rootDir string = "/Users/ed/al/cache/"

type Cache struct {
  Dir string
  Hash string
  Assemblyfile assemblyfile.Config
}

func New(dir string, config assemblyfile.Config) Cache {
  return Cache{
    Dir: dir,
    Hash: hash(config),
    Assemblyfile: config,
  }
}

func (c Cache) Save() {
  tarball, err := archive.Tar(c.Dir, archive.Gzip)
  checkerr(err)

  cacheFile, err := os.Create(c.path())
  checkerr(err)

  _, err = io.Copy(cacheFile, tarball)
  checkerr(err)
  cacheFile.Close()
  spanner.LogInfo("Cache for", c.Dir, "saved as", c.path())
}

func (c Cache) Restore() {
  path := c.path()
  if _, err := os.Stat(path); err == nil {
    err := archive.UntarPath(path, c.Dir)
    checkerr(err)
    spanner.LogInfo("Cache for", c.Dir, "restored from", path)
  } else {
    spanner.LogInfo("No Cache for", c.Dir, "to restore")
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
  return rootDir + c.Hash + ".tar.gz"
}

func checkerr(err error) {
  if err != nil {
    panic(err)
  }
}
