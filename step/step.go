package step

import (
	"github.com/assemblyline/spanner/cache"
	"github.com/assemblyline/spanner/logger"
	"os/exec"
)

var log = logger.New()

type Step struct {
	Name     string
	CacheDir string
	Cache    cache.Cache
	Script   [][]string
}

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
	if s.CacheDir != "" {
		s.Cache.Restore(s.CacheDir, s.Name)
	}
}

func (s Step) save() {
	if s.CacheDir != "" {
		s.Cache.Save(s.CacheDir, s.Name)
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
