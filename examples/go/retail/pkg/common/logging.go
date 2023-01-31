/*
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"context"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/logging"
)

/*
logging.go is a utility for loading and working with a cloud logger
*/

// Logger is the primary wrapper for holding the logging clinet and map
// of available / registered loggers and the configured severity level.
type Logger struct {
	client   *logging.Client
	loggers  map[string]*logging.Logger
	severity logging.Severity
}

// newLogger is a builder method used to generate the logger and the client.
// this is only available from the common package.
func newLogger(ctx context.Context, projectId string, logLevel string) (logger *Logger, err error) {
	client, err := logging.NewClient(ctx, projectId)
	if err == nil {
		logger = &Logger{client: client, severity: parseLogLevel(logLevel), loggers: make(map[string]*logging.Logger)}
	}
	return logger, err
}

// Close is used to close each open log, and ensure they are flushed prior
// to the system halting
func (logger *Logger) close() error {
	for k, v := range logger.loggers {
		v.StandardLogger(logger.severity).Println(fmt.Sprintf("Closing logger: %s", k))
		err := v.Flush()
		if err != nil {
			log.Default().Printf("Failed to flush logger: %s", k)
		}
	}
	return logger.client.Close()
}

// GetLogger returns a cloud logger
func (logger *Logger) GetLogger(name string) (log *logging.Logger) {
	if l, ok := logger.loggers[name]; ok {
		log = l
	} else {
		logger.loggers[name] = logger.client.Logger(name)
		log = logger.loggers[name]
	}
	return log
}

// GetStandardLogger returns a standard logger
func (logger *Logger) GetStandardLogger(name string) (log *log.Logger) {
	return logger.GetLogger(name).StandardLogger(logger.severity)
}

// parseLogLevel is used to parse the text input and ensure it's a valid
// value for the logging object.
func parseLogLevel(logLevel string) logging.Severity {
	switch lvl := strings.ToUpper(logLevel); lvl {
	case "DEBUG":
		return logging.Debug
	case "INFO":
		return logging.Info
	case "WARNING":
		return logging.Warning
	case "ERROR":
		return logging.Error
	case "EMERGENCY":
		return logging.Emergency
	case "CRITICAL":
		return logging.Critical
	default:
		return logging.Default
	}
}
