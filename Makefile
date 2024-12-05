.PHONY: lint test deps install

test:
	go test ./... -v

deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	golangci-lint run

install: deps
	go mod tidy
	go mod download
