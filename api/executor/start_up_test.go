package executor

import (
	"fmt"
	"testing"
)

func TestAutoStart(t *testing.T) {
	if err := AutoStart(true); err != nil {
		t.Error(err)
	}
}

func TestIsStartUp(t *testing.T) {
	val := IsStartUp()
	fmt.Println(val)
	// if !val {
	// 	t.Error("test fail.")
	// }
}
