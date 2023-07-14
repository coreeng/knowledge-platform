+++
title = "Tenant Managed Firewalls"
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b></b>"
+++

## Motivation

* Delegate control of Network Policies to tenants while still maintaining control of key parts e.g. if default deny is required, egress firewalling
* Speed up delivery by not requiring platform engineering team involvement in network policy creation
 
## Requirements

* Enable NetworkPolicies

## Additional Information

* What policies are needed on the network policies the tenants create?
* Tenants can create network policies in their namespaces
  Platform can still manage any cluster wide network policy


