package api

import "testing"

func TestGetProxies(t *testing.T) {
	proxy := GetProxies()
	if proxy.Name == "" {
		t.Error("get proxy name error, test fail.")
	}
}

func TestGetProxyInfoByName(t *testing.T) {
	name := "A-Vip1丨倍率2丨香港10丨 原生丨Netflix丨"
	actual := GetProxyInfoByName(name)
	if actual == "" {
		t.Error("get proxy info error, test fail.")
	}
}

func TestProxyDelayByName(t *testing.T) {
	name := "Vip1丨新加坡01 "
	actual := GetProxyDelayByName(name)
	if actual == "" {
		t.Error("get proxy delay error, test fail.")
	}
}
