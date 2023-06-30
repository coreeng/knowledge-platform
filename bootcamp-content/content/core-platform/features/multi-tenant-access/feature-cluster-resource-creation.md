+++
title = "Cluster Resource Creation"
date = 2022-01-15T11:50:10+02:00
weight = 3
chapter = false
pre = "<b></b>"
+++

### Motivation

* Ability for the platform engineering team to deploy cluster scoped resources across clusters 
  * Typically used to deploy RBAC for tenants, namespaces, any CRDs
* Ability for tenants to request the deployment of cluster scoped resources

### Requirements

* Ability to define which clusters get which resources
* Ability to define which namespaces get which resources
* Ability to update existing resources idempotently
* Ability to delete resources that are removed

### Additional Information

The building block for tenant kubernetes access.

Many features depend on a central platform function that created cluster-wide resources such as namespaces as well as all the roles and bindings to give tenants access.

There is an expanding open source ecosystem for this and each new core platform should evaluate new technologies in this space. Existing approaches have been:

* Terraform
* GitOps tools such as [Flux](https://fluxcd.io/), config-sync.




