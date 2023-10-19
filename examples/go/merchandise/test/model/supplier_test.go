package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSupplier(t *testing.T) {
	id := NewVersionID()

	supplier := NewSupplier(id, "01010101")

	assert.NotNil(t, supplier)
	assert.Equal(t, id.Id, supplier.Id.Id)
}
