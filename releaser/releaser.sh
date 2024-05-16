#!/usr/bin/env bash

set -e

RELEASE_VERSION=$(git describe --tags --abbrev=0)
RELEASE_DIR="release"

rm -fr $RELEASE_DIR
mkdir $RELEASE_DIR

echo ""
echo "Building Goker for version $RELEASE_VERSION"
echo ""


go build -o $RELEASE_DIR/goker -ldflags "-X github.com/skribisto/goker/internals/common.versionInfo=$RELEASE_VERSION"