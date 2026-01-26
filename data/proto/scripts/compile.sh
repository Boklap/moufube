#!/bin/sh

# Get proto directory and output directory
PROTO_DIR="/app/proto"
OUTPUT_DIR="/app/pb"

# Create output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Find all .proto files recursively and compile them
find "$PROTO_DIR" -name "*.proto" -type f | while read proto_file; do
  relative_path="${proto_file#$PROTO_DIR/}"
  output_file="$OUTPUT_DIR/${relative_path%.proto}.pb"

  echo "Compiling: $relative_path → $output_file"

  protoc \
    -I "$PROTO_DIR" \
    --go_out="$OUTPUT_DIR" \
    --go_opt=paths=source_relative \
    --go-grpc_out="$OUTPUT_DIR" \
    --go-grpc_opt=paths=source_relative \
    "$proto_file"

  if [ $? -eq 0 ]; then
    echo "✓ Success"
  else
    echo "✗ Failed"
    exit 1
  fi
done

echo ""
echo "All proto files compiled to $OUTPUT_DIR/"