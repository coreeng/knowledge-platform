+++
title = " "
date = 2022-01-15T11:50:10+02:00
weight = 4
chapter = false
pre = "<b>Multi-tenancy</b>"
+++

# Multi-Tenancy

Core Platforms typically have many departments onboarded with each department split into many teams. Traditionally, one might consider creating a separate platform per team or tenant to deploy and maintain their services. However the concept of multi-tenancy allows us to house multiple teams and even departments within a single kubernetes cluster, while respecting each tenant's limits, boundaries and autonomy. Taking this approach allows us to reduce the complexity and cardinality of our architecture by providing and maintaining a single feature-rich cluster as opposed to configurations for multiple clusters, the trade-off being that more time may be required up-front to replicate multi-cluster features within a single cluster. 

## Principles

Establish a set of principles with the client that tenants and platform engineers must strive to follow within a multi-tenanted cluster.

Some example principles that can be used as a foundation are:

* One tenant should not be able to adversely affect another tenant
* Developers should be given as much autonomy possible without breaking the other principles
* Non-developers/business areas/departments should be given as much autonomy possible without breaking the other principles
    - example: Can a department enforce a budget on their teams + applications?
    - Can departments approve new applications within their department 
* One tenant should not be able to see what another tenant is doing (optional)
    - i.e should you give read only access to one tenant to another tenants resources
    - this principle may affect the architecture i.e. what ingress controller is used

## Anecdotes
Some anecdotal problem points can be discussed with the client for relatability in highlighting problems and why multi-tenancy solves those problems. 

For example:
* If departments are in separate clusters with separate networks, occasionally there might be some governance that they have to go through requesting firewall rules for inter-team communication. Having the departments residing in a multi-tentanted cluster alleviates this extra step.






