.PHONY: clean lint test deps install docker-build docker-test

test:
	go test ./... -v

clean:
	rm -f godog-demo

deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	golangci-lint run

install: deps
	go mod tidy
	go mod download

docker-build:
	docker build -t godog-demo .

docker-test: docker-build
	docker run --rm godog-demo
	