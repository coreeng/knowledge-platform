+++
title = "Local Kubernetes with Tekton"
weight = 5
chapter = false
+++

## Motivation

Have a local Kubernetes cluster to use throughout the bootcamp.
Have an isolated CI/CD tool installed with full control over its settings.

## Requirements 

* A local Kubernetes cluster to work with e.g. Minikube or Kind
* All configuration of the cluster should be in version control as scripts, helm charts etc in a private repo bootcamp-<github handle>
* Tekton installed 
  * Pipelines
  * Triggers
  * Dashboard
* Tekton CLI installed
* Tekton tasks installed:
  * [git-clone](https://hub.tekton.dev/tekton/task/git-clone) 
* Ngrok Installed to allow GitHub web hooks to access your local cluster

## Additional Information 

There are instructions for setting up Tekton in the reference application in [tekton/README.md](https://github.com/coreeng/reference-application-java/blob/main/tekton/README.md).

