#!/usr/bin/env bash

set -euo pipefail

REPO_ROOT=$(git rev-parse --show-toplevel)
UI_DIR="${REPO_ROOT}/ui"

npm --prefix="${UI_DIR}" ci
npm --prefix="${UI_DIR}" run build

pushd "${REPO_ROOT}"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/api
docker buildx build --platform="linux/amd64" --push -t "lucasrod96/oss-projects:${IMAGE_TAG}" .
popd

IMAGE_DIGEST=$(docker buildx imagetools inspect "lucasrod96/oss-projects:${IMAGE_TAG}" --format "{{json .Manifest}}" | jq -r '.digest')
yq -i ".services.oss-projects.image = \"lucasrod96/oss-projects:${IMAGE_TAG}@${IMAGE_DIGEST}\"" docker-compose.yml
