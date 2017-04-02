#!/usr/bin/env bash
set -e

GO_CMD=${GO_CMD:-"install"}
GO_FLAGS=${GO_FLAGS:-}

echo ">> building wfe"

go "$GO_CMD" ${GO_FLAGS} -ldflags "-s" "github.com/citwild/wfe/cmd/wfe"