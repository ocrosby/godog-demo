.PHONY: clean lint test deps install docker-build docker-test

test:
	#go test ./... -v
	gotestsum --format testname --junitfile junit.xml -- -v

clean:
	rm -f godog-demo
	rm -f junit.xml

deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install gotest.tools/gotestsum@latest

lint:
	golangci-lint run

install: deps
	go mod tidy
	go mod download

docker-build:
	docker build -t godog-demo .

docker-test: docker-build
	docker run --rm godog-demo
	