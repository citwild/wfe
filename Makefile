pkgs = $(shell go list ./... | grep -v vendor)

all: presubmit build test

generate:
	@echo ">> generating code"
	@go generate $(pkgs)

format:
	@echo ">> formatting code"
	@go fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@go vet $(pkgs)

presubmit: vet
	@echo ">> checking go formatting"
	@./build/check_gofmt.sh

assets:
	@echo ">> building assets"
	@./build/assets.sh

build: assets
	@echo ">> building binaries"
	@./build/build.sh

test:
	@echo ">> running tests"
	@go test -short -race $(pkgs)

test-long: build
	@echo ">> running tests long"
	@go test -race $(pkgs)

release:
	@echo ">> building release binaries"
	@./build/release.sh $(VERSION)

docker:
	@echo ">> building docker image"
	@docker build -t wfe:$(shell git rev-parse --short HEAD) -f deploy/Dockerfile .

clean:
	@echo ">> removing compiled files"
	@go clean -i $(pkgs)

.PHONY: all generate format vet presubmit assets build test test-long release docker clean
