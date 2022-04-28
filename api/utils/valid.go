package utils

import (
	"log"
	"strings"
)

func IsUrlValid(url string) bool {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		log.Println("Invalid url.")
		return false
	}
	if !strings.Contains(url, "clash") {
		log.Println("Invalid clash describe url.")
		return false
	}
	return true
}
