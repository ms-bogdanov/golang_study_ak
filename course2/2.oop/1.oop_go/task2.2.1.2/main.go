package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger interface {
	Log(message string) error
}

type ConsoleLogger struct {
	Writer io.Writer
}

func (cl ConsoleLogger) Log(message string) error {
	_, err := cl.Writer.Write([]byte(message))
	if err != nil {
		return err
	}

	return nil
}

type FileLogger struct {
	File *os.File
}

func (fl FileLogger) Log(message string) error {
	_, err := fl.File.Write([]byte(message))
	if err != nil {
		return err
	}

	return nil
}

type RemoteLogger struct {
	Address string
}

func (rl RemoteLogger) Log(message string) error {
	fmt.Printf("Logging to remote server at %s: %s\n", rl.Address, message)

	return nil
}

func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		err := logger.Log(message)
		if err != nil {
			log.Println("Failed to log message:", err)
		}
	}
}

func main() {
	consoleLogger := &ConsoleLogger{Writer: os.Stdout}
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer file.Close()
	fileLogger := &FileLogger{File: file}
	remoteLogger := &RemoteLogger{Address: "http://example.com/log"}

	loggers := []Logger{consoleLogger, fileLogger, remoteLogger}
	LogAll(loggers, "This is a test log message.")
}
