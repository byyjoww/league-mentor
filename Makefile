SHELL := /bin/sh
PACKAGES := $(shell go list ./...)
ifdef DEBUG
$(info PACKAGES = $(PACKAGES))
$(info SOURCES = $(SOURCES))
endif

################################################################################
## Standard make targets
################################################################################

.DEFAULT_GOAL := all
.PHONY: all
all: fix lint build

.PHONY: build
build:
	go build .

.PHONY: run-api
run-api:
	build
	league-mentor start --type=api

.PHONY: dev-run-api
dev-run-api:
	@make build
	@make setup-dev-env
	./league-mentor start --type=api

.PHONY: run-client
run-client:
	build
	league-mentor start --type=client

.PHONY: dev-run-client
dev-run-client:
	@make build
	@make setup-dev-env
	./league-mentor start --type=client

.PHONY: setup-dev-env
setup-dev-env:
	echo "setting up dev environment - TODO: add local env variables"

################################################################################
## Tests
################################################################################

TESTABLE_PACKAGES := $$(go list ./... | egrep -v 'constants|mocks|testing')

.PHONY: test-unit
test-unit:
	echo 'Starting unit tests...'
	go test -timeout 50000ms -tags=unit -coverprofile=coverage.out ./...

.PHONY: test-integration
test-integration:
	echo 'Starting integration tests...'
	@GO111MODULE=on go test ${TESTABLE_PACKAGES} -tags=integration -coverprofile integration.coverprofile
	go test -tags=integration -coverprofile=coverage.out ./...
	echo 'Removing test environment...'

################################################################################
## Linters and formatters
################################################################################

.PHONY: fix
fix:
	go mod tidy

.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run
	go run golang.org/x/lint/golint $(PACKAGES)