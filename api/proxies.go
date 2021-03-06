package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	ProxyType = "π θηΉιζ©"
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

// θ·εζζδ»£η
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
	if proxy, ok := proxies.Proxies[ProxyType]; ok {
		return proxy
	}
	return Proxy{}
}

// θ·εεδΈͺδ»£ηδΏ‘ζ―
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

// θ·εεδΈͺδ»£ηηε»ΆθΏ
func GetProxyDelayByName(name string, ch chan map[string]string) string {
	if ok := strings.HasSuffix(name, "\t"); ok {
		name = name[:len(name)-2]
	}

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
		ch <- map[string]string{name: strconv.Itoa(delay.Delay)}
		return strconv.Itoa(delay.Delay)
	} else {
		var message Message
		json.Unmarshal(body, &message)
		ch <- map[string]string{name: message.Message}
		return message.Message
	}
}

// εζ’SelectorδΈ­ιδΈ­ηδ»£η
func ChangeProxyByName(name string) bool {
	b := Name{name}
	json, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, BaseUrl+"/proxies/"+ProxyType, bytes.NewBuffer(json))

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	return resp.StatusCode == 204
}
