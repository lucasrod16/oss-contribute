#!/bin/bash

set -euo pipefail

npm --prefix=ui ci
npm --prefix=ui run build

GOOS=linux GOARCH=amd64 go build -o ./bin/api

echo "${DOCKER_PASSWORD}" | docker login --username="${DOCKER_USERNAME}" --password-stdin
ls -lah ./bin/api
docker buildx build --no-cache --platform="linux/amd64" --push -t "lucasrod96/oss-contribute:${IMAGE_VERSION}" .

terraform init
terraform plan -var="image_version=${IMAGE_VERSION}"
terraform apply --auto-approve -var="image_version=${IMAGE_VERSION}"
