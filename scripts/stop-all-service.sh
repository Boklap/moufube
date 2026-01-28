#!/bin/bash

set -e

echo "Stopping all containers on moufube-dev network..."

CONTAINERS=$(docker ps --filter "network=moufube-dev" --format "{{.Names}}")

if [ -z "$CONTAINERS" ]; then
    echo "No running containers found on moufube-dev network"
else
    echo "Stopping containers: $CONTAINERS"
    echo "$CONTAINERS" | xargs docker stop
fi

echo "✅ All containers stopped"
