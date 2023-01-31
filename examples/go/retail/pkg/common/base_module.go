package common

import (
	"context"
	"fmt"
	"log"
	"strings"
)

// BaseModule is the primary structure used for holding state objects.
type BaseModule struct {
	// Config - The configuration loaded into the common
	Config *Config
	// BigQuery - The limited BigQuery object, hiding the client and close method
	// to ensure safe operation.
	BigQuery *BigQuery
	// logger - The cloud logger, with a hidden close method to ensure safe
	// behaviors of reference passing.
	logger *Logger
	// log - A standard log object used for recoding events into cloud stored
	// logs
	log *log.Logger
}

// LoadBaseModule - Reads a JSON configuration from disk and initializes the common.
// if any errors occur, a non-nil error will be returned. The intention of
// LoadModule is to be used in conjunction with the Main method of an implementing
// service.
func LoadBaseModule(configFile string) (module Module, err error) {
	var config *Config
	config, err = LoadConfig(configFile)
	if err == nil {
		module, err = NewBaseModule(config)
	}
	return module, err
}

// NewBaseModule is a programmatic constructor using a Config object.
func NewBaseModule(config *Config) (module Module, err error) {
	m := &BaseModule{Config: config}
	ctx := context.Background()
	if m.Config.Enabled {
		err = m.initializeLogger(ctx)
		if err == nil {
			// Initialize the default Logger
			m.log = m.logger.GetStandardLogger(m.generateLogName(DefaultLogName))
			// Initialize Big Query
			err = m.initializeBigQuery(ctx)
			if err == nil {
				module = m
			}
		}
	}
	return module, err
}

// initializeLogger reads the Config and initializes a logger based on the values
// of the configuration file.
func (module *BaseModule) initializeLogger(ctx context.Context) error {
	logger, err := newLogger(ctx, module.Config.ProjectId, module.Config.LogLevel)
	if err == nil {
		module.logger = logger
	}
	return err
}

// initializeBigQuery reads the Config and initializes the dataset.
func (module *BaseModule) initializeBigQuery(ctx context.Context) error {
	bq, err := NewBigQuery(
		ctx,
		module.Config.ProjectId,
		module.logger.GetStandardLogger(module.generateLogName(DBLogName)),
		module.generateDBName())
	if err == nil {
		module.BigQuery = bq
	}
	return err
}

// GetDefaultLogger - returns a default logger
func (module *BaseModule) GetDefaultLogger() *log.Logger {
	return module.log
}

// generateLogName is a helper method for creating the log name based on the
// prefix and suffix or common name and environment.
func generateLogName(prefix string, logName string, suffix string) string {
	return strings.ToLower(fmt.Sprintf(LogNameFormat, prefix, logName, suffix))
}

// GenerateLogName is a helper method on the Module for generating log names.
func (module *BaseModule) generateLogName(logName string) string {
	return generateLogName(module.Config.ProjectId, logName, module.Config.Environment)
}

// generateDBName is a helper method used to generate the dataset name
func (module *BaseModule) generateDBName() string {
	return strings.ToLower(fmt.Sprintf(DBNameFormat, module.Config.Name, module.Config.Environment))
}

func (module *BaseModule) GetConfig() *Config {
	return module.Config
}

func (module *BaseModule) GetBigQuery() *BigQuery {
	return module.BigQuery
}

func (module *BaseModule) GetLogger() *Logger {
	return module.logger
}

func (module *BaseModule) GetDefaultLog() *log.Logger {
	return module.log
}

// Close - is used to safely close all related objects and instances.
func (module *BaseModule) Close() {
	_ = module.BigQuery.close()
	module.log.Printf("Closed BQ Client")
	_ = module.logger.close()
}
