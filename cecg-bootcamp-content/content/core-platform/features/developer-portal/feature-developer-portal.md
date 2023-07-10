+++
title = "Developer Portal"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation

* Provide a single UI to all features of the core platform for tenants
* Hands free deployment of a Core Platform backstage instance

## Requirements

* Backstage instance in Version Control
  * Branded as Core Platform
  * Organisation as CECG
* Backstage deployed to the core platform
  * Ingress + TLS
* Database deployed via a cloud managed service 
* VCS integration enabled
* Kubernetes integration configured for the current cluster with the option to discovery multiple clusters
  * For any service entity, annotations in place for the Deployments, Ingress, 
* Integrated with access management i.e. no additional managing of users and groups specifically for the developer portal
* Potential tenants can request access
* Basic core platform technical documentation

## Questions

Should the backstage instance source code be in the same repo as the core platform or a separate one?

Can the developer portal be deployed to a non-core platform? Can it be a self contained module?

* Which components does it depend on? Just [Connected Kubernetes](../connected-kubernetes/) or [Multi Tenant Access](../multi-tenant-access/)?
* If it depends on both, what is the interface between the two? How does that affect backstage being used as the onboarding interface for tenants?

Which environment and how many instances of backstage?

How is backstage configured, packaged and deployed? 

* Is it part of the platform P2P or a fully separate lifecycle?

How is backstage kept up to date?

How does authentication and authorisation work?

Where will the database live and how will it be provisioned? It can depend on [Paired Cloud Account](../persistence/feature-paired-cloud-account) or live in the platform accounts.

How does authentication work got GitHub? 

* What type of account?
* What permissions does it have?
* Can it use OIDC?

What should be in the initial technical documentation?

* Onboarding docs?
