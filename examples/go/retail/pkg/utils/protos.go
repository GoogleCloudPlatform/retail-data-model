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

package utils

import (
	"errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ConvertMessageToType - Converts a raw message to a specific message type. This helper class is useful for
// converting a struct to a message
func ConvertMessageToType(from proto.Message, to proto.Message) (err error) {
	if from != nil {
		if to != nil {
			var b []byte
			b, err = protojson.Marshal(from)
			err = protojson.Unmarshal(b, to)
		} else {
			err = errors.New("to may not be nil")
		}
	} else {
		err = errors.New("from may not be nil")
	}
	return err
}
