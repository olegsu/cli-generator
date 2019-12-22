#!/bin/bash
set -e
echo "Generating CLI"
cli-generator generate --project-dir . --language go --spec ./build/cli-generator.yaml  --go-package github.com/olegsu/cli-generator

echo "Generating template map from template files"
go generate ${PWD}/scripts/generate.go
echo "All files geneated"
OUTFILE=/usr/local/bin/cli-generator-dev
echo "Building go binary"
go build -o $OUTFILE *.go

chmod +x $OUTFILE