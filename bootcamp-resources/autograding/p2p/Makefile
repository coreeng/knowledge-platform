AUTOGRADING_P2P_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
AUTOGRADING_ROOT_PATH := $(realpath $(AUTOGRADING_P2P_PATH)/..)

.PHONY: help-p2p-autograding
help-p2p-autograding:
	@grep -E '^[a-zA-Z0-9_-]+:.*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: autograde-p2p
autograde-p2p: build-p2p-autograding cleanup-p2p-autograding deploy-p2p-autograding watch-p2p-autograding-job ## Run p2p autograding

.PHONY: build-p2p-autograding
build-p2p-autograding: # Build the autograding image
	@eval $$(minikube docker-env) ;\
	docker build -t p2p-acceptance-tests $(AUTOGRADING_P2P_PATH)


.PHONY: deploy-p2p-autograding
deploy-p2p-autograding: # Deploy the autograding resources
	kubectl apply -f $(AUTOGRADING_P2P_PATH)/manifests/namespace.yaml
	kubectl apply -f $(AUTOGRADING_P2P_PATH)/manifests/rolebindigs.yaml
	kubectl apply -f $(AUTOGRADING_P2P_PATH)/manifests/job.yaml
	kubectl apply -f $(AUTOGRADING_ROOT_PATH)/pushgateway/namespace.yaml
	kubectl apply -f $(AUTOGRADING_ROOT_PATH)/pushgateway/service.yaml
	kubectl apply -f $(AUTOGRADING_ROOT_PATH)/pushgateway/deployment.yaml

.PHONY: cleanup-p2p-autograding
cleanup-p2p-autograding: # Cleanup the autograding jobs if they exist
	kubectl -n p2p-autograding delete job p2p-acceptance-tests --ignore-not-found

.PHONY: watch-p2p-autograding-job
watch-p2p-autograding-job: # Watch for the output of the autograding job
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