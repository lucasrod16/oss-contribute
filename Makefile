.PHONY: build-dev
build-dev:
	npm --prefix ui run build && CGO_ENABLED=0 go build -o ./bin/api
