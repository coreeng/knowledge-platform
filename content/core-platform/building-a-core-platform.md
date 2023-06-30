+++
title = "Building a Core Platform"
date = 2022-12-28T12:50:10+02:00
weight = 1
chapter = false
pre = "<b>3.1 </b>"
+++

Large infrastructure and platforming projects have a reputation for failing. Often built in isolation by
infrastructure teams hoping that delivery teams will use it.

In our view the definition of success of a platforming project is:
* Adoption: Delivery teams want to use the platform
* Speed of delivery: The platform speeds up product delivery for an organisation

Assuming no existing Core Platform. Our recommended process for building a new one with a client is:

{{<mermaid align="left">}}
graph LR;
  A(Definition) --> B(Build It: MVP with first tenants)
  B --> C(Launch & Adopt)
  C --> D(Evolve)
{{< /mermaid>}}

## Definition

The definition phase will start with 2-3 platform engineers working with the the organisation's infrastructure team and product tenants. The outputs are:

* Product Definition: MVP and Target
* Technical Architecture: MVP and Target
* Roadmap for the build it phase 
* Very high level roadmap for evolve phase

## Product Definition: MVP and Target

One size does not fit all. First we need to define what the minimum viable product is to unlock delivery and meet the organisation's initial go live requirements. 

An MVP phase in our experience can range from:
* 3 months: 2-3 Delivery teams/tenants in the same department
* 6 months: 3-4 Delivery teams/tenants spread across 1-2 departments
* 9 months: Enterprise ready for 3-4 Delivery teams/tenants in different parts of the organisation.

Within an MVP we recommend having a dev ready milestone to onboard the initial tenants into a dev environment to:
* Unblock them if they are bottlenecked on infrastructure
* Get their feedback

### Product Definition and Roadmap for MVP phase 

A platform is broken down into [features](/core-platform/features/). 
Features are grouped into high level categories called roadmap items.

Roadmap items are high level areas of the platform that can be understood by high level stakeholders. Each roadmap item is made up of many features. The details of the tasks for each epic will vary for each build. 

The current list of recommended roadmap items is:

* [Platform P2P](/core-platform/features/platform-path-to-prod/)
* [Connected Kubernetes](/core-platform/features/connected-kubernetes/)
* [Connectivity](/core-platform/features/connectivity/)
* [Multi Tenant Ingress + DNS](/core-platform/features/ingress/)
* [Multi Tenant Access](/core-platform/features/multi-tenant-access/)
* [Governance](/core-platform/features/governance/)
* [Kubernetes Networking](/core-platform/features/kubernetes-network/)
* [Security](/core-platform/features/platform-security/)
* [Platform Observability](/core-platform/features/platform-observability/)
* [Tenant Observability](/core-platform/features/tenant-observability/)

During the definition phase the scope of each roadmap item will be defined. 

A second phase roadmap can also be created to show what the key features that should be worked on after initial go live. For a second phase roadmap items can be an extension of the standard ones, with increased scope, and custom client specific ones. Defining this upfront puts stakeholders at ease that are vested in features that don't make the MVP.

