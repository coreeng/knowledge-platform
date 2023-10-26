.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

##### Installation tasks

.PHONY: bootstrap
bootstrap: ## Starts minikube and installs pre-requisites in the cluster
	./bootstrap/bootstrap-cluster.sh

.PHONY: autograde
autograde: ## Run autograding for the module "multi-tenancy"
	helm uninstall multi-tenancy-autograding --wait || true
	helm install --set module=multi-tenancy multi-tenancy-autograding oci://ghcr.io/coreeng/cecg-bootcamp-autograding
	@echo "Waiting for the jobs output..."
	@kubectl wait pods -n multi-tenancy-autograding --for condition=Ready --selector=job-name=multi-tenancy-acceptance-tests || true
	@kubectl -n multi-tenancy-autograding logs -f job/multi-tenancy-acceptance-tests

.PHONY: run
