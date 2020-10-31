GOPATH ?= $(GOPATH:)
GOLINT=$(GOPATH)/bin/golint
OS := $(shell uname -s)
BIN="./bin"

project=password-validator
name=goapi
version=latest

ifeq ($(OS), Darwin)
	SED_INPLACE = sed -i ''
endif
ifeq ($(OS), Linux)
	SED_INPLACE = sed -i''
endif

test:
	go test -v -cover -covermode=atomic ./...