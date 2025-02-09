run:
	./bin/go run ./app/main.go

tidy:
	./bin/go mod tidy && ./bin/go mod vendor

lint:
	docker compose run --rm app-lint

vet:
	./bin/go vet ./...

test_race:
	./bin/go test -race -timeout=60s -count 1 ./...

test:
	./bin/go clean -testcache
	./bin/go test -race -coverprofile=coverage.out ./...
	grep -v "_mock.go" coverage.out | grep -v mocks > coverage_no_mocks.out
	./bin/go tool cover -func=coverage_no_mocks.out
	rm coverage.out coverage_no_mocks.out

.PHONY: run tidy lint vet test_race test