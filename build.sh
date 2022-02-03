#!/bin/bash

set -o xtrace

BIN="go"
which go1.18beta2
if [ $? == 0 ]; then
  BIN=go1.18beta2
fi

set -e

PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "${PROJECT_DIR}"

VERSION=`git describe --tags --always`

GOOS=windows GOARCH=386 $BIN build -tags "vcs.describe=${VERSION}" -o "bin/rstats-win32-${VERSION}.exe" ./cmd/rstats
GOOS=windows GOARCH=amd64 $BIN build -tags "vcs.describe=${VERSION}" -o "bin/rstats-win64-${VERSION}.exe" ./cmd/rstats
GOOS=darwin GOARCH=amd64 $BIN build -tags "vcs.describe=${VERSION}" -o "bin/rstats-macos-${VERSION}" ./cmd/rstats
GOOS=linux GOARCH=amd64 $BIN build -tags "vcs.describe=${VERSION}" -o "bin/rstats-linux-${VERSION}" ./cmd/rstats
