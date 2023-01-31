package common

import (
	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	"google.golang.org/protobuf/proto"
)

// MessageLoaderWrapper - is a wrapper used to reduce total number of
// dependencies in other common projects.
type MessageLoaderWrapper struct {
	loader protobq.MessageLoader
}

// Load is a helper method for loading bq values based on a schema
func (w *MessageLoaderWrapper) Load(values []bigquery.Value, schema bigquery.Schema) error {
	return w.loader.Load(values, schema)
}

// GetMessage is a helper method for converting a wrapped message into a
// protobuf message
func (w *MessageLoaderWrapper) GetMessage() proto.Message {
	return w.loader.Message
}
