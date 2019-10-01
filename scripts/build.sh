#!/bin/bash
set -e
echo "Generating template map from template files"
go generate ${PWD}/scripts/generate.go
echo "All files geneated"
OUTFILE=/usr/local/bin/cli-generator-dev
echo "Building go binary"
go build -o $OUTFILE *.go

chmod +x $OUTFILE