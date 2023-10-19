package model

import (
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/retail-data-model/enums"
	"github.com/stretchr/testify/assert"
	"testing"
)

const NAME = "name"
const DEPARTMENT = "department"
const LOGICAL = "logical"
const PUBLIC = "public"

func TestLocation(t *testing.T) {

	cm := enums.Distance_CENTIMETER

	// Places the left back corner of the site as 0,0,0
	// Create a building 49x25 meters, with the support of 10 meter ceilings implied
	floorCoordinates := NewLocationCoordinate(0, cm, 0, 0, 0).
		AddVertex(0, -2500, 0).
		AddVertex(4900, -2500, 0).
		AddVertex(4900, 0, 0)

	store0101 := NewLocation(NewVersionID(),
		"STORE",
		"SELLING_LOCATION",
		PUBLIC,
		false,
		floorCoordinates)

	dairy := NewLocation(NewVersionID(),
		DEPARTMENT, LOGICAL, PUBLIC, false,
		NewLocationCoordinate(0, cm, 0, 0, 0).
			AddVertex(800, 0, 0).
			AddVertex(800, -600, 0).
			AddVertex(1400, -600, 0).
			AddVertex(400, -600, 0).
			AddVertex(400, -1100, 0).
			AddVertex(0, -1100, 0)).AddMeta("name", "Dairy")

	frozen := NewLocation(NewVersionID(), DEPARTMENT, LOGICAL, PUBLIC, false,
		NewLocationCoordinate(0, cm, 400, -601, 0).
			AddVertex(1400, -601, 0).
			AddVertex(1400, -1500, 0).
			AddVertex(400, -1500, 0)).AddMeta("name", "Frozen Food")

	coldCuts := NewLocation(NewVersionID(), DEPARTMENT, LOGICAL, PUBLIC, false,
		NewLocationCoordinate(0, cm, 0, -1101, 0).
			AddVertex(400, -1101, 0).
			AddVertex(400, -1400, 0).
			AddVertex(0, -1400, 0)).AddMeta(NAME, "Cold Cuts")

	produce := NewLocation(NewVersionID(), DEPARTMENT, LOGICAL, PUBLIC, false,
		NewLocationCoordinate(0, cm, 0, -1401, 0).
			AddVertex(1400, -1401, 0).
			AddVertex(1400, -1900, 0).
			AddVertex(0, -1900, 0)).AddMeta(NAME, "Produce")

	restRooms := NewLocation(NewVersionID(), PUBLIC, LOGICAL, PUBLIC, false,
		NewLocationCoordinate(0, cm, 0, -1901, 0).
			AddVertex(999, -1901, 0).
			AddVertex(999, -2500, 0).
			AddVertex(0, -2500, 0)).AddMeta(NAME, "Public Lavatories")

	checkout := NewLocation(NewVersionID(), "CONTROL", LOGICAL, PUBLIC, false,
		NewLocationCoordinate(0, cm, 1000, -1901, 0).
			AddVertex(3000, -1901, 0).
			AddVertex(3000, -2500, 0).
			AddVertex(1000, -2500, 0))

	dairy.ParentId = store0101.Id
	frozen.ParentId = store0101.Id
	coldCuts.ParentId = store0101.Id
	produce.ParentId = store0101.Id
	restRooms.ParentId = store0101.Id
	checkout.ParentId = store0101.Id

	assert.NotNil(t, store0101)

	b, err := json.Marshal(store0101)
	assert.Nil(t, err, "Error was not nil")
	fmt.Println(string(b))

}
