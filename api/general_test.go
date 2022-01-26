package api

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	actual := NewConfig()
	fmt.Println(actual.LogLevel)
}
