#!/bin/bash
set -e
set -o pipefail

REV="$(git rev-parse --short HEAD)"
VERSION="$1"
if [ -z "$1" ]
  then
    VERSION=0.0.0
fi

echo "revision: " $REV "version: " ${VERSION}

echo "testing project - chapter5"

go test ./chapter5

echo "golang build"

go build -ldflags "-X main.Build=${VERSION} -X main.Revision=${REV}" -o tutorial .

echo "DONE"

exit 0

