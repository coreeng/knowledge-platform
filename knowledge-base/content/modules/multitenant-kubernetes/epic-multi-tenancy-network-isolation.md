+++
title = "Multi Tenancy: Network Isolation"
weight = 6
chapter = false
+++

## Motivation

* Have a CNI (Container Network Interface) that can enable us to achieve network isolation by using network policies.
* Provide tenants with access to their own firewalling - so that tenants have autonomy over accessing their own applications or applications maintained by other tenants.
* Learn about a key principle of a multi tenant kubernetes cluster: Default Deny.

## Requirements

The requirements should be implemented as features in the repo added as part of the Multi Tenancy RBAC Epic:
`bootcamp-multi-tenancy-<github handle>`.

Implement the following set of Kubernetes Network features:

* CNI
* Tenant Managed Firewalls
* Default Deny

### Specific requirements for the features above

#### CNI

* Use [Cilium](https://cilium.io/) or [Calico](https://docs.tigera.io/calico/3.25/about/). 

#### Default deny

Platform responsibilities - when a tenant is onboarded the following rules apply:
* Default deny for inbound network traffic
* Applications in the same namespace can, by default, communicate with each other
* Applications in different namespaces can, by default, not communicate with each other (Application 1 cannot communicate with application 3 - use the applications deployed in the multi tenancy RBAC epic)

Tenant responsibilities:
* Create namespace level network policies to override the default deny - so that monitoring continues to work

## Questions / Defuzz / Decisions

* Which CNI will be used?


## Deliverables 

- [ ] Switch CNI to Calico/Cilium - add an ADR to explain why a CNI was chosen over the other
- [ ] Network policies for default deny - add an ADR to support the decisions for implementing this
- [ ] Network policies for allowing prometheus connectivity

