package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type User struct {
	Name     string    `yaml:"name"`
	Age      int       `yaml:"age"`
	Comments []Comment `yaml:"comments"`
}

type Comment struct {
	Text string `yaml:"text"`
}

func writeYAML(filePath string, data []User) error {
	dir := filepath.Dir(filePath)

	if err := os.MkdirAll(dir, os.FileMode(0755)); err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}
	defer file.Close()

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}

	_, err = file.Write(yamlData)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %s", err.Error())
	}

	return nil
}
