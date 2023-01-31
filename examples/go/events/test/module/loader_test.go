package module

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoader(t *testing.T) {
	assert.NotNil(t, eventModule)
}
