#!/bin/bash

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

docker run --rm \
  -v "$PROJECT_ROOT/data:/app" \
  proto-compiler \
  bash -c "chmod +x /app/proto/scripts/compile.sh && /app/proto/scripts/compile.sh"
