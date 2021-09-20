.PHONY: build, test, lint, vet

build:
	go build -v -o app .

test:
	go test -v ./...

lint:
	golint ./...

vet:
	go vet ./...

.DEFAULT_GOAL := build