+++
title = "Platform: Single Lifecycle"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation

* Reliably deploy a version of a platform to a single environment

## Requirements

* **Source control**: single repo for the source of all platform components
* **Packaging**: Immutable artifact used to deploy the platform
* **Versioning**: Every deployable artifact has a unique, incrementing, version
* **Deployment**: Given a versioned artifact and a target environment (e.g. a cloud account) the platform can be deployed
   * Promotion implemented in either [basic promotion](./feature-basic-promotion) or [automatic promotion](./feature-automatic-promotion)

## Additional Information

Advantages:
* Easy to understand.
* Easy to deploy a whole new version as it is a single life cycle.

Downsides:
* Longer deployments
* Larger blast radius

### Loose coupling

A single lifecycle does not mean a monolithic architecture e.g. everything coupled together. Through many platform builds
we've seen the advantage of having a platform be made up of loosely coupled components due to:
* Each can be deployed in sequence (or part of a DAG) so that deployment failure is easily identifiable to a single component
* Tests can be separated and run after the deployment to verify that feature rather than all at the end

## Questions

What should be done to prove this feature at the start of a platform build?
* Deploy one piece of cloud infrastructure 

What is the package for a platform?
* Docker image including deployment dependencies e.g. terraform, gcloud
* A git tag
* An archive that is 

## Depends on

* At least one environment from [Core Platform Environments](./feature-core-platform-environments)
