MAIN=./cmd/nofox
TAG_VERSION=$(shell git describe --tags --exact-match 2>/dev/null || git rev-parse --short HEAD)
VERSION=$(subst /,_,$(TAG_VERSION))

all: build
	@echo all built

build:
	@go build $(MAIN)

clean:
	@rm -rf nofox target
	@echo all removed

lint:
	@golangci-lint run ./...
	@echo all code is linted

format:
	@gofmt -w -s .

format-check:
	@gofmt -l .

run:
	@go run $(MAIN)

test:
	@go test -v ./...

test-cov:
	@mkdir -p target/test
	@go test -v -cover -coverprofile=./target/test/coverage.out ./...

sbom: build
	@mkdir -p target/sbom
	@cyclonedx-gomod bin -json -output ./target/sbom/nofox-$(VERSION).bom.json ./nofox

.PHONY: all build clean lint format format-check run test test-cov sbom
