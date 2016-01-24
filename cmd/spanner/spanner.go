package main

import (
	"fmt"
	"github.com/assemblyline/spanner/assemblyfile"
	"github.com/assemblyline/spanner/cache"
	"github.com/assemblyline/spanner/config"
	"github.com/assemblyline/spanner/logger"
	"os"
)

func handleFatal() {
	if err := recover(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var l = logger.New()
var fileStore = cache.NewFileStore("/var/assemblyline/cache/")

func main() {
	defer handleFatal()

	confFile, err := os.Open("/etc/assemblyline/spanner/" + os.Args[1] + ".toml")
	check(err)
	config, err := config.Read(confFile)
	check(err)

	assemblyFile, err := os.Open("Assemblyfile")
	check(err)
	af, err := assemblyfile.Read(assemblyFile)
	check(err)

	c := cache.New(af.Hash(), fileStore)

	l.Title("Building", af.Application.Name, "on", config.Builder.Name, config.Builder.Version)
	for _, step := range config.Step {
		step.Cache = c
		err := step.Exec()
		check(err)
	}
}
