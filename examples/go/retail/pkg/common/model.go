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
