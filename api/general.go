package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Config struct {
	Port           int           `json:"port"`
	SocketPort     int           `json:"socket-port"`
	RedirPort      int           `json:"redir-port"`
	TproxyPort     int           `json:"tproxy-port"`
	MixedPort      int           `json:"mixed-port"`
	Authentication []interface{} `json:"authentication"`
	AllowLAN       bool          `json:"allow-lan"`
	BindAddress    string        `json:"bind-address"`
	Mode           string        `json:"mode"`
	LogLevel       string        `json:"log-level"`
	Ipv6           bool          `json:"ipv6"`
}

// 基础配置构造函数
func NewConfig() *Config {
	config, err := GetConfigs()
	if err != nil {
		return &Config{}
	}
	return config
}

// 获取当前基础设置
func GetConfigs() (*Config, error) {
	resp, err := http.Get(BaseUrl + "/configs")
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	var config Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// PatchConfigs 增量修改clash配置
func PatchConfigs(key string, value interface{}) error {
	res := true
	switch key {
	case "port":
	case "socks-port":
	case "redir-port":
	case "mode":
	case "log-level":
	case "allow-lan":
		if v, ok := value.(bool); ok {
			config := struct {
				AllowLan bool `json:"allow-lan"`
			}{v}
			res = patch(config)
		}
	default:
		log.Println("invalid setting item.")
	}
	if !res {
		return errors.New("set " + key + " failed.")
	} else {
		return nil
	}
}

func patch(c interface{}) bool {
	config, err := json.Marshal(c)
	if err != nil {
		log.Println("marshal config failed.")
		return false
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPatch, BaseUrl+"/configs", bytes.NewBuffer(config))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	return resp.StatusCode == 204
}
