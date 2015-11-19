package config

import (
	"github.com/assemblyline/spanner/step"
	"github.com/naoina/toml"
	"io"
	"io/ioutil"
)

func Read(reader io.Reader) (Config, error) {
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := toml.Unmarshal(buf, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

type Config struct {
	Builder Builder
	Step    []step.Step
}

type Builder struct {
	Name    string
	Task    string
	Version string
}
