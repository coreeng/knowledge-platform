AUTOGRADING_NFT_OBS_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
AUTOGRADING_ROOT_PATH := $(realpath $(AUTOGRADING_NFT_OBS_PATH)/..)

help-nft-obs-autograding:
	@grep -E '^[a-zA-Z0-9_-]+:.*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: autograde-nft-obs
autograde-nft-obs: build-nft-obs-autograding cleanup-nft-obs-autograding deploy-nft-obs-autograding watch-nft-obs-autograding-job ## Run nft observability autograding

.PHONY: build-nft-obs-autograding
build-nft-obs-autograding: # Build the autograding image
	@eval $$(minikube docker-env) ;\
	docker build -t nft-obs-acceptance-tests $(AUTOGRADING_NFT_OBS_PATH)

.PHONY: deploy-nft-obs-autograding
deploy-nft-obs-autograding: # Deploy the autograding resources
	kubectl apply -f $(AUTOGRADING_NFT_OBS_PATH)/manifests/namespace.yaml
	kubectl apply -f $(AUTOGRADING_NFT_OBS_PATH)/manifests/rolebindigs.yaml
	kubectl apply -f $(AUTOGRADING_NFT_OBS_PATH)/manifests/job.yaml
	kubectl apply -f $(AUTOGRADING_ROOT_PATH)/pushgateway/namespace.yaml
	kubectl apply -f $(AUTOGRADING_ROOT_PATH)/pushgateway/service.yaml
	kubectl apply -f $(AUTOGRADING_ROOT_PATH)/pushgateway/deployment.yaml

.PHONY: cleanup-nft-obs-autograding
cleanup-nft-obs-autograding: # Cleanup the autograding jobs if they exist
	kubectl -n nft-obs-autograding delete job nft-obs-acceptance-tests --ignore-not-found

.PHONY: watch-nft-obs-autograding-job
watch-nft-obs-autograding-job: # Watch for the output of the autograding job
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