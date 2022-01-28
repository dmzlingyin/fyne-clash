package executor

import (
	"io"
	"net/http"
	"os"
)

const File = "/etc/clash/config.yaml"

func DownloadConfig(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return updateConfig(body)
}

func updateConfig(content []byte) error {
	return os.WriteFile(File, content, 0666)
}
