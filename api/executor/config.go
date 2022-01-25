package executor

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type General struct {
	SystemProxy bool `yaml:"system-proxy"`
	AutoStart   bool `yaml:"auto-start"`
}

func Parse() (bool, bool, error) {
	file, err := os.Open("../config.yaml")
	if err != nil {
		return false, false, err
	}
	defer file.Close()

	config, err := io.ReadAll(file)
	if err != nil {
		return false, false, err
	}

	var g General
	err = yaml.Unmarshal(config, &g)
	if err != nil {
		return false, false, err
	}

	return g.SystemProxy, g.AutoStart, nil
}

func UpdateConfig(checked bool, field string) error {

	return nil
}
