package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Timeout   int    `yaml:"timeout"`
	Separator string `yaml:"separator"`
}

// ReadConfig returns a structure with data from config-file
func ReadConfig() (*Config, error) {
	// read config file
	configData, err := ioutil.ReadFile("./configs/config.yaml")
	if err != nil {
		// log.error
		return nil, err
	}
	// decode config
	cfg := new(Config)
	err = yaml.Unmarshal(configData, cfg)
	if err != nil {
		// log.error
		return nil, err
	}
	return cfg, nil
}
