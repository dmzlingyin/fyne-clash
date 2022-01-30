package api

import "testing"

func TestReload(t *testing.T) {
	err := Reload()
	if err != nil {
		t.Error("test fail.")
	}
}
