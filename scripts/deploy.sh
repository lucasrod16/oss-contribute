#!/bin/bash

set -euo pipefail

cd ui && npm ci --legacy-peer-deps && npm run build
cd - || exit 1

GOOS=linux GOARCH=amd64 go build -o ./bin/api
docker buildx build --platform="linux/amd64" --push -t "lucasrod96/oss-contribute:${IMAGE_VERSION}" .

terraform init
terraform plan -var="image_version=${IMAGE_VERSION}"
terraform apply --auto-approve -var="image_version=${IMAGE_VERSION}"
