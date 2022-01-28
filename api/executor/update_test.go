package executor

import "testing"

func TestDownloadConfig(t *testing.T) {
	err := DownloadConfig("https://www.jafiyun.fun/link/oR4pib4Kggdk1Ik0?clash=1")
	if err != nil {
		t.Error("test fail.")
	}
}
