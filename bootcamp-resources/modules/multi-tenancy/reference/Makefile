registry = localhost:5000
hnc_version = v1.0.0

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

##### Public tasks

.PHONY: bootstrap
bootstrap: ## Starts minikube and installs pre-requisites in the cluster
	./bootstrap/bootstrap-cluster.sh

.PHONY: run
run: ## Run the multi tenancy reference implementation pod in the cluster- this will create all resources needed
	@eval $$(minikube docker-env) ;\
	docker build -t coreeng/multi-tenancy-bootcamp-reference reference
	kubectl delete ns multi-tenancy-reference || true
	kubectl create ns multi-tenancy-reference
	kubectl apply -f reference/manifests
	kubectl run -n multi-tenancy-reference -it ref \
		--image coreeng/multi-tenancy-bootcamp-reference \
		--image-pull-policy="Never" \
		--command -- ./run.sh
	kubectl delete ns multi-tenancy-reference

##### Private tasks - for debugging and development

.PHONY: validate-onboarding-manifests
validate-onboarding-manifests:
	@echo "Evaluating all cue manifests"
	cd reference && cue eval ./onboard
	@echo "Print all generated yaml manifests"
	cd reference && cue cmd print ./onboard

.PHONY: test-tenant-rbac
test-tenant-rbac:
	# Test that tenants cannot access each other's namespaces
	cd reference && cue cmd testRbac ./onboard/tests

.PHONY: test-network-isolation
test-network-isolation:
	# Tests that the default deny applies correctly
	cd reference && cue cmd testNetworkIsolation ./onboard/tests

.PHONY: test-cilium
test-cilium:
	kubectl create ns cilium-test
	kubectl apply -n cilium-test -f https://raw.githubusercontent.com/cilium/cilium/v1.9/examples/kubernetes/connectivity-check/connectivity-check.yaml
	kubectl get pods -n cilium-test --watch