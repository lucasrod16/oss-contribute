.PHONY: build-dev
build-dev:
	npm --prefix ui run build && go build -o ./bin/api
