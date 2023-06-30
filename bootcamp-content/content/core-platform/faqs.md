+++
title = "FAQs"
date = 2022-12-28T12:50:10+02:00
weight = 2
chapter = false
pre = "<b>3.2 </b>"
+++

#### What are the key elements of the Core Platform that make it different to a managed service like EKS or GKE?
A Core Platform is typically built on top of EKS and GKE but is so much more. A managed control plane (GKE, EKS) is about 1-5% of what a Core Platform is. It is a host of other services that offer:
* Deployed in a highly available way with considered upgrades vs a managed platform like GKE set to autopilot / upgrade braking tenant applications
* It is opinionated, offering a paved path for tenant engineers a.k.a product teams
* Be corporate connected. Routable in and out of corporate networks and cloud providers used
* Is multi-tenant. See [Why are Core Platforms multi-tenanted?](/core-platform/faqs/#why-are-core-platforms-multi-tenanted)
* Have an ingress solution that scales for the client (GKE/AWS LB Controllers are good for cluster edge load balancing but not for tenant ingress)
* Principled approach to deploying / upgrading / testing.

For example, a recent core platform build had the following deployed along with a managed control plane:
* JSPolicy: With the companies’s corporate standards encoded as policies
* Custom validating web hooks for the companies corporate standards that couldn’t be implemented using JSPolicy
* Ingress solution using AWS LB Controller for cluster edge + highly available Traefik per tenant
* and much much more

#### Why are Core Platforms Multi-Tenanted?
We advocate for multi-tenancy until there is a real blocker e.g. regulatory.  Why?

* A Core Platform team typically has full autonomy in a cluster, whereas inter cluster can involve other teams/departments such as networking
* In cluster communication is more feature full
    * Service Discovery
    * Network Policies
* Operational simplicity
    * Managing clusters is not an easy task, regardless of how much automation there is around deployment
    * Upgrades require careful consideration, better to do this once
    * Version drift between clusters

Some of these things are technically possible across cluster but typically much harder to setup/maintain.

And some more technical reasons:

* Getting corporate IP space - a single multi tenanted cluster per environment can get a corporate IP range once and then never depend on corporate networking changes again.

That said we do work with clients who want more of a “Core Platform” vending machine but there is still always some level of multi tenancy e.g. teams in the same department
