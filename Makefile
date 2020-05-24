.PHONY: all install test

all: help

# Self documenting makefile. Double hashes signify help comments.
help:                   ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

install:                ## Installs tables-to-go. Same behavior like `go install -mod=vendor .`
	go install -mod=vendor .

test:                   ## Runs unit tests with race flag enabled
	go test -mod=vendor -race ./...

integration-test:       ## Runs integration tests
	go test -mod=vendor -tags=integration ./test/...

sqlite3:                ## Installs tables-to-go with sqlite3 driver and the \
                        ## User Authentication feature enabled. \
                        ## For more information see the documentation of the driver at \
                        ## - https://github.com/mattn/go-sqlite3#compilation \
                        ## - https://github.com/mattn/go-sqlite3#user-authentication
	CGO_ENABLED=1 go install -mod=vendor -tags="sqlite3 sqlite_userauth" .
