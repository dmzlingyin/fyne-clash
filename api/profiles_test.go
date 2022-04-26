package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReload(t *testing.T) {
	assert.Equal(t, nil, Reload(), "Reload failed.")
}
