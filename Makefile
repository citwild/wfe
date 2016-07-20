pkgs = $(shell go list ./... | grep -v vendor)

all: build

get-deps:
	@echo ">> getting dependencies"
	@go get ./...

format:
	@echo ">> formatting code"
	@go fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@go vet $(pkgs)

generate:
	@echo ">> generating code"
	@go generate $(pkgs)

build:
	@echo ">> building binaries"
	@go build $(pkgs)

test:
	@echo ">> running tests"
	@go test -short -race $(pkgs)

test-long: install
	@echo ">> running tests long"
	@go test -race $(pkgs)

install:
	@echo ">> installing binaries"
	@go install ./cmd/wfe

precommit: vet
	@gofmt -s -l $(pkgs)

.PHONY: all get-deps format vet generate build test test-long install precommit
