package assemblyfile

import (
	"github.com/naoina/toml"
	"io"
	"io/ioutil"
)

func Read(assemblyfile io.Reader) (Config, error) {
	buf, err := ioutil.ReadAll(assemblyfile)
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
	Application Application
	Build       Build
	Test        Test
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
