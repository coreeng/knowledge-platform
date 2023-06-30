+++
title = "Tenant Definition"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation

Have a clear definition of what a tenant is on this platform.

How do we need to represent tenants so they will be recognized in the client ecosystem?

Decide if the tenant definition is the same for P2P and core platform (that would provide the best interface).

Can this become the single definition for all sub-systems including aws accounts etc?

 ## Requirements

Define the tenant definition, including decisions on:

Hierarchy of tenant access, documented as ADR

What integrations are required within that hierarchy  and with other external systems, and should there be any automation?

Any integration with CMDB?

Migration or creation of identities in AD groups, or other systems: Do we use existing ones that the tenants have or are new ones created specific to the platform.

How tenants map to P2P access and Core Platform access

 

## Questions

How does this interact with other systems?

Is there any existing layout that exists in service now that this should match to?

Does it map to org structure? Is a single application or a team that manages many applications.

As an example should we map out one of our tenants? E.g.

Do we need to capture the department / area? Do any approvals need to happen at the department / area level?

Consider the interactions that need to happen within the tenant hierarchy to facilitate/achieve a self service model (e.g. request for capacity, limits, etc..)

Can this depend on any standard labels/tags that go onto things like account requests?

 

