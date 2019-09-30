#!/bin/bash
set -e
OUTFILE=/usr/local/bin/cli-generator-dev
go build -o $OUTFILE *.go

chmod +x $OUTFILE