.PHONY: build, lint, vet

build:
	go build -v -o apiserver ./cmd/apiserver/

lint:
	golint ./...

vet:
	go vet ./...
.DEFAULT_GOAL := build