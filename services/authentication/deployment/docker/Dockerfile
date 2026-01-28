# Development
FROM golang:1.25.4 AS development

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
RUN apt-get update \
    && go install golang.org/x/tools/gopls@v0.21.0 \
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b /tmp/golangci v2.7.2 \
    && go install github.com/air-verse/air@v1.63.4 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.0 \

# Move the installed golang-ci from temp folder to go/bin
# We do this to fix the bug where VSCode will pass the v1 flags as argument
# When running the linter.
    && mv /tmp/golangci/golangci-lint /go/bin/golangci-lint-v2

CMD ["sleep", "infinity"]