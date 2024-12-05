FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download && go install github.com/cucumber/godog/cmd/godog@latest

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o godog-demo ./cmd/godog-demo && chmod +x /app/godog-demo

# Set the entry point to run the tests
ENTRYPOINT ["/app/godog-demo"]