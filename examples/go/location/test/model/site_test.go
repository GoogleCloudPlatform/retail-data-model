package model

import (
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/retail-data-model/enums"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSite(t *testing.T) {

	id := NewVersionID()

	pl := NewPhysicalLocation(
		NewAddress("10",
			"1234 Somewhere Street",
			"",
			"",
			"",
			"Canton",
			"Georgia",
			"30122",
			"US-GA"),
		enums.Climate_AM,
		34,
		272,
		817,
		"NORTH",
		-84,
		500,
		354, "EAST")

	site := NewSite(
		id,
		"STORE",
		pl, "OP_TEST_PARTY", "10", enums.TimeZone_UTCMinus0500EasternTime, -24.0)

	assert.NotNil(t, site)
	b, err := json.Marshal(site)
	assert.Nil(t, err)
	fmt.Println(string(b))

}
