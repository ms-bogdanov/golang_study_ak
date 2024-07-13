#!/bin/bash
#chmod +x debug.sh

echo "Debug started..."
go build -o myprogram $1
dlv exec myprogram
echo "Debug ended."