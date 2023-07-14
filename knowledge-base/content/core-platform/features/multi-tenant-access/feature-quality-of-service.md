+++
title = "Quality of Service and Resource Quotas"
date = 2022-01-15T11:50:10+02:00
weight = 5
chapter = false
pre = "<b></b>"
+++

## Motivation
* To provide fair use of a shared multi-tenanted cluster for those that use it.
* If tenants have service tiers within their cluster, provide mechanisms to appropriately provide service per tier.
* To be able to prioritise the bandwidth needs of different tenants across a multi-tenanted cluster.
* To be able to prioritise the storage needs of different tenants across a multi-tenanted cluster.
* To be able to prioritise pod scheduling needs of differenbt tenants across a multi-tenanted cluster.


## Requirements
* Implement mechanism to allow rate limits to pods.
* Introduce templates for storage classes tenants can define and implement
* Implement a pod priority mechanism which allows tenants to schedule pods based on priority.

## Questions

Does the client need a service tier for it's tenants?

## Additional Information
