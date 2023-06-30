+++
title = "IaC Setup"
weight = 4
chapter = false
+++

##  Motivation

* Learn about the steps involved in setting up an IaC P2P.
* Learn about basic cloud provider setup.
* Learn best practices for authentication with a cloud provider from a P2P tool.
* Learn how CECG recommend bootstrapping new projects by doing an infrastructure P2P that deploys a small cloud resource in Sprint 0.

## Requirements

* A new private GitHub repository in the same organisation for your IaC project named `bootcamp-core-platform-<githb handle>`
* Initialisation step documented in the README that creates a state bucket and any service account required for the IaC P2P
* All IaC tools should keep the remote state in a storage bucket
* README with instructions on how to run the IaC locally and a description of any CI/CD you setup
* P2P
    * Repo should protect the main branch
    * PRs should be required to merge to main
* CI 
    * On PR, run a plan, and linting
* CD 
    * On merge to main deploy the latest
* Create a VPC with the CIDR range: 10.0.0.0/24 for nodes
    * Region: europe-west2
* The GCP project should be configurable

## Additional Information / Questions / Defuzz

You should receive a GCP sandbox project or AWS account.
Pick an IaC tool during planning or the epic defuzz
Ensure no secrets such as API keys are stored in the repo
Select a CI/CD tool: Tekton or Github Actions