package docker

import (
	"github.com/assemblyline/spanner/assemblyfile"
	docker "github.com/fsouza/go-dockerclient"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Client struct {
	cgroup io.Reader
	Client *docker.Client
}

func New() Client {
	cgroup, err := os.Open("/proc/self/cgroup")
	if err != nil {
		panic(err)
	}
	return Client{
		cgroup: cgroup,
		Client: client(),
	}
}

func (d Client) SaveContainer(c assemblyfile.Config) string {
	options := docker.CommitContainerOptions{
		Container:  d.ContainerId(),
		Repository: c.Application.Repo,
	}

	image, err := d.Client.CommitContainer(options)
	if err != nil {
		panic(err)
	}

	return image.ID
}

func (d Client) ContainerId() string {
	s, err := ioutil.ReadAll(d.cgroup)
	if err != nil {
		panic(err)
	}
	cgroup := strings.Split(string(s[:]), "\n")[0]
	row := strings.Split(cgroup, "/")
	return row[len(row)-1]
}

func client() *docker.Client {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}
	return client
}
