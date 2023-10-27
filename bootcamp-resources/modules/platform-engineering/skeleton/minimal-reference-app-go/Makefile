REGISTRY := cecg
IMAGE_NAME := minimal-ref-app
IMAGE_TAG :=

.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the docker image for the app
ifndef IMAGE_TAG
	$(error "IMAGE_TAG variable needs to be defined")
endif
	@eval $$(minikube docker-env) ;\
	docker build -t $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG) .

.PHONY: build-and-push
build-and-push: ## Build and push the docker image for the app
ifndef IMAGE_TAG
	$(error "IMAGE_TAG variable needs to be defined")
endif
	docker build -t $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG) .
	docker push $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)

run:
	go run cmd/main.go