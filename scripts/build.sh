#!/bin/bash

# This script builds the application from source for multiple platforms.
set -e
set -x

VERSION=${VERSION:-(development)}
BINARY=${BINARY:-delmo}

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

# Delete the old dir
echo "==> Removing old directory..."
rm -f bin/*
rm -rf pkg/*
mkdir -p bin

TARGETS=${TARGETS:-linux/amd64 darwin/amd64}
# If its dev mode, only build for ourself
if [[ ! -z "${DELMO_DEV}" ]]; then
  TARGETS=$(go env GOOS)/$(go env GOARCH)
fi


# Build!
echo "==> Building..."
gox -osarch="${TARGETS}" --output="pkg/${BINARY}-{{.OS}}-{{.Arch}}" -ldflags="-X main.Version=${VERSION}" ./...

ls pkg

DEV_BINARY="${BINARY}-$(go env GOOS)-$(go env GOARCH)"
cp pkg/${DEV_BINARY} bin/

echo "Built version: $(bin/${DEV_BINARY} --version)"
