package assemblyfile

import (
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
)

func Read(path string) Config {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := toml.Unmarshal(buf, &config); err != nil {
		panic(err)
	}
	return config
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
