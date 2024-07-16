#!/bin/bash
#chmod +x gofmt.sh

if [ -z "$1" ]; then
 echo "Usage: gofmt.sh "
 exit 1
fi

go fmt $1