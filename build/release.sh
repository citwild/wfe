#!/usr/bin/env bash
set -e

VERSION=${1:-$(git rev-parse --short HEAD)}
GO_FLAGS=${GO_FLAGS:-}

export GO_CMD="build"
export GO_FLAGS="-a -tags release ${GO_FLAGS}"
build/build.sh

echo ">> building wfe docker image"
docker build -t wfe:$VERSION -f deploy/Dockerfile .