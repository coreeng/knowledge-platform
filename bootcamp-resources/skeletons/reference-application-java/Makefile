projectDir := $(realpath $(dir $(firstword $(MAKEFILE_LIST))))
os := $(shell uname)
VERSION ?= $(shell git rev-parse --short HEAD)
## localhost by default, can be overridden at runtime when running make target e.g REGISTRY=minikube:5000 make <target_name>
REGISTRY ?= localhost:5000
image_name = reference-service
image_tag = latest
tenant_namespace = tenant-1

.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed 's/^.*Makefile://' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# P2P tasks

.PHONY: local
local: ## Build the reference application, and run functional and non-functional tests locally
	build local-stubbed-functional local-stubbed-nft

.PHONY: build
build: ## Build the reference application
	docker compose run --quiet-pull --rm gradle_build sh -c 'gradle --no-daemon service:build'

.PHONY: local-stubbed-functional
local-stubbed-functional: ## Run functional tests locally
	docker compose build service downstream database --no-cache
	docker compose up -d service downstream database waitForHealthyPods
	docker compose run --rm gradle_build sh -c 'SERVICE_ENDPOINT="http://service:8080" gradle --no-daemon functional:clean functional:test'
	docker compose down
	sudo rm -rf db-data

.PHONY: local-stubbed-nft
local-stubbed-nft: ## Run non-functional tests locally
	docker compose build service downstream database --no-cache
	docker compose up -d database downstream service waitForHealthyPods
	docker compose run --rm k6 run ./nft/ramp-up/test.js
	docker compose down

.PHONY: stubbed-functional
stubbed-functional: ## Perform functional tests against stubs
	docker compose run --rm gradle_build sh -c 'SERVICE_ENDPOINT="http://service:8080" gradle --no-daemon functional:clean functional:test'

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
	docker compose build service downstream --no-cache
	docker compose up -d downstream database
	docker compose run --service-ports --rm service

# Minikube local tasks

.PHONY: docker-build
docker-build: ## Build the reference application, and package as a Docker image
	docker build --file Dockerfile.service --tag $(image_name):$(image_tag) .

.PHONY: docker-push
docker-push: ## Push the Docker image to an image registry
	docker image push $(image_name):$(image_tag)

include ../../autograding/p2p/Makefile
include ../../autograding/nft-observability/Makefile

AUTOGRADE_TASKS = $(shell grep -E '^autograde[a-zA-Z0-9_-]+:.*?$$' $(MAKEFILE_LIST) | sed 's/^.*Makefile://' | awk 'BEGIN {FS = ":"}; {printf "%s ", $$1}')

.PHONY: autograde
autograde: $(AUTOGRADE_TASKS) ## Run all autograding tasks