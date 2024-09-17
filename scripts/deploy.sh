#!/bin/bash

set -e

IMAGE_VERSION="v0.0.5"

cd ui && npm run build
cd - || exit 1

GOOS=linux GOARCH=amd64 go build -o ./bin/api
docker buildx build --platform="linux/amd64" --push -t "lucasrod96/oss-contribute:${IMAGE_VERSION}" .

terraform plan -var="image_version=${IMAGE_VERSION}"
terraform apply --auto-approve -var="image_version=${IMAGE_VERSION}"
