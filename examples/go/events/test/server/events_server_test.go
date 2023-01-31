package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	assert.NotNil(t, server)
	assert.Equal(t, "localhost", server.Host)
	assert.Equal(t, int32(50051), server.Port)
}
