package model

import (
	"encoding/json"
	"github.com/GoogleCloudPlatform/retail-data-model/merchandise/pb"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func NewRequireStringAttribute() *CategoryTemplateAttributeRule {
	return NewCategoryTemplateAttributeRule(
		pb.CategoryTemplateAttributeRule_STRING,
		0,
		true,
		false,
		".*",
	)
}

func TestProductTemplate(t *testing.T) {
	template := NewCategoryTemplate(NewVersionID(), "Consumer Electronic", "Consumer electronics (CE) are electronic devices that are primarily used for entertainment, communication, or information processing.")
	assert.NotNil(t, template)
	b, _ := json.Marshal(template)
	log.Printf(string(b))
}
