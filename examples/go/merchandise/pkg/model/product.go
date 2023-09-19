package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/enums"
	merchandise "github.com/GoogleCloudPlatform/retail-data-model/merchandise/pb"
)

type I18nProductHeader merchandise.I18NProductHeader

func NewI18nProductHeader(locale Locale, brand string, name string, shortDescription string, longDescription string) *I18nProductHeader {
	return &I18nProductHeader{
		Locale:           (enums.Locale)(locale),
		Brand:            brand,
		Name:             name,
		ShortDescription: shortDescription,
		LongDescription:  longDescription,
		SearchTerms:      make([]string, 0),
		Images:           make([]*common.Image, 0),
	}
}

func (n *I18nProductHeader) AddSearchTerm(val string) {
	n.SearchTerms = append(n.SearchTerms, val)
}

func (n *I18nProductHeader) AddImage(image *Image) {
	n.Images = append(n.Images, (*common.Image)(image))
}

type Product merchandise.Product

type Variant merchandise.Product_Variant

func NewProductDetails() *merchandise.Product_Detail {
	return &merchandise.Product_Detail{
		Headers:    make([]*merchandise.I18NProductHeader, 0),
		Attributes: make([]*merchandise.ProductAttribute, 0),
	}
}

func NewProduct(header *I18nProductHeader, rating float64) *Product {
	product := &Product{
		Id:           (*common.VersionID)(NewVersionID()),
		BusinessKeys: make([]*common.BusinessKey, 0),
		Detail:       NewProductDetails(),
		Rating:       rating,
		Measurements: make([]*common.NamedMeasure, 0),
		Variations:   make([]*merchandise.Product_Variant, 0),
		Suppliers:    make([]*merchandise.Supplier, 0),
	}
	product.Detail.Headers = append(product.Detail.Headers, (*merchandise.I18NProductHeader)(header))
	return product
}

func NewVariant() *Variant {
	return &Variant{
		Details: NewProductDetails(),
	}
}

func (p *Product) AddBusinessKey(key string, values ...string) {
	p.BusinessKeys = append(p.BusinessKeys, (*common.BusinessKey)(NewBusinessKey(key, values)))
}

func (p *Product) AddHeader(header *I18nProductHeader) {
	p.Detail.Headers = append(p.Detail.Headers, (*merchandise.I18NProductHeader)(header))
}

func (p *Product) AddAttribute(attribute *ProductAttribute) {
	p.Detail.Attributes = append(p.Detail.Attributes, (*merchandise.ProductAttribute)(attribute))
}

func (p *Product) AddNamedMeasure(measure *NamedMeasure) {
	p.Measurements = append(p.Measurements, (*common.NamedMeasure)(measure))
}

func (p *Product) AddVariation(variant *Variant) {
	p.Variations = append(p.Variations, (*merchandise.Product_Variant)(variant))
}
