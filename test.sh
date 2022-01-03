#!/bin/bash

set -o xtrace
set -e

PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "${PROJECT_DIR}"

BIN="go"
which go1.18beta1
if [ $? == 0 ]; then
  BIN=go1.18beta1
fi

$BIN test ./...
