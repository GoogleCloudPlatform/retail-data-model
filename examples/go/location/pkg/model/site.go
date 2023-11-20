package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/enums"
	location "github.com/GoogleCloudPlatform/retail-data-model/location/pb"
)

type Site location.Site
type PhysicalLocation location.Site_PhysicalLocation

func NewPhysicalLocation(
	address *Address,
	climate enums.Climate,
	latDeg int32,
	latMin int32,
	latSec int32,
	latDirCode string,
	lngDeg int32,
	lngMin int32,
	lngSec int32,
	lngDirCode string) *PhysicalLocation {

	return &PhysicalLocation{
		Address:                (*common.Address)(address),
		Climate:                climate,
		LatitudeDegrees:        latDeg,
		LatitudeMinutes:        latMin,
		LatitudeSeconds:        latSec,
		LatitudeDirectionCode:  latDirCode,
		LongitudeDegrees:       lngDeg,
		LongitudeMinutes:       lngMin,
		LongitudeSeconds:       lngSec,
		LongitudeDirectionCode: lngDirCode,
	}
}

func NewSite(id *VersionID,
	siteType string,
	physicalLocation *PhysicalLocation,
	operationalPartyId string,
	icaoCode string,
	timeZone enums.TimeZone,
	angleOffsetFromNorth float64) *Site {
	return &Site{
		Id:                 (*common.VersionID)(id),
		SiteType:           siteType,
		Location:           (*location.Site_PhysicalLocation)(physicalLocation),
		OperationalPartyId: operationalPartyId,
		IcaoCode:           icaoCode,
		TimeZone:           timeZone,
		AngleOffsetDegrees: angleOffsetFromNorth,
		Contacts:           make([]*common.Contact, 0),
		Meta:               make([]*common.BusinessKey, 0),
		LocationIds:        make([]string, 0),
	}
}

func (site *Site) addContact(contact *common.Contact) *Site {
	site.Contacts = append(site.Contacts, contact)
	return site
}

func (site *Site) AddLocationId(locationId string) *Site {
	site.LocationIds = append(site.LocationIds, locationId)
	return site
}

func (site *Site) AddMeta(key string, value string) *Site {
	site.Meta = append(site.Meta, &common.BusinessKey{Name: key, Value: value})
	return site
}
