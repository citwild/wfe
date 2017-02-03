pkgs = $(shell go list ./... | grep -v vendor)

all: get-deps test

get-deps:
	@echo ">> getting dependencies"
	@go get -u ./...
	@go get github.com/shurcooL/vfsgen

format:
	@echo ">> formatting code"
	@go fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@go vet $(pkgs)

generate:
	@echo ">> generating code"
	@go generate $(pkgs)

test:
	@echo ">> running tests"
	@go test -short -race $(pkgs)

test-long: install
	@echo ">> running tests long"
	@go test -race $(pkgs)

install:
	@echo ">> installing binary"
	@go install ./cmd/wfe

.PHONY: all get-deps format vet generate test test-long install
