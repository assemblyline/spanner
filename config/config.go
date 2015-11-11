package config

import (
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
	Step    []Step
}

type Builder struct {
	Name    string
	Version string
}

type Step struct {
	Cache         string
	Script        []string
	RequiredFiles []string
}
