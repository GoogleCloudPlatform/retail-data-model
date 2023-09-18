package model

import (
	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/enums"
	location "github.com/GoogleCloudPlatform/retail-data-model/location/pb"
)

type GeoSegmentGroup location.GeoSegmentGroup

func (gs *GeoSegment) addGeoSegment(segment *GeoSegment) {
	gs.Children = append(gs.Children, (*location.GeoSegment)(segment))
}
func (gsGroup *GeoSegmentGroup) addGeoSegmentId(id string) {
	gsGroup.GeoSegmentIds = append(gsGroup.GeoSegmentIds, id)
}

type GeoSegment location.GeoSegment

func (gs *GeoSegment) addSite(site *Site) {
	gs.Sites = append(gs.Sites, (*location.Site)(site))
}

func NewGeoSegmentGroup(
	id *VersionID,
	group enums.GeoSegmentGroup) *GeoSegmentGroup {
	return &GeoSegmentGroup{
		Id:            (*common.VersionID)(id),
		Type:          group,
		GeoSegmentIds: make([]string, 0),
	}
}

func NewGeoSegment(id *VersionID, t enums.GeoSegment, abbr string, description string) *GeoSegment {
	return &GeoSegment{
		Id:           (*common.VersionID)(id),
		Type:         t,
		Abbreviation: abbr,
		Description:  description,
		Sites:        make([]*location.Site, 0),
		Children:     make([]*location.GeoSegment, 0),
	}
}
