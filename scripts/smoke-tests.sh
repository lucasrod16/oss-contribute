#!/bin/bash

set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NONE='\033[0m'

cleanup() {
    echo -e "${BLUE}=== Application logs ===${NONE}"
    docker logs oss-contribute
    echo -e "${BLUE}=== Application logs ===${NONE}"
    docker compose down
}

trap cleanup EXIT

smoke_test() {
  local endpoint=$1
  local expected_content_type=$2

  echo -e "${BLUE}Checking '${endpoint}'...${NONE}"

  curl -fsSL "http://localhost:8080${endpoint}" > /dev/null

  content_type=$(curl -s -I "http://localhost:8080${endpoint}" | grep -i "Content-Type:" | awk '{print $2}' | tr -d '\r')

  if [[ "$content_type" != *"$expected_content_type"* ]]; then
    echo -e "${RED}FAILED${NONE}"
    echo -e "${RED}Error: '${endpoint}' returned Content-Type '$content_type', expected '$expected_content_type'${NONE}"
    exit 1
  fi

  echo -e "${GREEN}PASSED${NONE}"
}

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

smoke_test "/" "text/html"
smoke_test "/repos" "application/json"

echo "All smoke tests passed!"
echo "Done!"
