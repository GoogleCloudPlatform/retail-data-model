package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Address common.Address
type VersionID common.VersionID

func NewVersionID() *VersionID {
	u, _ := uuid.NewRandom()
	return &VersionID{
		Id:        u.String(),
		Version:   0,
		Created:   timestamppb.Now(),
		Effective: timestamppb.Now(),
	}
}

func NewAddress(
	geoSegmentId string,
	line1 string,
	line2 string,
	line3 string,
	line4 string,
	city string,
	territory string,
	postalCode string,
	countrySubdivision string) *Address {
	return &Address{
		GeoSegmentId:                     geoSegmentId,
		Line_1:                           line1,
		Line_2:                           line2,
		Line_3:                           line3,
		Line_4:                           line4,
		City:                             city,
		Territory:                        territory,
		PostalCode:                       postalCode,
		Iso_3166_2CountrySubDivisionCode: countrySubdivision,
	}
}
