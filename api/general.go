package api

import (
	"encoding/json"
	"io"
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
