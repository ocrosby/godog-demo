# godog-demo

A demonstration project showing how to write BDD acceptance tests in Go using [GoDog](https://github.com/cucumber/godog) (the official Cucumber framework for Go). The test suite exercises the public [JSONPlaceholder](https://jsonplaceholder.typicode.com) REST API across multiple resource domains, plus a local calculator example.

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Project Structure](#project-structure)
- [Running Tests](#running-tests)
- [Task Reference](#task-reference)
- [Docker](#docker)
- [Contributing](#contributing)
- [License](#license)
- [References](#references)

---

## Overview

This project demonstrates:

- Writing Gherkin feature files that describe API behavior from the user's perspective
- Implementing step definitions in Go using the GoDog framework
- Sharing scenario state via a context object
- Running BDD tests with `gotestsum` for rich terminal output and JUnit XML reporting
- Using [Task](https://taskfile.dev) as the project task runner

The test suite targets the JSONPlaceholder API and covers albums, comments, photos, posts, todos, and users — plus a standalone calculator feature to illustrate local unit-style BDD.

---

## Features

- BDD acceptance tests written in Gherkin, implemented in Go
- Full coverage of JSONPlaceholder resource domains (albums, comments, photos, posts, todos, users)
- Local calculator feature for self-contained BDD examples
- JUnit XML test report output (`junit.xml`) for CI integration
- Randomized scenario execution order to catch ordering-dependent failures
- Linting with `golangci-lint`
- Docker support for fully isolated test runs
- [Task](https://taskfile.dev) for all development workflows

---

## Prerequisites

Ensure the following are installed before setting up the project.

### Required

| Tool | Minimum Version | Install |
|------|----------------|---------|
| [Go](https://go.dev/dl/) | 1.25 | `brew install go` |
| [Task](https://taskfile.dev/installation/) | 3.x | `brew install go-task` |
| [Git](https://git-scm.com) | any | `brew install git` |

### Installed automatically by `task deps`

| Tool | Purpose |
|------|---------|
| [golangci-lint](https://golangci-lint.run) | Go linter |
| [gotestsum](https://github.com/gotestyourself/gotestsum) | Enhanced test runner with JUnit output |

### Optional

| Tool | Purpose |
|------|---------|
| [Docker](https://www.docker.com/get-started) | Run tests in an isolated container |

---

## Installation

### 1. Clone the repository

```sh
git clone https://github.com/ocrosby/godog-demo.git
cd godog-demo
```

### 2. Install development tools

This installs `golangci-lint` and `gotestsum` to your `$GOPATH/bin`:

```sh
task deps
```

### 3. Install Go module dependencies

```sh
task install
```

This runs `task deps` (above), then `go mod tidy` and `go mod download`.

### Verify the setup

```sh
task --list
```

You should see all available tasks listed. If `task` is not found, ensure `$(go env GOPATH)/bin` is on your `PATH`:

```sh
export PATH="$(go env GOPATH)/bin:$PATH"
```

Add that line to your shell profile (`~/.zshrc`, `~/.bashrc`, etc.) to make it permanent.

---

## Configuration

This project has no runtime configuration files. All behavior is driven by the feature files in `features/` and the test harness in `features/steps/`.

### Environment

The test suite makes live HTTP requests to `https://jsonplaceholder.typicode.com`. No API key or authentication is required. Ensure outbound HTTPS access is available from your machine (or container).

### Test output

`gotestsum` writes a JUnit XML report to `junit.xml` in the project root after each test run. This file is consumed by most CI systems (GitHub Actions, Jenkins, GitLab CI) for test result visualization. It is listed in `.gitignore` and excluded from version control.

### Linter configuration

`golangci-lint` uses its default configuration. To customize it, add a `.golangci.yml` file to the project root. See the [golangci-lint documentation](https://golangci-lint.run/usage/configuration/) for available options.

---

## Project Structure

```
godog-demo/
├── features/                   # Gherkin feature files (one per domain)
│   ├── albums.feature
│   ├── calculator.feature
│   ├── comments.feature
│   ├── photos.feature
│   ├── posts.feature
│   ├── todos.feature
│   ├── users.feature
│   └── steps/                  # Go step definitions and test harness
│       ├── main_test.go        # TestFeatures entry point, wires all scenarios
│       ├── album_steps.go
│       ├── calculator_steps.go
│       ├── comment_steps.go
│       ├── common_steps.go     # Shared steps (error checking, response assertions)
│       ├── photo_steps.go
│       ├── post_steps.go
│       ├── todo_steps.go
│       └── user_steps.go
├── pkg/
│   ├── calculator.go           # Calculator implementation (local BDD example)
│   ├── helpers/
│   │   └── request_helper.go   # HTTP request utilities
│   └── models/                 # Domain model structs (Album, Comment, etc.)
│       ├── album.go
│       ├── comment.go
│       ├── photo.go
│       ├── post.go
│       ├── todo.go
│       └── user.go
├── Dockerfile
├── Taskfile.yml
├── go.mod
├── go.sum
└── README.md
```

---

## Running Tests

All tests are run through the `features/steps` package using the standard `go test` toolchain, wrapped by `gotestsum` for better output.

### Run all tests

```sh
task test
```

This runs `gotestsum --format testname --junitfile junit.xml -- -v -count=1 ./features/steps` and streams results to the terminal as they complete. A `junit.xml` report is written to the project root.

### Run tests for a specific feature

Pass Go test filter flags directly if you need to target one domain:

```sh
gotestsum --format testname -- -v -count=1 -run TestFeatures/albums ./features/steps
```

Supported sub-test names match the feature file paths registered in `main_test.go`:

| Sub-test name | Feature file |
|--------------|--------------|
| `albums` | `features/albums.feature` |
| `comments` | `features/comments.feature` |
| `photos` | `features/photos.feature` |
| `posts` | `features/posts.feature` |
| `todos` | `features/todos.feature` |
| `users` | `features/users.feature` |
| `calculator` | `features/calculator.feature` |

### Run with the race detector

```sh
gotestsum --format testname -- -race -count=1 ./features/steps
```

### View the JUnit report

After running tests, `junit.xml` contains the full test report in JUnit XML format. Most CI platforms parse this automatically. To inspect it locally:

```sh
cat junit.xml
```

---

## Task Reference

All project workflows are managed by [Task](https://taskfile.dev). Run `task` with no arguments to list available tasks.

```sh
task              # list all tasks
task install      # install tools and download Go modules
task deps         # install golangci-lint and gotestsum
task lint         # run golangci-lint
task test         # run the full test suite
task clean        # remove build artifacts (godog-demo binary, junit.xml)
task docker-build # build the Docker image
task docker-test  # build image and run tests inside Docker
```

### task install

Installs development tools (`task deps`) then runs `go mod tidy` and `go mod download`. Run this once after cloning or after any dependency change.

```sh
task install
```

### task deps

Installs `golangci-lint` and `gotestsum` using `go install`. These land in `$(go env GOPATH)/bin`.

```sh
task deps
```

### task lint

Runs `golangci-lint run` against the entire module. Fix any reported issues before opening a pull request.

```sh
task lint
```

### task test

Runs the full BDD test suite using `gotestsum`. Scenarios execute in randomized order. Output uses the `testname` format — each scenario prints as a single line. A JUnit XML report is written to `junit.xml`.

```sh
task test
```

### task clean

Removes the compiled `godog-demo` binary and `junit.xml` test report from the project root.

```sh
task clean
```

### task docker-build

Builds the `godog-demo` Docker image using the project `Dockerfile`.

```sh
task docker-build
```

### task docker-test

Builds the Docker image (via `task docker-build`) then runs the container. Useful for verifying the test suite passes in a clean, isolated environment.

```sh
task docker-test
```

---

## Docker

The project includes a `Dockerfile` for running tests in an isolated environment without requiring a local Go installation.

### Build the image

```sh
task docker-build
# or directly:
docker build -t godog-demo .
```

### Run tests in the container

```sh
task docker-test
# or directly:
docker run --rm godog-demo
```

---

## Contributing

Contributions are welcome. Please follow these steps:

1. Fork the repository and create a feature branch from `main`.
2. Install dependencies: `task install`
3. Make your changes. Add or update feature files and step definitions as appropriate.
4. Ensure all tests pass: `task test`
5. Ensure there are no lint errors: `task lint`
6. Open a pull request against `main` with a clear description of the change and why it was made.

### Commit style

This project follows [Conventional Commits](https://www.conventionalcommits.org/):

```
feat(albums): add scenario for updating an album
fix(steps): correct assertion for empty response body
docs: update task reference in README
```

---

## License

This project is provided as a demonstration and is not under an official open-source license. See the repository for any applicable terms.

---

## References

- [GoDog — GitHub](https://github.com/cucumber/godog)
- [GoDog examples](https://github.com/cucumber/godog/tree/main/_examples)
- [Cucumber / Gherkin reference](https://cucumber.io/docs/gherkin/reference/)
- [JSONPlaceholder](https://jsonplaceholder.typicode.com) — the fake REST API used by the test suite
- [Task — taskfile.dev](https://taskfile.dev)
- [gotestsum](https://github.com/gotestyourself/gotestsum)
- [golangci-lint](https://golangci-lint.run)
- [GOLANG & API Testing with GoDog](https://medium.com/propertyfinder-engineering/golang-api-testing-with-godog-2de8944d2511)
