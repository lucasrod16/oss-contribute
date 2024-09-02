.PHONY: ui build

ui:
	cd ui && npm run build

build: ui
	go build -o ./bin/api

build-linux-arm: ui
	GOOS=linux GOARCH=arm64 go build -o ./bin/api

build-linux-amd: ui
	GOOS=linux GOARCH=amd64 go build -o ./bin/api

image-prod: build-linux-amd
	docker buildx build --platform="linux/amd64" --push -t lucasrod96/oss-contribute:v0.0.3 .
