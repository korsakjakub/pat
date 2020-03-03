SRC := $(shell go list ./... | grep -v /vendor)

.PHONY: test
test:
	go test $(SRC)

BINARY := pat
MAIN := main.go

.PHONY: linux
linux:
	mkdir -p build
	go build -o build/$(BINARY) cmd/$(MAIN)

.PHONY: build
build: test linux
