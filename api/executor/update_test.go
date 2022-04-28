package executor

import "testing"

func TestDownloadConfig(t *testing.T) {
	err := DownloadConfig("https://stc-anycast.net/link/vFN5GEw7JDDmggpH?sub=3&client=clash")
	if err != nil {
		t.Error("test fail.")
	}
}
