package step

import (
	"github.com/assemblyline/spanner/cache"
	"github.com/assemblyline/spanner/logger"
	"os/exec"
)

var log = logger.New()

//Step represents a runnable script that updates the contents of a particular directory
//Optionly caches that directory
type Step struct {
	Name   string
	Dir    string
	Cache  cache.Cache
	Script [][]string
}

//Exec runs the script in order to update the filesystem
//If the step is cacheable the cache is restored to the filesystem before the script is run and updated after
func (s Step) Exec() error {
	log.StepTitle(s.Name)
	s.restore()
	for _, command := range s.Script {
		if err := run(command[0], command[1:]...); err != nil {
			return err
		}
	}
	s.save()
	return nil
}

func (s Step) restore() {
	if s.Dir != "" {
		s.Cache.Restore(s.Dir, s.Name)
	}
}

func (s Step) save() {
	if s.Dir != "" {
		s.Cache.Save(s.Dir, s.Name)
	}
}

func run(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = log.Out()
	cmd.Stderr = log.Err()
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
