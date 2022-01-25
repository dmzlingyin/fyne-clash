package executor

import (
	"testing"
)

func TestParse(t *testing.T) {
	_, _, err := Parse()
	if err != nil {
		t.Error(err)
	}
}
