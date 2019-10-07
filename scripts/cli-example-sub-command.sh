#!/bin/bash
set -e

make build-local

echo "Generating new CLI: cli-example-sub-command"
cli-generator-dev generate --project-dir cli-example-sub-command --language go --spec ./examples/cli-example-sub-command.yaml --go-package github.com/cli-example-sub-command/cli-example-sub-command --create-handlers --run-init-flow --verbose

cd cli-example-sub-command

echo "Runnig go mod init & go mod tidy"
go mod init github.com/cli-example-sub-command/cli-example-sub-command
go mod tidy

echo "Running make build"
make build
echo "Running compiled binary"
./cli-example-sub-command cmd sub

echo "Done"