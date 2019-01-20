package config

import (
	"time"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	baseConfigPath = "./config/base.yaml"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Hostport string        `yaml:"hostport"`
	Scheme   string        `yaml:"scheme"`
	Timeout  time.Duration `yaml:"timeout"`
}

func Load() (*Config, error) {
	configBytes, err := ioutil.ReadFile(baseConfigPath)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	if err := yaml.Unmarshal(configBytes, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
