package main

import "gopkg.in/yaml.v2"

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

func getConfigFromYAML(data []byte) (Config, error) {
	res := Config{}
	if err := yaml.Unmarshal(data, &res); err != nil {
		return Config{}, err
	}
	return res, nil
}
