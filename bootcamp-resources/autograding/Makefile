CURRENT_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

MODULE :=
SUPPORTED_MODULES := "multi-tenancy nft-onservability p2p platform-engineering"

PUSHGATEWAY_CHART_PATH = $(realpath $(CURRENT_PATH)/helm-charts/pushgateway)
CHART_PATH := $(realpath $(CURRENT_PATH)/helm-charts/autograding)
CHART_NAME := $(shell grep -E '^name' $(CHART_PATH)/Chart.yaml | awk '{printf $$2}')
CHART_VERSION := $(shell grep -E '^version' $(CHART_PATH)/Chart.yaml | awk '{printf $$2}'))

# Docker related variables
DOCKER_IMAGE_NAME := ghcr.io/coreeng/$(MODULE)-autograding

.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed 's/^.*Makefile://' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: supported-modules
supported-modules: ## Retrieve a list of supported modules to be used with the make tasks
	@echo $(SUPPORTED_MODULES)

.PHONY: autograde
autograde: build install watch-autograding-job ## Run multi tenancy autograding. Usage: make MODULE=<module-name> autograde

# Public tasks
.PHONY: build
build: ## Build the autograding image (Used by Github actions). Usage: make MODULE=<module-name> build
	$(eval MODULE_PATH := $(realpath $(CURRENT_PATH)/modules/$(MODULE)))
	@eval $$(minikube docker-env) ;\
	docker build -t $(DOCKER_IMAGE_NAME) $(MODULE_PATH)

.PHONY: build-and-push
build-and-push: ## Build and push the autograding image (Used by Github actions). Usage: make MODULE=<module-name> build-and-push
	$(eval MODULE_PATH := $(realpath $(CURRENT_PATH)/modules/$(MODULE)))
	docker build -t $(DOCKER_IMAGE_NAME) $(MODULE_PATH)
	docker push $(DOCKER_IMAGE_NAME)

.PHONY: install-pushgateway
install-pushgateway:
	helm repo update
	helm uninstall bootcamp-pushgateway-autograding --wait || true
	helm install bootcamp-pushgateway-autograding $(PUSHGATEWAY_CHART_PATH)

.PHONY: port-forward-pushgateway
port-forward-pushgateway: ## Port forward push gateway
	kubectl -n pushgateway-autograding port-forward service/pushgateway 9091:9091

.PHONY: install
install: ## Installs the helm charts for a module. Usage: make MODUlE=<module-name> install
	helm repo update
	helm uninstall $(CHART_NAME)-$(MODULE) --wait || true
	helm install --set module=$(MODULE) --set localMode=true $(CHART_NAME)-$(MODULE) $(CHART_PATH)
# Temporary workaround to install monitoring stack as dependency until we find alternative solution
ifeq ($(MODULE), platform-engineering)
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo update
	helm uninstall kube-prometheus-stack --wait || true
	helm install -n platform-engineering-autograding kube-prometheus-stack --set prometheus.prometheusSpec.serviceMonitorSelectorNilUsesHelmValues=false prometheus-community/kube-prometheus-stack
endif

.PHONY: watch-autograding-job
watch-autograding-job:
	$(eval NAMESPACE := $(shell echo $(MODULE)"-autograding"))
	$(eval JOB_NAME := $(shell echo $(MODULE)"-acceptance-tests"))
	@echo "Waiting for the jobs output..."
	@kubectl wait pods -n $(NAMESPACE) --for condition=Ready --selector=job-name=$(JOB_NAME) || true
	@kubectl -n $(NAMESPACE) logs -f job/$(JOB_NAME)
