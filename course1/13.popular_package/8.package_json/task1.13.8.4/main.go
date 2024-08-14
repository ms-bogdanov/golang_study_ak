package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	data := []map[string]interface{}{
		{
			"name": "Elliot",
			"age":  25,
		},
		{
			"name": "Fraser",
			"age":  30,
		},
	}
	err := writeJSON("users.json", data)
	if err != nil {
		panic(err)
	}
}

func writeJSON(filePath string, data interface{}) error {
	dir := filepath.Dir(filePath)

	if err := os.MkdirAll(dir, os.FileMode(0755)); err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}

	return nil
}
