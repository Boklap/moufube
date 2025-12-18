# Development
FROM golang:1.25.4 AS development

WORKDIR /app

RUN go install github.com/air-verse/air@latest \
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.7.2

COPY go.mod go.sum ./

RUN go mod download

CMD ["sleep", "infinity"]

# Production
FROM golang:1.25.4-alpine3.22 AS build
