package api

import "testing"

func TestAutoStart(t *testing.T) {
	if err := AutoStart(true); err != nil {
		t.Error(err)
	}
}
