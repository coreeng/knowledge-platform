CURRENT_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
MODULE :=
SUPPORTED_MODULES := "multi-tenancy nft-onservability p2p"

# Charts related variables
CHART_REPO_NAME := my-local-helm-repo
CHART_REPO_PORT := 8080
CHART_REPO := http://localhost:$(CHART_REPO_PORT)

# Docker related variables
DOCKER_IMAGE_NAME := cecg/$(MODULE)-autograding

.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed 's/^.*Makefile://' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: supported-modules
supported-modules: ## Retrieve a list of supported modules to be used with the make tasks
	@echo $(SUPPORTED_MODULES)

.PHONY: autograde
autograde: start-helm-repo-locally build install watch-autograding-job cleanup ## Run multi tenancy autograding. Usage: make MODULE=<module-name> autograde

# Public tasks
.PHONY: build
build: ## Build the autograding image (Used by Github actions). Usage: make MODULE=<module-name> build
	$(eval MODULE_PATH := $(realpath $(CURRENT_PATH)/$(MODULE)))
	@eval $$(minikube docker-env) ;\
	docker build -t $(DOCKER_IMAGE_NAME) $(MODULE_PATH)

.PHONY: build-and-push
build-and-push: ## Build and push the autograding image (Used by Github actions). Usage: make MODULE=<module-name> build-and-push
	$(eval MODULE_PATH := $(realpath $(CURRENT_PATH)/$(MODULE)))
	docker build -t $(DOCKER_IMAGE_NAME) $(MODULE_PATH)
	docker push $(DOCKER_IMAGE_NAME)

.PHONY: install
install: start-helm-repo-locally upload-charts-locally install-charts ## Install the autograding helm charts. Usage: make MODULE=<module-name> install

.PHONY: install-pushgateway
install-pushgateway: MODULE=pushgateway
install-pushgateway: install ## Install pushgateway

.PHONY: port-forward-pushgateway
port-forward-pushgateway: ## Port forward push gateway
	kubectl -n pushgateway-autograding port-forward service/pushgateway 9091:9091

.PHONY: upload-charts-locally
upload-charts-locally: start-helm-repo-locally ## Upload charts locally. Usage: make MODUlE=<module-name> upload-charts-locally.
	cd $(CHART_PATH) && helm package .
	cd $(CHART_PATH) && curl -X "DELETE" $(CHART_REPO)/api/charts/$(CHART_NAME)/$(CHART_VERSION)
	cd $(CHART_PATH) && curl --data-binary "@$(CHART_NAME)-$(CHART_VERSION).tgz" $(CHART_REPO)/api/charts
	cd $(CHART_PATH)  && rm -rf *.tgz

# Private tasks

.PHONY: init-charts-data
init-charts-data:
ifndef MODULE
	$(error "MODULE variable needs to be defined")
endif
	$(eval CHART_PATH := $(realpath $(CURRENT_PATH)/$(MODULE)/charts))
	$(eval CHART_NAME := $(shell grep -E '^name' $(CHART_PATH)/Chart.yaml | awk '{printf $$2}'))
	$(eval CHART_VERSION := $(shell grep -E '^version' $(CHART_PATH)/Chart.yaml | awk '{printf $$2}'))
	$(eval NAMESPACE := $(shell grep -E '^namespace' $(CHART_PATH)/values.yaml | awk '{printf $$2}'))

.PHONY: start-helm-repo-locally
start-helm-repo-locally: init-charts-data
	docker run --rm --name chartmuseum -d \
	  -p 8080:8080 \
	  -e DEBUG=1 \
	  -e STORAGE=local \
	  -e STORAGE_LOCAL_ROOTDIR=/chartstorage \
	  -v $(CURRENT_PATH)/chartstorage:/chartstorage \
	  ghcr.io/helm/chartmuseum:v0.16.0 &
	curl --retry 10 --retry-connrefused http://localhost:8080 > /dev/null
	helm repo add $(CHART_REPO_NAME) $(CHART_REPO) || true

.PHONY: install-charts
install-charts:
	helm repo update
	helm uninstall $(CHART_NAME) --wait || true
	helm install $(CHART_NAME) $(CHART_REPO_NAME)/$(CHART_NAME)

.PHONY: watch-autograding-job
watch-autograding-job:
	$(eval NAMESPACE := $(shell grep -E '^namespace' $(CHART_PATH)/values.yaml | awk '{printf $$2}'))
	@echo "Waiting for the jobs output..."
	@kubectl wait pods -n $(NAMESPACE) --for condition=Ready --selector=job-name=$(MODULE)-acceptance-tests
	@kubectl -n $(NAMESPACE) logs -f job/$(MODULE)-acceptance-tests

.PHONY: cleanup
cleanup:
	# Stops the chartmuseum repo for the helm charts
	docker stop chartmuseum