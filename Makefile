.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test ./...

.DEFAULT_GOAL := build