+++
title = "Canary Operator"
weight = 4
chapter = false
+++

## Motivation

Provide a better tenant interface by building deployment models directly into the platform rather than external scripts

* A single CRD for tenants to update rather than many Kubernetes resources


More control for the platform team for upgrading tenant resources in place rather than relying on the tenants to upgrade tooling

* E.g. as the underlying Kubernetes resources are created by a platform operator, it is easier to do things like upgrade Kubernetes as the platform operator can upgrade any underlying resources rather than waiting for tenants to do it 

## Requirements

We'll be implementing the same requirements as in the *core-platform* Canary Deployments epic but this time
inside the platform as an operator and CRD! The deployment steps should be as follows:

* Deploy 1 new replica with the new version
* Wait until that replica is healthy
* Enable that replica to receive traffic through Ingress
* Have some type of check that needs to stay healthy for a configurable amount of time, e.g. checking a metric
* Once check has passed for the configured amount of time replace the remaining replicas with the new version if the new replica is healthy

The user interface should be a single deployment of the Custom Resource,
the user should not have to directly manipulate: 

* Ingresses
* Services
* Deployments
 
The CRD should contain the information required to create all of these for the user e.g.

* Number of replicas
* Image

## Questions / Defuzz / Decisions

See the [Operator Pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)
We recommend using [Kube Builder](https://book.kubebuilder.io/)
Your operator should have unit and integration tests

## Deliverables (For Epic)

* ...
