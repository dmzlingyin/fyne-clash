/*
系统代理实现部分. 根据Qv2ray官网介绍, KDE的系统配置文件可能并不会被相应的应用程序读取, 因此可能达不到预期的效果.
配置文件路径: ~/.config/kioslaverc

ProxyType的取值分别代表不同的代理配置：
		0 无代理
		1 使用手动配置的代理服务器
		2 使用代理自动配置 URL
		3 自动检测代理配置
		4 使用系统代理服务器配置
*/
package executor

import (
	"log"
	"strconv"

	"gopkg.in/ini.v1"
)

var (
	path    = "/home/lingyin/.config/kioslaverc"
	address = "http://127.0.0.1"
	port    = "7890"
)

const (
	NoProxy = iota
	ManualProxy
	AutoConfig
	AutoDetect
	SystemProxy
)

func IsSystemProxy() bool {
	sec := ReadFile()
	pat, err := sec.Key("ProxyType").Int()
	if err != nil {
		log.Println(err)
	}
	if pat == SystemProxy {
		return true
	} else {
		return false
	}
}

func ReadFile() *ini.Section {
	cfg, err := ini.Load(path)
	if err != nil {
		log.Println(err)
	}
	return cfg.Section("Proxy Settings")
}

func SetProxy(proxyType int) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return err
	}
	sec := cfg.Section("Proxy Settings")

	if proxyType == SystemProxy {
		sec.Key("ProxyType").SetValue(strconv.Itoa(SystemProxy))
		sec.Key("httpProxy").SetValue(address + ":" + port)
		sec.Key("httpsProxy").SetValue(address + ":" + port)
		sec.Key("ftpProxy").SetValue(address + ":" + port)
		sec.Key("socksProxy").SetValue(address + ":" + port)
		sec.Key("NoProxyFor").SetValue("127.0.0.1,localhost")
	} else {
		sec.Key("ProxyType").SetValue(strconv.Itoa(NoProxy))
	}

	cfg.SaveTo(path)
	return nil
}
