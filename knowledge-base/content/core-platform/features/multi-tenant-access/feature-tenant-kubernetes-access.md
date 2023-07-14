+++
title = "Multi-Tenant Kubernetes Access"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation

* Support multiple tenants in the same Kubernetes cluster to reduce the operational overhead
* To be able to onboard new tenants (where tenants could be: departments, teams, or just a service) quickly i.e. not providing heavy weight, time consuming processes such as allocating network ranges, getting cloud accounts
* Decouple tenants from underlying infrastructure 

## Requirements

* Tenant definition schema agreed
* Tenant definition is separate from the Kubernetes resource creation (e.g. a separate sub module)
* Tenant onboarding process
* Tenant access constraints agreed
  * What is a tenant allowed to do in k8s?
* One tenant shouldn't be able to access the resources of other tenants
* Tenant resources should get metadata applied 
* (often) Namespace creation for tenants
* (often) Tenants can’t do cluster scoped operations

## Questions

How will tenants onboard? Options:

* PR to a config repo (could be the same repo, different folder)
* Via custom backstage
* Via a central request system such as Service Now

How will tenant resources (the implementation e.g. HNC resources, namesapces) be deployed?

* Is it it in the same lifecycle as the platform IaC code base? We highly recommend against this due to the coupling of the IaC deployment with the tenant resources deployment.

What happens when we onboard a tenant? 

* Does it include anything outside of Kubernetes? If so, are these better created by a k8s operator rather than the low level IaC layer.

What is the tenant definition? It is the inforation to describe a single tenant. Typically it will include:

* tenant name - this could be a team or a department, or even one per service
* tenant applications
* governance metadata e.g. cost centre, service now resolver

Should the tenant definition be agnostic of any technology? 

* E.g. a custom yaml that is then used to generate the artifacts (namespaces)

What is the tenant topology? How many levels does it have?

* What needs to be represented: Orgs, Departments, Teams, Applications?

Should tenants have a mechanism to create namespaces? E.g. with HNC. Or is it a request to the platform team?

What technology should tenant segregation be implemented with?

* Pure namespaces
* HNC
* Vclusters


## Additional Information

A typical module structure for implementing is:

* `tenant-definition`
  * Technology agnostic definition of tenants
* `connected-kubernetes`
* `multi-tenant-kubernetes`


* Typically depends on Cluster Resource Creation

For tenant access we want to give them as much autonomy as possible without breaking the platforming principle of: one tenant should not be able to adversely affect another.

Our current best practice is to use the [Hierarchical Namespaces](https://github.com/kubernetes-sigs/hierarchical-namespaces) that way each tenant can have their own namespace parent with the permission to create sub namespaces, this:

* Decouples the platform team from knowing about any of the team’s applications

It has some downsides:

* Makes it harder to create standard environments e.g. for functional tests

A middle ground is to create the standard environments as namespaces in non-prod clusters and allow any other namespaces to be created with the HNC.

