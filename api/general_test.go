package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	actual := NewConfig()
	fmt.Println(actual.LogLevel)
}

func TestPatchConfig(t *testing.T) {
	assert.Equal(t, true, PatchConfigs("allow-lan", true), "allow-lan test failed.")
	assert.Equal(t, true, PatchConfigs("allow-lan", false), "allow-lan test failed.")
}
