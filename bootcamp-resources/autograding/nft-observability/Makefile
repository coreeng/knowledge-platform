AUTOGRADING_NFT_OBS_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
AUTOGRADING_ROOT_PATH := $(realpath $(AUTOGRADING_NFT_OBS_PATH)/..)

.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: autograde
autograde: build cleanup deploy watch-autograding-job ## Run nft observability autograding

.PHONY: build
build: ## Build the autograding image
	docker build -t cecg/nft-observability-autograding $(AUTOGRADING_NFT_OBS_PATH)

.PHONY: build-and-push
build-and-push: ## Build and push the autograding image - this task is used by the github actions
	docker build -t cecg/nft-observability-autograding $(AUTOGRADING_NFT_OBS_PATH)
	docker push cecg/nft-observability-autograding

.PHONY: deploy
deploy: ## Deploy the autograding resources
	kubectl apply -f $(AUTOGRADING_NFT_OBS_PATH)/manifests/namespace.yaml
	kubectl apply -f $(AUTOGRADING_NFT_OBS_PATH)/manifests/rolebindigs.yaml
	kubectl apply -f $(AUTOGRADING_NFT_OBS_PATH)/manifests/job.yaml
	kubectl apply -f $(AUTOGRADING_ROOT_PATH)/pushgateway/namespace.yaml
	kubectl apply -f $(AUTOGRADING_ROOT_PATH)/pushgateway/service.yaml
	kubectl apply -f $(AUTOGRADING_ROOT_PATH)/pushgateway/deployment.yaml

.PHONY: cleanup
cleanup: ## Cleanup the autograding jobs if they exist
	kubectl -n nft-obs-autograding delete job nft-obs-acceptance-tests --ignore-not-found

.PHONY: watch-autograding-job
watch-autograding-job: ## Watch for the output of the autograding job
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