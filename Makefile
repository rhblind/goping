# Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=inctl

all: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
test:
	gotestsum --format testname
lint:
	golangci-lint run
gosec:
	gosec ./...
govulncheck:
	govulncheck ./...
security: gosec govulncheck

check: build test lint security clean
