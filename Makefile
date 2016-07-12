pkgs = $(shell go list ./... | grep -v vendor)

all: build test

format:
	go fmt $(pkgs)

generate:
	go generate $(pkgs)

build:
	go build $(pkgs)

test: 
	go test -race $(pkgs)

install:
	go install ./cmd/wfe

.PHONY: all format generate build test install
