# Development
FROM golang:1.25.4 AS development

WORKDIR /app

RUN apt-get update && apt-get install -y unzip \
    && go install golang.org/x/tools/gopls@v0.21.0 \
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b /tmp/golangci v2.7.2 \
    && go install github.com/air-verse/air@v1.63.4 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.0 \
    && PB_REL="https://github.com/protocolbuffers/protobuf/releases" \
    && curl -LO $PB_REL/download/v30.2/protoc-30.2-linux-x86_64.zip

RUN unzip protoc-30.2-linux-x86_64.zip -d /usr/local \
    && chmod +x /usr/local/bin/protoc

# Move the installed golang-ci from temp folder to go/bin
# We do this to fix the bug where VSCode will pass the v1 flags as argument
# When running the linter.
RUN mv /tmp/golangci/golangci-lint /go/bin/golangci-lint-v2

COPY go.mod go.sum ./

COPY go.mod ./

RUN go mod download

ENV PATH="/go/bin:/usr/local/go/bin:/usr/local/bin:$PATH"

CMD ["sleep", "infinity"]