+++
title = "Reference Application"
weight = 4
chapter = false
+++

## Motivation

Learn about the steps for a mature path to production via the bootcamp reference application.

## Requirements

* Copy the  Java (bootcamp-resources/skeletons/reference-application-java) or Go (bootcamp-resources/skeletons/reference-application-go) bootcamp reference applications as a private repo in the same organisation with the name `reference-application-(java|go)-<github handle>`
* Checkout and verify workstation setup by running it locally, including building, running functional and non-functional tests 
* Update the registry to one that can be used by your local cluster
* Deploy the application to a local Kubernetes cluster and run the functional tests 
* Complete the CI requirements for the reference application
  1. Branch protected for the `main` branch 
      * No force pushes to main branch
  2. Pull-requests are required for merging to main branch
      * At least one peer review :eyes:
      * “Local testing” taken place :test_tube:
  3. On merge to main branch, the following should be enforced:
      * Versioned image built
      * Deployed functional tests run
      * Deployed NFT tests run


## Questions / Defuzz / Decisions

There are two technologies that can be used for the P2P:
* GitHub Actions
* Tekton

Agree with your product owner which to use for this epic.

**Q: What is "local" testing?** Any test that can be run without the need for a real environment. See Testing Strategy.

**Q: What is deployed testing?** See Deployed vs non-Deployed tests

**Q: What registry can be used?**
* Dockerhub
* Any within your organisation that your local cluster has access to
* Push directly to docker in minikube: https://minikube.sigs.k8s.io/docs/handbook/pushing/#1-pushing-directly-to-the-in-cluster-docker-daemon-docker-env

### Tekton

Get a locally running [Tekton Cluster for this epic](https://tekton.dev/docs/getting-started/).
For deployed testing the application can be deployed to your local cluster.

### GitHub Actions

Example GitHub actions are in the .github folder in the reference application.
For deployed testing the application can be deployed to a minikube in [github actions](https://minikube.sigs.k8s.io/docs/tutorials/setup_minikube_in_github_actions/)

### Makefile Documentation

The Java and Go reference applications both make use of a Makefile for build tasks.  For consistency and compatibility 
with the consolidated P2P, the Makefile should contain the following targets:

* `build` - Build the reference application
* `docker-build` - Build the reference application, and package as a Docker image
* `docker-push` - Push the Docker image to an image registry
* `local` - Build the reference application, and run functional and non-functional tests locally
* `local-stubbed-functional` - Run functional tests locally
* `local-stubbed-nft` - Run non-functional tests locally
* `run-local` - Run the reference application locally
* `stubbed-functional` - Perform functional tests against stubs
* `stubbed-nft` - Perform non-functional tests against stubs
* `extended-stubbed-nft` - Perform extended non-functional tests against stubs
* `integrated` - Perform integrated tests

The Makefile also contains a `help` target which will generate the usage documentation shown above.  This relies upon 
Makefile targets being documented with a comment beginning `##` after their declaration.

## Deliverables (For Epic)

- ... 
