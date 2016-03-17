package docker

import (
	"github.com/assemblyline/spanner/assemblyfile"
	docker "github.com/fsouza/go-dockerclient"
	"io/ioutil"
	"os"
	"strings"
)

//Client for interactions with the docker daemon
type Client struct {
	cgroup string
	Client *docker.Client
}

//SaveContainer commits the current container
func (d Client) SaveContainer(c assemblyfile.Config) string {
	options := docker.CommitContainerOptions{
		Container:  d.ContainerID(),
		Repository: c.Application.Repo,
	}

	image, err := d.Client.CommitContainer(options)
	if err != nil {
		panic(err)
	}

	return image.ID
}

//ContainerID returns the ID of the current container
func (d Client) ContainerID() string {
	if d.cgroup == "" {
		d.cgroup = "/proc/self/cgroup"
	}
	cgroup, err := os.Open(d.cgroup)
	if err != nil {
		panic(err)
	}
	s, err := ioutil.ReadAll(cgroup)
	if err != nil {
		panic(err)
	}
	cgroups := strings.Split(string(s[:]), "\n")[0]
	row := strings.Split(cgroups, "/")
	return row[len(row)-1]
}

func client() *docker.Client {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}
	return client
}
