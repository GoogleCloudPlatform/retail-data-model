package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/enums"
	merchandise "github.com/GoogleCloudPlatform/retail-data-model/merchandise/pb"
)

type ProductTemplate merchandise.ProductTemplate

func NewProductTemplate(id *VersionID, name string, description string) *ProductTemplate {
	return &ProductTemplate{
		Id:          (*common.VersionID)(id),
		Name:        name,
		Description: description,
	}
}

func (productTemplate *ProductTemplate) addTemplateAttribute(attribute *ProductTemplateAttribute) {
	productTemplate.Attributes = append(productTemplate.Attributes, (*merchandise.ProductTemplateAttribute)(attribute))
}

// Template Attribute

type ProductTemplateAttribute merchandise.ProductTemplateAttribute

func NewProductTemplateAttributeFromGroup(group *ProductAttributeRuleGroup) *ProductTemplateAttribute {
	_group := &merchandise.ProductTemplateAttribute_RuleGroup{RuleGroup: (*merchandise.ProductAttributeRuleGroup)(group)}
	return &ProductTemplateAttribute{
		Value: _group,
	}
}

func NewProductTemplateAttributeFromRule(rule *ProductAttributeRule) *ProductTemplateAttribute {
	_attr := &merchandise.ProductTemplateAttribute_Rule{Rule: (*merchandise.ProductAttributeRule)(rule)}
	return &ProductTemplateAttribute{
		Value: _attr,
	}
}

// Template Attribute Rule Group

type ProductAttributeRuleGroup merchandise.ProductAttributeRuleGroup

func NewProductAttributeRuleGroup(id *VersionID) *ProductAttributeRuleGroup {
	return &ProductAttributeRuleGroup{
		Id:                    (*common.VersionID)(id),
		ProductAttributeRules: make([]*merchandise.ProductAttributeRule, 0),
	}
}

func (group *ProductAttributeRuleGroup) AddProductAttributeRule(rule *ProductAttributeRule) {
	group.ProductAttributeRules = append(group.ProductAttributeRules, (*merchandise.ProductAttributeRule)(rule))
}

// Template Attribute Rules

type ProductAttributeRule merchandise.ProductAttributeRule
type ProductAttributeRuleType merchandise.ProductAttributeRule_Type

func NewProductAttributeRule(t ProductAttributeRuleType, ordinal int32, required bool, allowOverride bool, validationRegEx string) *ProductAttributeRule {
	return &ProductAttributeRule{
		Type:            (merchandise.ProductAttributeRule_Type)(t),
		Ordinal:         ordinal,
		Required:        required,
		AllowOverride:   allowOverride,
		ValidationRegEx: validationRegEx,
		Labels:          make([]*merchandise.AttributeLabel, 0),
	}
}

func (rule *ProductAttributeRule) AddAttributeLabel(label *AttributeLabel) {
	rule.Labels = append(rule.Labels, (*merchandise.AttributeLabel)(label))
}

type AttributeLabel merchandise.AttributeLabel

func NewAttributeLabel(locale enums.Locale, name string, description string) *AttributeLabel {
	return &AttributeLabel{Locale: locale, Name: name, Description: description}
}

type ProductAttributeRuleGroupValue merchandise.ProductAttributeRuleGroupValue

type ProductAttributeRuleValue merchandise.ProductAttributeRuleValue

func (n *ProductAttributeRuleValue) AddStringOrNumber(num float64, str string) {
	switch v := n.Value.(type) {
	case *merchandise.ProductAttributeRuleValue_NumberValue:
		v.NumberValue.Value = append(v.NumberValue.Value, num)
		break
	case *merchandise.ProductAttributeRuleValue_StringValue:
		v.StringValue.Value = append(v.StringValue.Value, str)
		break
	}
}

type NumericProductAttributeRuleValue ProductAttributeRuleValue

func NewNumericProductAttributeRuleValue(id *VersionID, ordinal int32) *NumericProductAttributeRuleValue {
	return &NumericProductAttributeRuleValue{
		ProductTemplateId: (*common.VersionID)(id),
		FieldOrdinal:      ordinal,
		Value: &merchandise.ProductAttributeRuleValue_NumberValue{
			NumberValue: &merchandise.ProductAttributeRuleValue_Number{
				Value: make([]float64, 0),
			},
		},
	}
}

func (n *NumericProductAttributeRuleValue) AddNumber(num float64) {
	(*ProductAttributeRuleValue)(n).AddStringOrNumber(num, "")
}

type StringProductAttributeRuleValue ProductAttributeRuleValue

func NewStringProductAttributeRuleValue(id *VersionID, ordinal int32) *StringProductAttributeRuleValue {
	return &StringProductAttributeRuleValue{
		ProductTemplateId: (*common.VersionID)(id),
		FieldOrdinal:      ordinal,
		Value: &merchandise.ProductAttributeRuleValue_StringValue{
			StringValue: &merchandise.ProductAttributeRuleValue_String{
				Value: make([]string, 0),
			},
		},
	}
}

func (n *StringProductAttributeRuleValue) AddString(val string) {
	(*ProductAttributeRuleValue)(n).AddStringOrNumber(0, val)
}
