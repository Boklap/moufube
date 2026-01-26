FROM golang:1.25.4

WORKDIR /app

RUN apt-get update && apt-get install -y unzip \
    && PB_REL="https://github.com/protocolbuffers/protobuf/releases" \
    && curl -LO $PB_REL/download/v30.2/protoc-30.2-linux-x86_64.zip \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.0

RUN unzip protoc-30.2-linux-x86_64.zip -d /usr/local \
    && chmod +x /usr/local/bin/protoc

ENV PATH="/go/bin:/usr/local/go/bin:/usr/local/bin:$PATH"

CMD ["sleep", "infinity"]