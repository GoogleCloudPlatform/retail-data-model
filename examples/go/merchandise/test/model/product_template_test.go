package model

import (
	"encoding/json"
	"github.com/GoogleCloudPlatform/retail-data-model/enums"
	"github.com/GoogleCloudPlatform/retail-data-model/merchandise/pb"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func NewRequireStringAttribute() *ProductAttributeRule {
	return NewProductAttributeRule(
		pb.ProductAttributeRule_STRING,
		0,
		true,
		false,
		".*",
	)
}

func NewSafetyFeatures() *ProductAttributeRuleGroup {
	rg := NewProductAttributeRuleGroup(NewVersionID(), "Safety Features", "Common safety features")

	rg.AddProductAttributeRule(NewRequireStringAttribute().AddAttributeLabel(NewAttributeLabel(enums.Locale_EN_US, "Materials", "Consumer electronics products often have a variety of different features, such as different screen sizes, different camera resolutions, and different storage capacities.")))
	rg.AddProductAttributeRule(NewRequireStringAttribute().AddAttributeLabel(NewAttributeLabel(enums.Locale_EN_US, "Age Range", "Some consumer electronics products are specifically designed for children or adults.")))
	rg.AddProductAttributeRule(NewRequireStringAttribute().AddAttributeLabel(NewAttributeLabel(enums.Locale_EN_US, "Compliance", "Some consumer electronics products must comply with certain regulations. For example, a television must comply with the Federal Communications Commission (FCC) Part 15 regulations.")))
	return rg
}
func TestProductTemplate(t *testing.T) {

	template := NewProductTemplate(NewVersionID(), "Consumer Electronic", "Consumer electronics (CE) are electronic devices that are primarily used for entertainment, communication, or information processing.")

	color := NewRequireStringAttribute().AddAttributeLabel(NewAttributeLabel(enums.Locale_EN_US, "Color", "The primary color of the item."))
	features := NewRequireStringAttribute().AddAttributeLabel(NewAttributeLabel(enums.Locale_EN_US, "Features", "Consumer electronics products often have a variety of different features, such as different screen sizes, different camera resolutions, and different storage capacities."))

	template.addTemplateAttribute(NewProductTemplateAttributeFromRule(color))
	template.addTemplateAttribute(NewProductTemplateAttributeFromRule(features))
	template.addTemplateAttribute(NewProductTemplateAttributeFromGroup(NewSafetyFeatures()))

	assert.NotNil(t, template)
	assert.Equal(t, 3, len(template.Attributes))

	b, _ := json.Marshal(template)
	log.Printf(string(b))
}
