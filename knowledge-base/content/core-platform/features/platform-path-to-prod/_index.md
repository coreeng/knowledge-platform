+++
title = ""
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b>Platform Path to Prod </b>"
+++

# Platform Path to Prod

This section describes how to manage the path to production for the Kubernetes platforms we manage and build. It is primarily our approach to quality that mandates Platform P2P features. However, if there are requirements for very quick rebuilds / many instances of a Core Platform it can influence the selected features.

## What is a Platform P2P?

A Core Platform is typically made up of:

* IaC:
  * Base networking, IAM
  * Base Kubernetes, typically provisioned with a cloud provider e.g. EKS/GKE
* Kubernetes resources
  * external-dns
  * ingress controllers
  * policy controllers
* Custom controllers

The approach to CI/CD for platforms depends on:

* The expected size of the platform team
* The skill of the platform team members
* The size of the platform
* The number of instances of the core platform. Ranging from:
  * Single, multi-tenanted core platform per environment
  * A small number of instances 
  * Many core platforms e.g. one per business area / team

## Deployment models
 
Based on the above we categorise three main deployment models for platforms

### 1. Monolithic

A single lifecycle for all IaC and things deployed to the cluster.

#### Advantages

* Conceptually very simple. A single versioned deployment.
* Easy to build a new platform from scratch
* Everything can rely on a specific version of everything else

#### Disadvantages

* Longer deployment times
* Longer PR validation
* Larger blast radius for every deployment

### 2. Fully decoupled

Break the platform into decoupled components e.g.

* Base K8s layer
* Ingress Controllers
* Policy Agents

#### Advantages

* Low blast radius of deployments
* Quicker deployment times
* Quicker PR validation

#### Disadvantages

* Inevitable dependencies between components
* Conceptually much more difficult to understand what is running in any environment
* Can introduce duplication at the build/tooling layer
* Difficult to reliably integration test and promote known working versions together
* Harder to stand up new clusters - need to add new configuration in n places rather than a single

### 3. Middle ground

Have the majority of the platform deployed as one and key components, with large test lifecycles, pulled out to have their own life cycle.

 
### Continuous Integration for Platforms

Continuous integration is how we build and test the main branch. Generally we follow [Trunk Based Development](https://trunkbaseddevelopment.com/) , where main is expected to always be in a good state, ready for deployment at any time. This is required to enable CD.

With CI, every commit is tested before it can be merged into the main branch. If you’re familiar with application development, CI for apps is pretty straightforward - a PR or branch commit triggers a remote build, which runs a suite of tests. 

For platforms, the situation is similar, but more complex due to the heavy reliance on cloud components. Provisioning a platform usually means that the vast majority of what we do interacts with cloud resources - we can’t simply stub this out, since the tests would not be very good. As a result, we have to build a robust build and test platform that can create and test cloud resources in isolation.

Our main approach for handling this is to build dedicated, independent CI environments which can be deployed to and tested against.

Another difficulty is the time it requires to provision cloud resources can be quite lengthy. So we have to generally operate in non-idempotent states - we can’t simply wipe out our CI environments between test runs or the length of builds would get too long. To give an example, for a GKE platform where many CRDs are installed, it can take up to 60 minutes to provision and
GKE's control plane scaling to kick in. 

### Continuous Delivery for Platforms

We embrace CD for platform roll outs to ensure high quality, incremental releases are continually being deployed and used. This is in contrast to a more traditional, periodic manual release cycle where releases are done big-bang after large amounts of changes have been made. 

## Models for CI/CD

Depending on the maturity of the organisation and size of the project, there are different levels of automation that can be taken towards implementing CI/CD for a client.

### 1. CI with manually triggered deployments

#### When to use:

* We have infrastructure to stand up CI automation (Jenkins/Tekton).
* We have a fairly substantial team (3+ devs) where CI will benefit the day to day dev process and ensure a green main build.
* Have a cloud environment for running tests in.
* Don’t have the infra or time to implement CD. Or we only have a few clusters (<= 4) to manage.

In this model, we implement CI for our development process. This facilitates multiple engineers working concurrently on a code base, especially with a PR process where we can automate checks on PRs.

For CI:

* Stand up Jenkins, or alternatively a Tekton pipeline (if you have a pre-existing Kubernetes cluster, or use something like microk8s to stand up a no-ops cluster just for CI).
* Every build should deploy the platform to the test environment and run a suite of tests.
* Results are reported on PRs.

### 2. CI and CD

For Platform CD, this means that each version of the platform is automatically promoted through environments without any human interaction.
It is a full implementation of [automatic promotion](./feature-automatic-promotion).

#### When to use:

* Large teams/orgs and many clusters (> 4) are planned to be deployed.
* We have the bandwidth to dedicate serious effort into CD.
* We have infrastructure for implementing both CI and CD.

In this model, we implement full CI and CD.

For CI, follow (2).

For CD:

* Create a genesis cluster which will manage deployments of all other clusters and cloud resources.
* Use a CD tool like [Argo CD](https://argo-cd.readthedocs.io/en/stable/) or a custom controller 
* Figure out how you will promote between environments - what are the SLOs that indicate a specific platform version is okay to promote? Common ones are minimum deployment age and number of alerts.

