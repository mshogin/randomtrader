PROJECT_NAME := randomtrader
PKG := "github.com/mshogin/$(PROJECT_NAME)"
PKG_LIST := $(go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
LINTPKG = github.com/golangci/golangci-lint/cmd/golangci-lint@v1.24.0
ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif
LINTBIN = $(GOPATH)/bin/golangci-lint

.PHONY: all dep build clean test coverage coverhtml lint

all: build

linter:
	go get golang.org/x/lint/golint
	golint ./...
	GO111MODULE=on go get $(LINTPKG)
	test -z "$$($(LINTBIN) run --verbose | tee /dev/stderr)"

check: linter test

test: ## Run unittests
	go test -coverprofile=coverage.txt -covermode=atomic  ./...

race: dep ## Run data race detector
	go test -race -coverprofile=coverage.txt -covermode=atomic  ./...

msan: dep ## Run memory sanitize
	go test -msan -short ${PKG_LIST}

coverage: ## Generate global code coverage report
	./tools/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	./tools/coverage.sh html;

dep: ## Get the dependencies
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod tidy
	go get -v -d ./...
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/lint/golint

build: dep ## Build the binary file
	go build -tags=prod -o cmd/randomtrader/randomtrader cmd/randomtrader/main.go
	go build -tags=prod -o cmd/datacollector/randomtrader-datacollector cmd/datacollector/main.go

clean: ## Remove previous build
	rm -f cmd/randomtrader/randomtrader

help: ## Display this help screen
	grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
