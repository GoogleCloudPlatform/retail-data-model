package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	merch "github.com/GoogleCloudPlatform/retail-data-model/merchandise/pb"
)

type Template merch.Template

func NewTemplate(id *VersionID, name string, description string) *Template {
	return &Template{
		Id:             (*common.VersionID)(id),
		Name:           name,
		Description:    description,
		AttributeRules: make([]*merch.TemplateAttributeRule, 0),
	}
}

func (template *Template) addTemplateAttributeRule(attribute *TemplateAttributeRule) {
	template.AttributeRules = append(template.AttributeRules, (*merch.TemplateAttributeRule)(attribute))
}

// Template Attribute Rules

type TemplateAttributeRule merch.TemplateAttributeRule

func NewTemplateAttributeRule(t merch.TemplateAttributeRule_FieldType, ordinal int32, required bool, allowOverride bool, validationRegEx string) *TemplateAttributeRule {
	return &TemplateAttributeRule{
		FieldType:       t,
		Ordinal:         ordinal,
		Required:        required,
		AllowOverride:   allowOverride,
		ValidationRegex: validationRegEx,
	}
}
