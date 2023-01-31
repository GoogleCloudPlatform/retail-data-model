package common

/*
Module is the fundamental building block of the Retail Data Model packages.
A common represents a single group of service, datasource, and other resources
required to run the common as a native cloud component group.

Modules control the life-cycle of the assets required to run, this includes
provisioning compute, creating datasets and schemas based on protobuf definitions
found in the API package.
*/

import (
	"log"
)

// Constants here define the naming conventions used by the database, logging,
// and pubsub topics used by the modules. This notion uses the following
// convention rules from the configuration file: <module_name>_?_environment
// this convention allows the common to exist across multiple projects and
// variant environments to support business and testing needs.
const (
	// LogNameFormat is a constant used as "<module_name>_<log_name>_<environment>"
	LogNameFormat  = "%s_%s_%s"
	DBNameFormat   = "%s_%s"
	DefaultLogName = "app"
	DBLogName      = "bq"
)

// Module is the basic building block for retail domain modules. Each common
// represents a clear-cut subdomain model of retail. Please see the default
// documentation for more details on available modules.
type Module interface {
	// GetConfig returns a Configuration object
	GetConfig() *Config
	// GetBigQuery returns a BigQuery wrapper object
	GetBigQuery() *BigQuery
	// GetLogger returns the underlying cloud logger
	GetLogger() *Logger
	// GetDefaultLog returns a native log with a default criticality level.
	GetDefaultLog() *log.Logger
	// Close safely releases all resources
	Close()
}
