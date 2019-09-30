#!/bin/bash
set -e

make build-local

echo "Generating new CLI: greet"
cli-generator-dev generate --project-dir ../greet --language go --spec ./examples/greet.yaml --go-package github.com/greet/greet --create-handlers

cd ../greet

echo "Runnig go mod init & go mod tidy"
go mod init github.com/greet/greet
go mod tidy

go run main.go welcome --array test1 --array test2

echo "Done"