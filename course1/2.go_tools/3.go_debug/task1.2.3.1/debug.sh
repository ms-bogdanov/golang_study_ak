#!/bin/bash
#chmod +x debug.sh

echo "Debug started..."
dlv debug main.go
echo "Debug ended."