package executor

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	val := IsSystemProxy()
	fmt.Println(val)

	err := SetProxy(0)
	if err != nil {
		t.Error(err)
	}
}
