package model

import (
	"encoding/json"
	"github.com/GoogleCloudPlatform/retail-data-model/merchandise/pb"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func NewRequireStringAttribute() *ProductTemplateAttributeRule {
	return NewProductAttributeRule(
		pb.ProductTemplateAttributeRule_STRING,
		0,
		true,
		false,
		".*",
	)
}

func NewSafetyFeatures() *ProductTemplateAttributeRuleGroup {
	rg := NewProductAttributeRuleGroup(NewVersionID(), "Materials", "Common safety features")

	return rg
}
func TestProductTemplate(t *testing.T) {
	template := NewProductTemplate(NewVersionID(), "Consumer Electronic", "Consumer electronics (CE) are electronic devices that are primarily used for entertainment, communication, or information processing.")
	template.addTemplateAttribute(NewProductTemplateAttributeFromGroup(NewSafetyFeatures()))

	assert.NotNil(t, template)
	assert.Equal(t, 3, len(template.Attributes))

	b, _ := json.Marshal(template)
	log.Printf(string(b))
}
