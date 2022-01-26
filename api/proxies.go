package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Proxy struct {
	All  []string `json:"all"`
	Name string   `json:"name"`
	Now  string   `json:"now"`
	Type string   `json:"type"`
	UDP  bool     `json:"udp"`
}

type Proxies struct {
	Proxies map[string]Proxy `json:"proxies"`
}

type Delay struct {
	Delay int `json:"delay"`
}

type Message struct {
	Message string `json:"message"`
}

type Name struct {
	Name string `json:"name"`
}

// 获取所有代理
func GetProxies() Proxy {
	resp, err := http.Get(BaseUrl + "/proxies")
	if err != nil {
		log.Print("http.Get error.")
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Print("io.ReadAll error.")
	}

	var proxies Proxies
	err = json.Unmarshal(body, &proxies)
	if err != nil {
		log.Fatal(err)
	}

	if proxy, ok := proxies.Proxies["Proxy"]; ok {
		return proxy
	}
	return Proxy{}
}

// 获取单个代理信息
func GetProxyInfoByName(name string) string {
	resp, err := http.Get(BaseUrl + "/proxies/" + name)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

// 获取单个代理的延迟
func GetProxyDelayByName(name string) string {
	query := "timeout=" + TimeOut + "&url=" + URL
	resp, err := http.Get(BaseUrl + "/proxies/" + name + "/delay?" + query)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		var delay Delay
		if err := json.Unmarshal(body, &delay); err != nil {
			delay.Delay = -1
		}
		return strconv.Itoa(delay.Delay)
	} else {
		var message Message
		json.Unmarshal(body, &message)
		return message.Message
	}
}

// 切换Selector中选中的代理
func ChangeProxyByName(name string) bool {
	b := Name{name}
	json, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, BaseUrl+"/proxies/Proxy", bytes.NewBuffer(json))

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode == 204 {
		fmt.Println(string(body))
		return true
	} else {
		fmt.Println(string(body))
		return false
	}
}
