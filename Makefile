pkgs = $(shell go list ./... | grep -v vendor)

all: test

get-deps:
	@echo ">> getting dependencies"
	@go get -u ./...

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
	@echo ">> installing binaries"
	@go install ./cmd/wfe

docker:
	@echo ">> building docker image"
	@go build -o ./deploy/wfe ./cmd/wfe
	@docker build -t wfe -f deploy/Dockerfile deploy

.PHONY: all get-deps format vet generate test test-long install docker
