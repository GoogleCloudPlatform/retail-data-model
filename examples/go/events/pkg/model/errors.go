package model

import (
	"errors"
)

var ConversionError = errors.New("failed to convert event type")
var RecordIdExists = errors.New("record id already exists")
var RecordNotFound = errors.New("record not found")
var RecordNotChanged = errors.New("the update record is equal to the existing record, no change made")
