package assemblyfile

import (
	"github.com/naoina/toml"
	"io"
	"io/ioutil"
)

//Read and parse the Assemblyfile from an io.Reader
func Read(assemblyfile io.Reader) (Config, error) {
	buf, err := ioutil.ReadAll(assemblyfile)
	if err != nil {
		return Config{}, err
	}
	return Parse(buf)
}

//Parse the Assemblyfile from the raw bytes
func Parse(rawConfig []byte) (Config, error) {
	var config Config
	if err := toml.Unmarshal(rawConfig, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

type Config struct {
	Application Application
	Build       Build
	Test        Test
}

func (c Config) Hash() []byte {
	return []byte(c.Application.Name + c.Application.Repo + c.Build.Builder + c.Build.Version)
}

type Application struct {
	Name string
	Repo string
}

type Build struct {
	Builder string
	Version string
}

type Test struct {
	Script  []string
	Env     map[string]interface{}
	Service map[string]Service
}

type Service struct {
	Version    string
	Properties []string
}
