+++
title = "Multi Tenancy: RBAC"
weight = 5
chapter = false
+++

## Motivation

* Work around the k8s limitations for implementing RBAC in a multi tenanted platform
* Have an onboarding mechanism for tenants to get their RBAC/Namespaces setup
* Decouple the platform team from the tenant teams

## Requirements

* Everything to meet these requirements should be in a separate repo you create in the organisation `bootcamp-multi-tenancy-<github handle>`
* Tenants can onboard to the platform with a single action e.g. a PR

* We're going to implement the following [RBAC](/core-platform/features/multi-tenant-access/) features:

  * [Tenant Kubernetes access](/core-platform/features/multi-tenant-access/feature-tenant-kubernetes-access/) (partially)
  * [Cluster Resource Creation](/core-platform/features/multi-tenant-access/feature-cluster-resource-creation)

### Specific requirements for the features above

#### Tenant Kubernetes Access

A tenant definition (i.e. the thing that they are PRing) contains:
* Tenant name
* List of application names
* Cost code (numeric) - this is just an example of mandatory metadata that we want to associate with all resource for a tenant

There can be many tenants. 

Validation step - this is the step run on CI:
* Ensure there can't be two tenants with the same name
* Any input provided by the tenant that becomes a resource name is valid

Deployment step - this is the step run on CD:
* Convert the tenant configuration into k8 resources using [hierarchical namespaces](https://kubernetes.io/blog/2020/08/14/introducing-hierarchical-namespaces/).

E.g. the output from hns

```
kubectl hns tree cecg                                                                                                                                                         
cecg (whole org)
├── tenant-a 
│   ├── app-1 <- deploy your reference application here
│   ├── app-2
│   └── team-a-monitoring
└── tenant-b 
    ├── app-3 <- deploy your reference application here 
    └── team-b-monitoring
```

* Each team should get its own monitoring stack as installed in "Monitoring Stack Setup"
  * If your local Kubernetes Setup is resource limited, limit this to just team-a 
* Each team should be deploying a version of the reference application into their own namespace
* Everything you create / install should be in code, i.e. you should be able to blast away and start from scratch with a new minikube cluster
  * Repo for all the kubernetes resource creation can be public in your GitHub profile

* Keep everything for a specific application in its repository

#### Cluster Resource Creation

- The functionality to sync the resource creation with the actual infrastructure is out of scope. 

## Additional Information

* Multiple Kubernetes clusters will be used in practice between non-prod and prod, but doing it all in the same cluster teaches the concepts
* De-scope any quotas 
* In this epic you'll be adopting three roles:
  * Platform engineering team member: working on platform onboarding service 
  * Team A member - deploying app 1 and 2
  * Team B member - deploying app 3

## Questions / Defuzz / Decisions

As a new tenant to a platform. How is that tenant onboarded? Is it a PR? A service now request? 

Should the tenants be tightly coupled to the implementation i.e. do they know that HNC is used? I.e. what is in the PR? 

How will the functionality be tested? 

Does your solution work if a tenant is removed or updated? 

How would this be deployed for the bootcamp versus in a more realistic setup?

* Bootcamp:

  * One command that validates the input e.g a `make` task
  * One command that installs all the resources e.g. a `make` task 
  * Tests should be run with one `make` command locally
* More realistic:

  * CI: Validates any configuration against a schema. Versions the configuration. 
  * CD: Deploys the resources to the first environment then promotes through other environments. Quality Gates are in place. 

Key decisions requiring [ADRs](/core-engineer/adrs/):

* How are Kubernetes resources deployed?
