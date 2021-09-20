.PHONY: build, test, lint, vet

build:
	go build -v -o wget ./cmd/app/

lint:
	golint ./...

test:
	go test -v ./...

vet:
	go vet ./...

.DEFAULT_GOAL := build