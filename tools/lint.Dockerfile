FROM golang:1.25.4

WORKDIR /app
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b /tmp/golangci v2.7.2 \
    && mv /tmp/golangci/golangci-lint /go/bin/golangci-lint-v2

CMD ["sleep", "infinity"]