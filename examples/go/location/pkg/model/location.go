package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/enums"
	location "github.com/GoogleCloudPlatform/retail-data-model/location/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type LocationMeasure location.LocationMeasure
type Location location.Location

func (loc *Location) AddMerchandiseGroup(merchGroup *MerchantGroup) *Location {
	loc.MerchandiseGroups = append(loc.MerchandiseGroups, (*location.Location_MerchGroup)(merchGroup))
	return loc
}

func (loc *Location) AddLocationMeasure(measure *LocationMeasure) *Location {
	loc.Measures = append(loc.Measures, (*location.LocationMeasure)(measure))
	return loc
}

func (loc *Location) AddMeta(key string, value string) *Location {
	loc.Meta = append(loc.Meta, &common.BusinessKey{Name: key, Value: value})
	return loc
}

type MerchantGroup location.Location_MerchGroup
type LocationCoordinate location.LocationCoordinate

func (lc *LocationCoordinate) AddVertexWithOrdinal(vertex *LocationVertex) *LocationCoordinate {
	lc.LocationVertices = append(lc.LocationVertices, (*location.LocationCoordinate_LocationVertex)(vertex))
	return lc
}

func (lc *LocationCoordinate) AddVertex(x float64, y float64, z float64) *LocationCoordinate {
	lc.AddVertexWithOrdinal(NewLocationVertex(int32(len(lc.LocationVertices)), x, y, z))
	return lc
}

type LocationVertex location.LocationCoordinate_LocationVertex

type Distance location.LocationMeasure_Distance
type Count location.LocationMeasure_Count
type Capacity location.LocationMeasure_Capacity
type Area location.LocationMeasure_Area
type Weight location.LocationMeasure_Weight

func NewDistance(
	name string,
	uom *Distance,
	value float64) *LocationMeasure {
	return &LocationMeasure{
		Name:  name,
		Uom:   (*location.LocationMeasure_Distance)(uom),
		Value: value,
	}
}

func NewCount(name string, count *Count, value float64) *LocationMeasure {
	return &LocationMeasure{
		Name:  name,
		Uom:   (*location.LocationMeasure_Count)(count),
		Value: value,
	}
}

func NewCapacity(name string, capacity *Capacity, value float64) *LocationMeasure {
	return &LocationMeasure{
		Name:  name,
		Uom:   (*location.LocationMeasure_Capacity)(capacity),
		Value: value,
	}
}

func NewArea(name string, area *Area, value float64) *LocationMeasure {
	return &LocationMeasure{
		Name:  name,
		Uom:   (*location.LocationMeasure_Area)(area),
		Value: value,
	}
}

func NewWeigh(name string, weight *Weight, value float64) *LocationMeasure {
	return &LocationMeasure{
		Name:  name,
		Uom:   (*location.LocationMeasure_Weight)(weight),
		Value: value,
	}
}

func NewMerchGroup(merchandiseHierarchyId string, effectiveDate time.Time, performanceScore float64) *MerchantGroup {
	return &MerchantGroup{
		MerchandiseHierarchyId: merchandiseHierarchyId,
		EffectiveDate:          timestamppb.New(effectiveDate),
		PerformanceScore:       performanceScore,
	}
}

func NewLocationVertex(
	ordinal int32,
	x float64,
	y float64,
	z float64) *LocationVertex {
	return &LocationVertex{
		Ordinal: ordinal,
		X:       x,
		Y:       y,
		Z:       z,
	}
}

func NewLocationCoordinate(
	northOffset float64,
	uom enums.Distance,
	x float64,
	y float64,
	z float64) *LocationCoordinate {
	return &LocationCoordinate{
		AngleOffsetDegrees: northOffset,
		UnitOfMeasure:      uom,
		X:                  x,
		Y:                  y,
		Z:                  z,
		LocationVertices:   make([]*location.LocationCoordinate_LocationVertex, 0),
	}
}

func NewLocation(
	id *VersionID,
	locationType string,
	locationFunctionType string,
	inventoryLocationSecurityType string,
	stockLedgerControlFlag bool,
	locationCoordinate *LocationCoordinate,
) *Location {
	return &Location{
		Id:                            (*common.VersionID)(id),
		ParentId:                      nil,
		LocationType:                  locationType,
		LocationFunctionType:          locationFunctionType,
		InventoryLocationSecurityType: inventoryLocationSecurityType,
		StockLedgerControlFlag:        stockLedgerControlFlag,
		LocationCoordinate:            (*location.LocationCoordinate)(locationCoordinate),
		MerchandiseGroups:             make([]*location.Location_MerchGroup, 0),
		Meta:                          make([]*common.BusinessKey, 0),
	}
}
