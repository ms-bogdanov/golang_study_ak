package main

import (
	"io"
	"log"
	"os"
)

func ReadString(filePath string) string {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err.Error())
	}

	defer file.Close()

	r, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %s", err.Error())
	}

	return string(r)
}
