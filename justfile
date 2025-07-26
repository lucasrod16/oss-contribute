build-dev:
	npm --prefix ui run build
	CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/api

test-unit:
	@echo "================RUNNING UNIT TESTS================\n"
	go test -race -v -count=1 -failfast ./internal/...
	@echo "================FINISHED RUNNING UNIT TESTS================\n"

test-smoke:
	@echo "================RUNNING SMOKE TESTS================\n"
	./scripts/smoke-tests.sh
	@echo "================FINISHED RUNNING SMOKE TESTS================\n"

test-e2e:
	@echo "================RUNNING E2E TESTS================\n"
	./scripts/e2e-tests.sh
	@echo "================FINISHED E2E TESTS================\n"

test-e2e-dev:
	@echo "================RUNNING E2E TESTS================\n"
	npm --prefix=ui run test:e2e-dev
	@echo "================FINISHED E2E TESTS================\n"

test: test-unit test-smoke test-e2e

lint:
	npm --prefix ui run lint

lint-fix:
	npm --prefix ui run lint:fix

release tag:
    ./scripts/release.sh {{tag}}
