MODULE_NAME=$(shell grep ^module go.mod | cut -d " " -f2)
GIT_COMMIT_HASH=$(shell git rev-parse HEAD)
LD_FLAGS=-ldflags="-X $(MODULE_NAME)/internal/config.gitCommitHash=$(GIT_COMMIT_HASH)"

.PHONY: app-run
app-run:
	@go run ./internal/api/main.go

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: lint-fix
lint-fix:
	@golangci-lint run --fix

.PHONY: docker-up
docker-db:
	@docker-compose -f ./docker-compose.yml up

.PHONY: docker-stop
docker-db-stop:
	@docker-compose -f ./docker-compose.yml stop

.PHONY: generate
generate: swagger
	@go generate ./...

.PHONY: swagger
swagger:
	@go run github.com/swaggo/swag/cmd/swag@v1.7.4 init -g internal/api/main.go -o internal/swagger/docs

.PHONY: test
test:
	@go test -v ./...

	.PHONY: build
build: generate
	@go build -o ./bin/api $(LD_FLAGS) ./cmd/api