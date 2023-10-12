projectDir := $(realpath $(dir $(firstword $(MAKEFILE_LIST))))
os := $(shell uname)
VERSION ?= $(shell git rev-parse --short HEAD)
registry = minikube:5000
image_name = bootcamp/reference-service-go-pub
image_tag = latest

.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed 's/^.*Makefile://' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# P2P tasks

.PHONY: local
local: build local-stubbed-functional local-stubbed-nft ## Build the reference application, and run functional and non-functional tests locally

.PHONY: build
build: ## Build the project
	docker compose run --quiet-pull --rm go_build sh -c 'go build && cd ../../ && go test ./...'

.PHONY: local-stubbed-functional
local-stubbed-functional: ## Build the reference application
	docker compose build service database downstream godog --no-cache
	docker compose up -d service downstream database waitForHealthyPods
	docker compose run --rm godog sh -c 'cd functional/godogs; go test'
	docker compose down -v

.PHONY: local-stubbed-nft
local-stubbed-nft: ## Run non-functional tests locally
	docker compose build service database downstream  --no-cache
	docker compose up -d service downstream database waitForHealthyPods
	docker compose run --rm k6 run ./nft/ramp-up/test.js
	docker compose down -v

.PHONY: stubbed-functional
stubbed-functional: ## Perform functional tests against stubs
	docker compose run --rm godog sh -c 'cd functional/godogs; godog run'

.PHONY: stubbed-nft
stubbed-nft: ## Perform non-functional tests against stubs
	docker compose run --rm k6 run ./nft/ramp-up/test.js

.PHONY: extended-stubbed-nft
extended-stubbed-nft: ## Perform extended non-functional tests against stubs
	@echo "Not implemented!"

.PHONY: integrated
integrated: ## Perform integrated tests
	@echo "Not implemented!"

# Custom tasks
.PHONY: run-local
run-local: ## Run the reference application locally
	docker compose build service --no-cache
	docker compose up -d database downstream
	docker compose run --service-ports --rm service

# Minikube local tasks
.PHONY: docker-build
docker-build: ## Build and tag the service image
	docker build --file Dockerfile.service --tag $(registry)/$(image_name):$(image_tag) .

.PHONY: docker-push
docker-push: ## Push the service image
	docker image push $(registry)/$(image_name):$(image_tag)

# Autograding related tasks
.PHONY: autograde-p2p
autograde-p2p: ## Run autograding for the module "p2p"
	echo "Not implemented yet"

.PHONY: autograde-nft-obs
autograde-nft-obs: ## Run autograding for the module "nft-observability"
	echo "Not implemented yet"


