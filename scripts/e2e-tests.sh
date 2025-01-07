#!/bin/bash

set -euo pipefail

BLUE='\033[0;34m'
NONE='\033[0m'

cleanup() {
    echo -e "${BLUE}=== Application logs ===${NONE}"
    docker logs oss-contribute
    echo -e "${BLUE}=== Application logs ===${NONE}"
    docker compose down
}

trap cleanup EXIT

health_check() {
  echo "Waiting for app to be ready..."

  start_time=$(date +%s)

  until curl -fsL "http://localhost:8080" > /dev/null; do
    sleep 2
    current_time=$(date +%s)
    elapsed_time=$((current_time - start_time))
    if [ $elapsed_time -ge 10 ]; then
        echo "Error: Timed out waiting for app to become healthy."
        exit 1
    fi
  done

  echo "app is ready!"
}

npm --prefix ui ci
npm --prefix ui run build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/api

docker compose up -d --build
health_check
npm --prefix ui run test:e2e
