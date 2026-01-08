# Development
FROM golang:1.25.4 AS development

WORKDIR /app
RUN go install golang.org/x/tools/gopls@v0.21.0 \
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b /tmp/golangci v2.7.2 \
    && go install github.com/air-verse/air@v1.63.4

# Move the installed golang-ci from temp folder to go/bin
# We do this to fix the bug where VSCode will pass the v1 flags as argument
# When running the linter.
RUN mv /tmp/golangci/golangci-lint /go/bin/golangci-lint-v2

# COPY go.mod go.sum ./

COPY go.mod ./

RUN go mod download

CMD ["sleep", "infinity"]

# Production
FROM golang:1.25.4-alpine3.22 AS build
