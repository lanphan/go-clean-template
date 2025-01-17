ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
.DEFAULT_GOAL := help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

compose-up: ## Run docker-compose
	docker-compose up --build -d postgres && docker-compose logs -f
.PHONY: compose-up

compose-up-integration-test: ## Run docker-compose with integration test
	docker-compose up --build --abort-on-container-exit --exit-code-from integration
.PHONY: compose-up-integration-test

compose-down: ## Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

swag-v1: ## swag init
	swag init -g internal/controller/http/v1/router.go
.PHONY: swag-v1

build: ## build binary file
	GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -tags migrate -o ./bin/app ./cmd/app
.PHONY: build

run: swag-v1 ## swag then run
	go mod tidy && go mod download && \
	DISABLE_SWAGGER_HTTP_HANDLER='' GIN_MODE=debug go run -tags migrate ./cmd/app
.PHONY: run

docker-rm-volume: ## remove docker volume
	docker volume rm go-clean-template_pg-data
.PHONY: docker-rm-volume

linter-golangci: ## check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

linter-hadolint: ## check by hadolint linter
	git ls-files --exclude='Dockerfile*' -i -c | xargs hadolint
.PHONY: linter-hadolint

linter-dotenv: ## check by dotenv linter
	dotenv-linter
.PHONY: linter-dotenv

test: ## run test
	go test -v -cover -race ./internal/...
.PHONY: test

migrate-create:  ## create new migration
	migrate create -ext sql -dir db/migrations 'migrate_name'
.PHONY: migrate-create

migrate-up: ## migration up
	migrate -path db/migrations -database '$(PG_URL)?sslmode=disable' up
.PHONY: migrate-up

migrate-down: ## migration down
	migrate -path db/migrations -database '$(PG_URL)?sslmode=disable' down
.PHONY: migrate-down

test-coverage: ## test-coverage
	go test -v ./... -covermode=count -coverpkg=./... -coverprofile coverage/coverage.out
	go tool cover -html coverage/coverage.out -o coverage/coverage.html
.PHONY: test-coverage

package-dependency: ## package-dependency
	godepgraph -s -novendor -p golang.org/x,google.golang.org,github.com/swaggo,gopkg.in,github.com/jackc,github.com/ethereum,go.uber.org,github.com/gin-gonic,github.com/gin-contrib,github.com/bits-and-blooms,github.com/go-co-op,github.com/xuri,github.com/upper,github.com/ilyakaznacheev ./cmd/app | dot -Tpng -o godepgraph.png
.PHONY: package-dependency

size-analysis: ## check size of each package
	goweight ./cmd/app
.PHONY: size-analysis
