package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	merchandise "github.com/GoogleCloudPlatform/retail-data-model/merchandise/pb"
)

type Supplier merchandise.Supplier

func NewSupplier(id *VersionID, glCode string) *Supplier {
	return &Supplier{
		Id:                (*common.VersionID)(id),
		GeneralLedgerCode: glCode}
}
