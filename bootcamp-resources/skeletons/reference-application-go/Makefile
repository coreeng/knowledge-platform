projectDir := $(realpath $(dir $(firstword $(MAKEFILE_LIST))))
os := $(shell uname)
VERSION ?= $(shell git rev-parse --short HEAD)
registry = minikube:5000
image_name = bootcamp/reference-service-go-pub
image_tag = latest

# Autograding related settings
#TODO - when we create a public repo we need to update the CHART_REPO
CHART_REPO := my-local-helm-repo
P2P_AUTOGRADING_CHART := bootcamp-p2p-autograding
NFT_OBS_AUTOGRADING_CHART := bootcamp-nft-obs-autograding
MULTI_TENANCY_AUTOGRADING_CHART := bootcamp-multi-tenancy-autograding
PUSHGATEWAY_AUTOGRADING_CHART := bootcamp-pushgateway-autograding

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
autograde-p2p: install-autograding-dependencies ## Run autograding for the module "p2p"
	helm repo update
	helm uninstall $(P2P_AUTOGRADING_CHART) --wait || true
	helm install $(P2P_AUTOGRADING_CHART) $(CHART_REPO)/$(P2P_AUTOGRADING_CHART)
	@echo "Waiting for job output.."
	@kubectl wait pods -n p2p-autograding --for condition=Ready --selector=job-name=p2p-acceptance-tests
	@kubectl -n p2p-autograding logs -f job/p2p-acceptance-tests

.PHONY: autograde-nft-obs
autograde-nft-obs: install-autograding-dependencies## Run autograding for the module "nft-observability"
	helm repo update
	helm uninstall $(NFT_OBS_AUTOGRADING_CHART) --wait || true
	helm install $(NFT_OBS_AUTOGRADING_CHART) $(CHART_REPO)/$(NFT_OBS_AUTOGRADING_CHART)
	@echo "Waiting for job output.."
	@kubectl wait pods -n nft-observability-autograding --for condition=Ready --selector=job-name=nft-observability-acceptance-tests
	@kubectl -n nft-observability-autograding logs -f job/nft-observability-acceptance-tests

.PHONY: autograde-multi-tenancy
autograde-multi-tenancy: install-autograding-dependencies## Run autograding for the module "multi-tenancy"
	helm repo update
	helm uninstall $(MULTI_TENANCY_AUTOGRADING_CHART) --wait || true
	helm install $(MULTI_TENANCY_AUTOGRADING_CHART) $(CHART_REPO)/$(MULTI_TENANCY_AUTOGRADING_CHART)
	@echo "Waiting for job output.."
	@kubectl wait pods -n multi-tenancy-autograding --for condition=Ready --selector=job-name=multi-tenancy-acceptance-tests
	@kubectl -n multi-tenancy-autograding logs -f job/multi-tenancy-acceptance-tests

.PHONY: install-autograding-dependencies
install-autograding-dependencies:
	helm repo update
	helm uninstall bootcamp-pushgateway-autograding --wait || true
	helm install $(PUSHGATEWAY_AUTOGRADING_CHART) $(CHART_REPO)/$(PUSHGATEWAY_AUTOGRADING_CHART)

