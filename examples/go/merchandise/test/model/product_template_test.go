package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestProductTemplate(t *testing.T) {
	template := NewTemplate(NewVersionID(), "Consumer Electronic", "Consumer electronics (CE) are electronic devices that are primarily used for entertainment, communication, or information processing.")
	assert.NotNil(t, template)
	b, _ := json.Marshal(template)
	log.Printf(string(b))
}
