package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	merchandise "github.com/GoogleCloudPlatform/retail-data-model/merchandise/pb"
)

type ProductTemplate merchandise.ProductTemplate

func NewProductTemplate(id *VersionID, name string, description string) *ProductTemplate {
	return &ProductTemplate{
		Id:          (*common.VersionID)(id),
		Name:        name,
		Description: description,
		Attributes:  make([]*merchandise.ProductTemplateAttribute, 0),
	}
}

func (productTemplate *ProductTemplate) addTemplateAttribute(attribute *ProductTemplateAttribute) {
	productTemplate.Attributes = append(productTemplate.Attributes, (*merchandise.ProductTemplateAttribute)(attribute))
}

// Template Attribute

type ProductTemplateAttribute merchandise.ProductTemplateAttribute

func NewProductTemplateAttributeFromGroup(group *ProductTemplateAttributeRuleGroup) *ProductTemplateAttribute {
	_group := &merchandise.ProductTemplateAttribute_RuleGroup{RuleGroup: (*merchandise.ProductTemplateAttributeRuleGroup)(group)}
	return &ProductTemplateAttribute{
		Value: _group,
	}
}

func NewProductTemplateAttributeFromRule(rule *ProductTemplateAttributeRule) *ProductTemplateAttribute {
	_attr := &merchandise.ProductTemplateAttribute_Rule{Rule: (*merchandise.ProductTemplateAttributeRule)(rule)}
	return &ProductTemplateAttribute{
		Value: _attr,
	}
}

// Template Attribute Rule Group

type ProductTemplateAttributeRuleGroup merchandise.ProductTemplateAttributeRuleGroup

func NewProductAttributeRuleGroup(id *VersionID, name string, description string) *ProductTemplateAttributeRuleGroup {
	return &ProductTemplateAttributeRuleGroup{
		ProductTemplateAttributeRules: make([]*merchandise.ProductTemplateAttributeRule, 0),
	}
}

func (group *ProductTemplateAttributeRuleGroup) AddProductAttributeRule(rule *ProductTemplateAttributeRule) {
	group.ProductTemplateAttributeRules = append(group.ProductTemplateAttributeRules, (*merchandise.ProductTemplateAttributeRule)(rule))
}

// Template Attribute Rules

type ProductTemplateAttributeRule merchandise.ProductTemplateAttributeRule

func NewProductAttributeRule(t merchandise.ProductTemplateAttributeRule_FieldType, ordinal int32, required bool, allowOverride bool, validationRegEx string) *ProductTemplateAttributeRule {
	return &ProductTemplateAttributeRule{
		FieldType:       t,
		Ordinal:         ordinal,
		Required:        required,
		AllowOverride:   allowOverride,
		ValidationRegex: validationRegEx,
	}
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
		FieldOrdinal: ordinal,
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
		FieldOrdinal: ordinal,
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
