ACCEPTANCE_CRITERIA_ROOT := $(realpath $(CURDIR)/..)

.PHONY: run-p2p-fast-feedback-ac-test
run-p2p-autograding: build-image cleanup-jobs apply-manifests wait-and-print-job-output

.PHONY: build-image
build-image: ## Builds the image using minikube's docker
	@eval $$(minikube docker-env) ;\
	docker build -t p2p-acceptance-tests .

.PHONY: apply-manifests
apply-manifests: ## Applies the necessary manifests for the acceptance criteria to run
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/p2p/manifests/namespace.yaml
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/p2p/manifests/rolebindigs.yaml
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/p2p/manifests/job.yaml
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/pushgateway/namespace.yaml
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/pushgateway/service.yaml
	kubectl apply -f $(ACCEPTANCE_CRITERIA_ROOT)/pushgateway/deployment.yaml

.PHONY: cleanup-jobs
cleanup-jobs: ## Cleans up any leftover jobs in the cluster
	kubectl -n p2p-autograding delete job p2p-acceptance-tests--ignore-not-found

.PHONY: wait-and-print-job-output
wait-and-print-job-output:
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