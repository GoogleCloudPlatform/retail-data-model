package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	merchandise "github.com/GoogleCloudPlatform/retail-data-model/merchandise/pb"
)

type CategoryTemplate merchandise.CategoryTemplate

func NewCategoryTemplate(id *VersionID, name string, description string) *CategoryTemplate {
	return &CategoryTemplate{
		Id:             (*common.VersionID)(id),
		Name:           name,
		Description:    description,
		AttributeRules: make([]*merchandise.CategoryTemplateAttributeRule, 0),
	}
}

func (template *CategoryTemplate) addCategoryTemplateAttributeRule(attribute *CategoryTemplateAttributeRule) {
	template.AttributeRules = append(template.AttributeRules, (*merchandise.CategoryTemplateAttributeRule)(attribute))
}

// Template Attribute Rules

type CategoryTemplateAttributeRule merchandise.CategoryTemplateAttributeRule

func NewCategoryTemplateAttributeRule(t merchandise.CategoryTemplateAttributeRule_FieldType, ordinal int32, required bool, allowOverride bool, validationRegEx string) *CategoryTemplateAttributeRule {
	return &CategoryTemplateAttributeRule{
		FieldType:       t,
		Ordinal:         ordinal,
		Required:        required,
		AllowOverride:   allowOverride,
		ValidationRegex: validationRegEx,
	}
}
