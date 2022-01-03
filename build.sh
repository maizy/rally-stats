#!/bin/bash

set -o xtrace

BIN="go"
which go1.18beta1
if [ $? == 0 ]; then
  BIN=go1.18beta1
fi

set -e

PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "${PROJECT_DIR}"

VERSION=`git describe --tags --always`


$BIN build -o bin -tags "vcs.describe=${VERSION}" -x ./...
