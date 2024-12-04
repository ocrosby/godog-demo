.PHONY: lint test deps install

deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	golangci-lint run

test:
	go test ./... -v

install: deps
	go mod tidy
	go mod download
