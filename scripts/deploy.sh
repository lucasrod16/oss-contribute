#!/bin/bash

set -euo pipefail

npm --prefix=ui ci
npm --prefix=ui run build

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/api

docker buildx build --platform="linux/amd64" --push -t "lucasrod96/oss-contribute:${IMAGE_VERSION}" .

IMAGE_DIGEST=$(docker buildx imagetools inspect "lucasrod96/oss-contribute:${IMAGE_VERSION}" --format "{{json .Manifest}}" | jq -r '.digest')

terraform init
terraform plan -var="image_digest=${IMAGE_DIGEST}"
terraform apply --auto-approve -var="image_digest=${IMAGE_DIGEST}"
