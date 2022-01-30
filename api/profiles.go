package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Path struct {
	Path string `json:"path"`
}

// 重新加载配置文件
func Reload() error {
	path := Path{"/etc/clash/config.yaml"}
	json, err := json.Marshal(path)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, BaseUrl+"/configs?force=true", bytes.NewBuffer(json))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()

	if resp.StatusCode == 204 {
		return nil
	} else {
		return err
	}
}
