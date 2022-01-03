#!/bin/bash

set -o xtrace
set -e

PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "${PROJECT_DIR}"

VERSION=`git describe --tags --always`

go1.18beta1 build -o bin -tags "vcs.describe=${VERSION}" -x ./...
