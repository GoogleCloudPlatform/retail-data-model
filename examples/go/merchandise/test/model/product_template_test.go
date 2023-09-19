package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductTemplate(t *testing.T) {

	template := NewProductTemplate(NewVersionID(), "Hard lines", "Durable Goods")

	assert.NotNil(t, template)
}
