package main

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func getYAML(cfg []Config) (string, error) {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
