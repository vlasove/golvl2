.PHONY: build, test, lint, vet

build:
	go build -v -o mancut ./cmd/app/

test:
	go test -v ./...

lint:
	golint ./...

vet:
	go vet ./...

.DEFAULT_GOAL := build