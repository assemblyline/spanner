package assemblyfile_test

import (
	"errors"
	af "github.com/assemblyline/spanner/assemblyfile"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

type unreadable struct{}

func (u unreadable) Read(p []byte) (i int, err error) {
	return 0, errors.New("could not read, because broken")
}

func fixture() af.Config {
	assemblyfile, _ := os.Open("Assemblyfile.fixture")
	config, _ := af.Read(assemblyfile)
	return config
}

func rawConfig() []byte {
	buf, _ := ioutil.ReadFile("Assemblyfile.fixture")
	return buf
}

func TestAssemblyfileApplicationConfig(t *testing.T) {
	config := fixture()

	expected := "Test App"
	if config.Application.Name != expected {
		t.Error("Expected", config.Application.Name, "to equal", expected)
	}

	expected = "foo.example.com/assemblyline/test"
	if config.Application.Repo != expected {
		t.Error("Expected", config.Application.Repo, "to equal", expected)
	}
}

func TestAssemblyfileParse(t *testing.T) {
	config, err := af.Parse(rawConfig())
	if err != nil {
		t.Error("Expected error to be nil")
	}

	expected := "Test App"
	if config.Application.Name != expected {
		t.Error("Expected", config.Application.Name, "to equal", expected)
	}

	expected = "foo.example.com/assemblyline/test"
	if config.Application.Repo != expected {
		t.Error("Expected", config.Application.Repo, "to equal", expected)
	}
}

func TestAssemblyfileBuildConfig(t *testing.T) {
	config := fixture()

	expected := "ruby"
	if config.Build.Builder != expected {
		t.Error("Expected", config.Build.Builder, "to equal", expected)
	}

	expected = "2.2.3"
	if config.Build.Version != expected {
		t.Error("Expected", config.Build.Builder, "to equal", expected)
	}
}

func TestAssemblyfileTestConfig(t *testing.T) {
	config := fixture()

	expected := []string{
		"bundle exec rake db:test:prepare",
		"bundle exec rake",
	}

	if !reflect.DeepEqual(config.Test.Script, expected) {
		t.Error("Expected", config.Test.Script, "to equal", expected)
	}

	expectedMap := map[string]interface{}{"RACK_ENV": "test", "AWESOME": true}

	if !reflect.DeepEqual(config.Test.Env, expectedMap) {
		t.Error("Expected", config.Test.Env, "to equal", expectedMap)
	}

	if len(config.Test.Service) != 3 {
		t.Error("Expected there to be 3 services")
	}

	expectedVersion := "9.4.1"
	postgresVersion := config.Test.Service["postgres"].Version
	if postgresVersion != expectedVersion {
		t.Error("Expected postgres version to be", expectedVersion, "but was", postgresVersion)
	}

	expectedProperties := "es.script.groovy.sandbox.enabled=true"
	properties := config.Test.Service["elasticsearch"].Properties[0]
	if expectedProperties != properties {
		t.Error("Expected elasticsearch properties to be", expectedProperties, "but was", properties)
	}
}

func TestAssemblyfileErrors(t *testing.T) {
	//io error
	u := unreadable{}
	_, err := af.Read(u)
	expected := "could not read, because broken"
	if err.Error() != expected {
		t.Error("Expected", err.Error(), "to equal", expected)
	}

	//invalid assemblyfile
	assemblyfile, _ := os.Open("Assemblyfile.broken")
	_, err = af.Read(assemblyfile)
	expected = "toml: unmarshal"
	if err.Error()[0:15] != expected {
		t.Error("Expected", err.Error()[0:15], "to equal", expected)
	}
}
