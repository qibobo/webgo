package config

import (
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Port int `yaml:"port"`
}

func LoadConfig(bytes []byte) (*Config, error) {
	conf := &Config{
		Port: 8080,
	}
	err := yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *Config) Validate() error {
	return nil
}
