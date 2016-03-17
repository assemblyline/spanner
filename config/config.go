package config

import (
	"github.com/assemblyline/spanner/step"
	"github.com/naoina/toml"
	"io"
	"io/ioutil"
)

//Read reads a spanner config from a reader, and returns the Config object
//returns an error if the config could not be read, or if it is invalid
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

//Config represents a set of Configuration for a spanner
type Config struct {
	Spanner Spanner
	Step    []step.Step
}

//Spanner represents some metadata about this spanner
type Spanner struct {
	Name    string
	Task    string
	Version string
}
