package module

import (
	"flag"
	"os"
	"testing"

	"github.com/rrmcguinness/retail-data-model/events/pkg/module"
)

var eventModule *module.EventModule

const (
	ConfigFile = "-configFile=examples/configs/events-test-config.json"
	SampleRate = "-sampleRate=1"
)

func setup() {
	var err error
	os.Args = append(os.Args, ConfigFile)
	os.Args = append(os.Args, SampleRate)

	// Parse all defined flags
	flag.Parse()

	// Load the event module
	eventModule, err = module.Load()
	if err != nil {
		panic(err)
	}
}

func teardown() {
	if eventModule != nil {
		eventModule.Close()
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
