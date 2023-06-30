+++
title = "Base Networking Design"
date = 2022-01-15T11:50:10+02:00
weight = 3
chapter = false
pre = "<b></b>"
+++

## Motivation

Understand how the network will be created. What is routable? Pods? Nodes?
This decisions often hold up a lot of other work and getting IP ranges can
take a long time at an organisation, so start it early.

## Requirements

* CIDR ranges decided, including an ADR
* What is routable, including an ADR
* Who is responsible for creating / managing the networks and subnets, including an ADR

## Questions

Are pods routable?

Are nodes routable?

What is the size of the node CIDR?

What is the size of the pod CIDR?




