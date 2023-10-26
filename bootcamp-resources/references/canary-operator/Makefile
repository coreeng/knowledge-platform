.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: boostrap-cluster
bootstrap-cluster: ## Start the minikube local cluster
	./scripts/bootstrap-cluster.sh

.PHONY: install
install: ## Install the CR and the canary operator in minikube
	cd src && make install
	@eval $$(minikube docker-env); \
	cd src && make docker-build
	cd src && make deploy