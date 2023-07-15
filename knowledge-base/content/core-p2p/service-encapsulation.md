+++
title = "Service Encapsulation"
date = 2022-12-26T19:32:10+02:00
weight = 2
chapter = false
+++
Our development recommendations are geared towards enabling Continuous Deployment.

[Trunk based development](https://cloud.google.com/architecture/devops/devops-tech-trunk-based-development) is a key capability identified by the [DORA research program](https://www.devops-research.com/research.html#capabilities). 

## Single Repository 
Everything required to build, deploy, and verify a service should be in a single repository, versioned together. This includes:

* Service code
* Deployment manifests e.g. Kubernetes YAML files
* Functional Tests
* Non-Functional Tests
* Configuration 
* Monitoring configuration: dashboards, alerts

## Trunk Based Development
A single main aka trunk branch that all engineers for a service work off. [Short lived branches](https://trunkbaseddevelopment.com/short-lived-feature-branches/) for feature work that are quickly merged into the main branch via a [Pull Request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests). 

Read through [Trunk Based Development](https://trunkbaseddevelopment.com/) 

## Immutable Artifacts
The main branch should be built via CI tooling into an immutable artifact such as:

* Docker Image
* Zip file (e.g. for serverless)
* RPM etc

And then promoted through environments. It should not be re-built for each environment, the artifact that is tested in one environment should be promoted to the next.
