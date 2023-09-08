projectDir := $(realpath $(dir $(firstword $(MAKEFILE_LIST))))
os := $(shell uname)
VERSION ?= $(shell git rev-parse --short HEAD)
## localhost by default, can be overridden at runtime when running make target e.g REGISTRY=minikube:5000 make <target_name>
REGISTRY ?= localhost:5000
image_name = reference-service
image_tag = latest
tenant_namespace = tenant-1

# Autograding related settings
#TODO - when we create a public repo we need to update the CHART_REPO
CHART_REPO := my-local-helm-repo
P2P_AUTOGRADING_CHART := $(CHART_REPO)/bootcamp-p2p-autograding
NFT_OBS_AUTOGRADING_CHART := $(CHART_REPO)/bootcamp-nft-obs-autograding
PUSHGATEWAY_AUTOGRADING_CHART := $(CHART_REPO)/bootcamp-pushgateway-autograding

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

# Autograding related tasks
.PHONY: autograde-p2p
autograde-p2p: ## Run p2p autograding
	helm repo update
	helm uninstall bootcamp-p2p-autograding || true
	sleep 5
	helm install bootcamp-p2p-autograding $(P2P_AUTOGRADING_CHART)
	@echo "Waiting for job output.."
	@slept=0
	@while ! kubectl -n p2p-autograding get pod -l job-name=p2p-acceptance-tests | grep -q "Running"; do \
		sleep 1; \
		((slept=slept+1)); \
		if [ $$slept -eq 15 ]; then \
			echo "15 Seconds have elapsed and the job has not run -- something is wrong, exiting.."; \
			break; \
		fi; \
	done; \
	kubectl -n p2p-autograding logs -f job/p2p-acceptance-tests

.PHONY: autograde-nft-obs
autograde-nft-obs: ## Run nft observability autograding
	helm repo update
	helm uninstall bootcamp-nft-obs-autograding || true
	sleep 5
	helm install bootcamp-nft-obs-autograding $(NFT_OBS_AUTOGRADING_CHART)
	@echo "Waiting for job output.."
	@slept=0
	@while ! kubectl -n nft-obs-autograding get pod -l job-name=nft-obs-acceptance-tests | grep -q "Running"; do \
		sleep 1; \
		((slept=slept+1)); \
		if [ $$slept -eq 15 ]; then \
			echo "15 Seconds have elapsed and the job has not run -- something is wrong, exiting.."; \
			break; \
		fi; \
	done; \
	kubectl -n nft-obs-autograding logs -f job/nft-obs-acceptance-tests
