#!/bin/sh

set -e

echo "Running golangci-lint-v2...\n"

echo "Linting api-gateway-dev"
if ! docker exec api-gateway-dev golangci-lint-v2 run; then
    echo "❌ Lint failed. Please fix the issues above before committing.\n"
    exit 1
fi

echo "Linting authentication-service"
if ! docker exec authentication-service-dev golangci-lint-v2 run; then
    echo "❌ Lint failed. Please fix the issues above before committing.\n"
    exit 1
fi

echo "✅ Lint passed"