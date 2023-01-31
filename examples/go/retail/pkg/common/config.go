package common

/*
config.go contains the definitions for the objects required to read
configuration files.
*/

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const (
	FlagConfigFile = "configFile"
)

// KeyValuePair is a simple stucture for holding key value pairs
type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Config is the primary configuration state holder and TO definition.
type Config struct {
	ProjectId   string            `json:"project_id"`
	Name        string            `json:"name"`
	Environment string            `json:"env"`
	Enabled     bool              `json:"enabled"`
	LogLevel    string            `json:"log_level"`
	Labels      map[string]string `json:"labels"`
}

// NewConfig is the programmatic constructor for the configuration struct.
func NewConfig(
	name string,
	environment string,
	enabled bool,
	logLevel string) *Config {

	return &Config{
		Name:        name,
		Environment: environment,
		Enabled:     enabled,
		LogLevel:    logLevel,
		Labels:      make(map[string]string)}
}

// LoadConfig is used for loading a configuration file from disk.
func LoadConfig(fileName string) (moduleConfig *Config, err error) {
	var b []byte
	b, err = os.ReadFile(fileName)
	if err == nil {
		tCft := &Config{}
		err = json.Unmarshal(b, tCft)
		if err == nil {
			moduleConfig = tCft
		}
	}
	return moduleConfig, err
}

func LoadConfigFromFlags() (*Config, error) {
	configFile := flag.String(FlagConfigFile, "", "The relative or exact path of a configuration file.")
	fmt.Printf("-----------------------CONFIG FILE---------- %s \n\n", *configFile)
	return LoadConfig(*configFile)
}
