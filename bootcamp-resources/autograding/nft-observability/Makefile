ACCEPTANCE_CRITERIA_ROOT := $(realpath $(CURDIR)/..)

.PHONY: run-nft-and-obs-ac-test
run-nft-and-obs-ac-test: build-image cleanup-jobs apply-manifests wait-and-print-job-output

.PHONY: build-image
build-image: ## Builds the image using minikube's docker
	@eval $$(minikube docker-env) ;\
	docker build -t nft-and-obs-ac .

.PHONY: apply-manifests
apply-manifests: ## Applies the necessary manifests for the acceptance criteria to run
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/nft-observability/manifests/namespace.yaml
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/nft-observability/manifests/rolebindigs.yaml
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/nft-observability/manifests/job.yaml
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/pushgateway/namespace.yaml
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/pushgateway/service.yaml
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/pushgateway/deployment.yaml

.PHONY: cleanup-jobs
cleanup-jobs: ## Cleans up any leftover jobs in the cluster
	kubectl -n nft-and-obs-ac delete job nft-acceptance-criteria --ignore-not-found

.PHONY: wait-and-print-job-output
wait-and-print-job-output:
	@echo "Waiting for job output.."
	@slept=0
	@while ! kubectl -n nft-and-obs-ac get pod -l job-name=nft-acceptance-criteria | grep -q "Running"; do \
		sleep 1; \
		((slept=slept+1)); \
		if [ $$slept -eq 15 ]; then \
			echo "15 Seconds have elapsed and the job has not run -- something is wrong, exiting.."; \
			break; \
		fi; \
	done; \
	kubectl -n nft-and-obs-ac logs -f job/nft-acceptance-criteria