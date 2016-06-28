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

golint

go fmt $(go list ./...)

go generate $(go list ./... | grep -v /vendor/)

go test $(go list ./... | grep -v '/vendor/' | grep -v '/builtin/bins/')  -timeout=30s -parallel=4 -bench=. -benchmem -cover

#go test ./chapter5 -bench=. -benchmem

go tool vet -all .

echo "golang build"

go build -ldflags "-X main.Build=${VERSION} -X main.Revision=${REV}" -v -o tutorial .

echo "DONE"

exit 0

