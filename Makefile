SERVICE		?= $(shell basename `go list`)
VERSION		?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || cat $(PWD)/.version 2> /dev/null || echo v0)
PACKAGE		?= $(shell go list)
PACKAGES	?= $(shell go list ./...)
FILES		?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")
INSTALL_PATH ?= /usr/local

.PHONY: help clean fmt lint vet build all

default: help

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

all:    ## clean, format, build and unit test
	make clean
	make gofmt
	make build
	make test

install: build   ## build and install go application executable
	sudo cp ./etcaid ${INSTALL_PATH}/bin

uninstall: ## uninstall application executable
	sudo rm ${INSTALL_PATH}/bin/etcaid

env:    ## Print useful environment variables to stdout
	@echo $(CURDIR)
	@echo $(SERVICE)
	@echo $(PACKAGE)
	@echo $(VERSION)

clean:  ## go clean
	go clean
	rm ./etcaid

fmt:    ## format the go source files
	go fmt ./...
	goimports -w $(FILES)

lint:   ## run go lint on the source files
	golint $(PACKAGES)

vet:    ## run go vet on the source files
	go vet ./...

doc:    ## generate godocs and start a local documentation webserver on port 8085
	godoc -http=:8085 -index

build:
	go build ./cmd/etcaid

test:
	go test -v ./... -short

test-it:
	go test -v ./...

test-all: test test-cover
