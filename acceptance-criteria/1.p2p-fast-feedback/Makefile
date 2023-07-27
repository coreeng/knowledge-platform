.PHONY: run-p2p-fast-feedback-ac-test
run-p2p-fast-feedback-ac-test: build-image cleanup-jobs apply-manifests wait-and-print-job-output

.PHONY: build-image
build-image: ## Builds the image using minikube's docker
	@eval $$(minikube docker-env) ;\
	docker build -t p2p-fast-feedback-ac .

.PHONY: apply-manifests
apply-manifests: ## Applies the necessary manifests for the acceptance criteria to run
	kubectl apply -f $(CURDIR)/manifests/namespace.yaml
	kubectl apply -f $(CURDIR)/manifests/rolebindigs.yaml
	kubectl apply -f $(CURDIR)/manifests/job.yaml
	kubectl apply -f $(CURDIR)/manifests/service.yaml
	kubectl apply -f $(CURDIR)/manifests/pushgateway.yaml

.PHONY: cleanup-jobs
cleanup-jobs: ## Cleans up any leftover jobs in the cluster
	kubectl -n p2p-fast-feedback-ac delete job p2p-acceptance-criteria --ignore-not-found

.PHONY: wait-and-print-job-output
wait-and-print-job-output:
	@echo "Waiting for job output.."
	@slept=0
	@while ! kubectl -n p2p-fast-feedback-ac get pod -l job-name=p2p-acceptance-criteria | grep -q "Running"; do \
		sleep 1; \
		((slept=slept+1)); \
		if [ $$slept -eq 15 ]; then \
			echo "15 Seconds have elapsed and the job has not run -- something is wrong, exiting.."; \
			break; \
		fi; \
	done; \
	kubectl -n p2p-fast-feedback-ac logs -f job/p2p-acceptance-criteria