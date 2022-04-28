package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProxies(t *testing.T) {
	assert.NotEqual(t, "", GetProxies().Name, "proxy is empty.")
}

func TestGetProxyInfoByName(t *testing.T) {
	name := "🇭🇰 香港 B06V"
	assert.NotEqual(t, "", GetProxyInfoByName(name), "get proxy info failed.")
}

func TestChangeProxyByName(t *testing.T) {
	name := "🇭🇰 香港 B06V"
	assert.Equal(t, true, ChangeProxyByName(name), "change proxy failed.")
}
