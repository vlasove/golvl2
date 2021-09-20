.PHONY: test, lint, vet

test:
	go test -v ./...

lint:
	golint ./...

vet:
	go vet ./...

.DEFAULT_GOAL := test