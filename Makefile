.PHONY: all install test integration-test

TAG=$(shell git describe --abbrev=0 --tags 2>&1)
TS=$(shell date '+%b %d %Y %T')

all: help

# Self documenting makefile. Double hashes signify help comments.
help:                   ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

install:                ## Installs tables-to-go. Requires `git` to be installed.
	@go install -mod=vendor -ldflags \
    	"-X 'main.buildTimestamp=$(TS)' -X 'main.versionTag=$(TAG)'" \
    	.

test:                   ## Runs unit tests with race flag enabled
	go test -mod=vendor -race ./...

integration-test:       ## Runs integration tests
	go test -mod=vendor -tags=integration ./internal/integration_tests/...
