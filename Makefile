.PHONY: test
test:
	go test

BINARY := pat
MAIN := main.go

.PHONY: linux
linux:
	mkdir -p build
	go build -o build/$(BINARY) cmd/$(MAIN)
	chmod +x build/$(BINARY)

.PHONY: build
build: test linux
