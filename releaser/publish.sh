#!/usr/bin/env bash

set -e

VERSION=$(git describe --tags --abbrev=0)

echo ""
echo " Publish Goker $VERSION"
echo ""

DOCKER_IMAGE="skribisto/goker"
DOCKER_TAG=$VERSION

docker buildx build --progress=plain -t $DOCKER_IMAGE:$VERSION -t $DOCKER_IMAGE:latest --push .