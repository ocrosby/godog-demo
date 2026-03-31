FROM golang:1.25

WORKDIR /app

# Cache module downloads as a separate layer so they are not re-fetched on
# every source change.
COPY go.mod go.sum ./
RUN go mod download

# Copy source after dependencies so the module cache layer is reused when
# only application code changes.
COPY . .

# Run the full BDD test suite. -count=1 disables the test result cache so
# the suite always executes fresh inside the container.
CMD ["go", "test", "-v", "-count=1", "./features/steps/..."]
