#!/bin/sh

set -e

BASE_DIR="$(dirname "$0")"
LINTER_DOCKERFILE="$BASE_DIR/../tools/lint.Dockerfile"

echo "Running golangci-lint-v2...\n"

echo "Building linter Docker image..."
docker build -t moufube-linter -f "$LINTER_DOCKERFILE" "$BASE_DIR"

echo "Running linter in container..."

docker run --rm -v "$BASE_DIR/../services:/app/services" moufube-linter sh -c '
    set -e
    SERVICES="api-gateway authentication"
    for service in $SERVICES; do
        echo "Linting $service..."
        cd /app/services/$service
        if ! golangci-lint-v2 run; then
            echo "❌ Lint failed for $service. Please fix the issues above before committing."
            exit 1
        fi
    done
    echo "✅ Lint passed"
'