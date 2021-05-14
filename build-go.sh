#!/bin/sh

SUM_FILE=go.sum

BUILD_DIR=build
SRC_DIR=src

rm -rf ${BUILD_DIR}

cd ${SRC_DIR}

go build -o ../${BUILD_DIR}/renamer main.go
