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
