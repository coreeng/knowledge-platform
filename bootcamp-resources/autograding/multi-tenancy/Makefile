CURRENT_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

# Charts related variables
CHART_PATH := $(realpath $(CURRENT_PATH)/charts)
CHART_REPO_NAME := my-local-helm-repo
CHART_REPO_PORT := 8080
CHART_REPO := http://localhost:$(CHART_REPO_PORT)
CHART_VERSION:=$(shell grep -E '^version' $(CHART_PATH)/Chart.yaml | awk '{printf $$2}')
CHART_NAME:=$(shell grep -E '^name' $(CHART_PATH)/Chart.yaml | awk '{printf $$2}')

# Docker related variables
DOCKER_IMAGE_NAME := cecg/multi-tenancy-autograding

.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed 's/^.*Makefile://' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: autograde
autograde: build install watch-autograding-job ## Run multi tenancy autograding

.PHONY: build
build: ## Build the autograding image - this task is used by the github actions
	@eval $$(minikube docker-env) ;\
	docker build -t $(DOCKER_IMAGE_NAME) $(CURRENT_PATH)

.PHONY: build-and-push
build-and-push: ## Build and push the autograding image - this task is used by the github actions
	docker build -t $(DOCKER_IMAGE_NAME) $(CURRENT_PATH)
	docker push $(DOCKER_IMAGE_NAME)

.PHONY: install
install: start-helm-repo-locally upload-charts-locally install-charts cleanup ## Install the autograding helm charts

.PHONY: start-helm-repo-locally
start-helm-repo-locally:
	docker run --rm --name chartmuseum -d \
	  -p 8080:8080 \
	  -e DEBUG=1 \
	  -e STORAGE=local \
	  -e STORAGE_LOCAL_ROOTDIR=/chartstorage \
	  -v $(CURRENT_PATH)/chartstorage:/chartstorage \
	  ghcr.io/helm/chartmuseum:v0.16.0 &

.PHONY: upload-charts-locally
upload-charts-locally:
	curl --retry 10 --retry-connrefused http://localhost:8080 > /dev/null
	helm repo add $(CHART_REPO_NAME) $(CHART_REPO) || true
	cd $(CHART_PATH) && helm package .
	cd $(CHART_PATH) && curl -X "DELETE" $(CHART_REPO)/api/charts/$(CHART_NAME)/$(CHART_VERSION)
	cd $(CHART_PATH) && curl --data-binary "@$(CHART_NAME)-$(CHART_VERSION).tgz" $(CHART_REPO)/api/charts
	cd $(CHART_PATH)  && rm -rf *.tgz

.PHONY: install-charts
install-charts:
	helm repo update
	helm uninstall $(CHART_NAME) || true
	sleep 10
	helm install $(CHART_NAME) $(CHART_REPO_NAME)/$(CHART_NAME)

.PHONY: watch-autograding-job
watch-autograding-job:
	@echo "Waiting for job output.."
	@slept=0
	@while ! kubectl -n multi-tenancy-autograding get pod -l job-name=multi-tenancy-acceptance-tests | grep -q "Running"; do \
		sleep 1; \
		((slept=slept+1)); \
		if [ $$slept -eq 15 ]; then \
			echo "15 Seconds have elapsed and the job has not run -- something is wrong, exiting.."; \
			break; \
		fi; \
	done; \
	kubectl -n multi-tenancy-autograding logs -f job/multi-tenancy-acceptance-tests

.PHONY: cleanup
cleanup:
	# Stops the chartmuseum repo for the helm charts
	docker stop chartmuseum