#!/bin/bash

echo "Compiling started..."

go build main.go

echo "Compiling complete."

echo "Trying to launch program"

go run main.go

echo "Program exited"
