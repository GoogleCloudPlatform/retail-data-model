package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/enums"
	merchandise "github.com/GoogleCloudPlatform/retail-data-model/merchandise/pb"
)

type Product merchandise.Product
type ProductAttribute merchandise.ProductAttribute
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

func NewProduct(header *I18nProductHeader, rating float64) *Product {
	product := &Product{
		Id:              (*common.VersionID)(NewVersionID()),
		BusinessKeys:    make([]*common.BusinessKey, 0),
		Headers:         make([]*merchandise.I18NProductHeader, 0),
		Attributes:      make([]*merchandise.ProductAttribute, 0),
		Rating:          rating,
		Measurements:    make([]*common.NamedMeasure, 0),
		Variations:      make([]*common.VersionID, 0),
		Substitutions:   make([]*common.VersionID, 0),
		RelatedProducts: make([]*common.VersionID, 0),
		Suppliers:       make([]*merchandise.Supplier, 0),
	}
	product.Headers = append(product.Headers, (*merchandise.I18NProductHeader)(header))
	return product
}

func (p *Product) AddBusinessKey(key string, value string) {
	p.BusinessKeys = append(p.BusinessKeys, (*common.BusinessKey)(NewBusinessKey(key, value)))
}

func (p *Product) AddHeader(header *I18nProductHeader) {
	p.Headers = append(p.Headers, (*merchandise.I18NProductHeader)(header))
}

func (p *Product) AddAttribute(attribute *ProductAttribute) {
	p.Attributes = append(p.Attributes, (*merchandise.ProductAttribute)(attribute))
}

func (p *Product) AddNamedMeasure(measure *NamedMeasure) {
	p.Measurements = append(p.Measurements, (*common.NamedMeasure)(measure))
}

func (p *Product) AddVariation(id *VersionID) {
	p.Variations = append(p.Variations, (*common.VersionID)(id))
}

func (p *Product) AddSubstitute(id *VersionID) {
	p.Substitutions = append(p.Substitutions, (*common.VersionID)(id))
}

func (p *Product) AddRelatedProduct(id *VersionID) {
	p.RelatedProducts = append(p.RelatedProducts, (*common.VersionID)(id))
}

func (p *Product) AddSupplier(supplier *Supplier) {
	p.Suppliers = append(p.Suppliers, (*merchandise.Supplier)(supplier))
}
