all: help

.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: run-local
run-local: ## Build and run docker container locally
	docker build . -t local-tp && \
	  docker run -p 8080:8080 \
	    --volume ./:/site \
	    --env LIVE_RELOAD=true\
	    --env BASE_URL=http://localhost:8080\
	  	local-tp
