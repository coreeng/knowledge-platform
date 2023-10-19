# 1. ADR for Network Isolation

Date: 2023-04-06

## Status

In review.

## Context

In a multi-tenant environment the following principles apply:
- the tenants can not impact each-other accidentally and should work as it they were completely isolated. That’s true for resources as well as network. 
  For example, if a tenant pod is compromised, it should not be able to maliciously attack other pods/tenants, 
  which can be achieved with default deny, as you’ll need to whitelist all the connectivity needed for any given application.
- the tenants should have the ability to update the network policies for their own resources so that platform 
  team intervention is minimised. 

## Decision

Network isolation between tenants can be achieved by:
- creating a default deny network policy for ingress which applies to all namespaces in an organisation. 
  This is achieved using Cilium in combination with the HNC feature of network policy propagation.
  When the default deny policy is applied to the parent namespace of an organisation it gets propagated to all the 
  subnamespaces. 
- relying on the HNC functionality of network policy propagation means that tenants cannot go and delete the policy in their 
  subnamespaces, but they have the rights to create new rules and whitelist access to let's say the monitoring namespace.


## Consequences

- Cilium needs to be installed and maintained
- Cilium UI can be used to visualise the connections in the cluster
- Teams will also need to use Cilium to apply their own network policies

