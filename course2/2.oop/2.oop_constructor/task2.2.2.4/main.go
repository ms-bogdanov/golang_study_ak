package main

import (
	"fmt"
	"io"
	"os"
)

type Logger interface {
	Log(message string)
}

type FileLogger struct {
	file *os.File
}

func (f FileLogger) Log(message string) {
	fmt.Fprintln(f.file, message)
}

type ConsoleLogger struct {
	out io.Writer
}

func (c ConsoleLogger) Log(message string) {
	fmt.Fprintln(c.out, message)
}

type LogSystem struct {
	logger Logger
}

type LogOption func(*LogSystem)

func WithLogger(logger Logger) LogOption {
	return func(ls *LogSystem) {
		ls.logger = logger
	}
}

func NewLogSystem(opts ...LogOption) *LogSystem {
	logSystem := &LogSystem{}
	for _, opt := range opts {
		opt(logSystem)
	}
	return logSystem
}

func (ls *LogSystem) Log(message string) {
	if ls.logger != nil {
		ls.logger.Log(message)
	}
}
