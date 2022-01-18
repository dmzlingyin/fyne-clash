package api

import (
	"encoding/json"
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

var (
	BaseUrl = "http://localhost:9090"
	TimeOut = "5000"
	URL     = "http://www.gstatic.com/generate_204"
)

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
