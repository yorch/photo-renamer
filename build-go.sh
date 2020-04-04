#!/bin/sh

SUM_FILE=go.sum

rm -f \
    ${SUM_FILE} \
    src/${SUM_FILE}

cd src/

go build -o renamer main.go
