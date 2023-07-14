+++
title = "UDP Ingress"
date = 2022-01-15T11:50:10+02:00
weight = 9
chapter = false
pre = "<b></b>"
+++

## Motivation
Tenant applications require UDP ingress i.e. not over HTTP e.g. vendor products with custom protocols.

## Requirements
Raw UDP Ingress, any arbitrary UDP port / protocol.

## Additional Information
Some requirements should be gathered, if required, before Platform Ingress is implemented as it will affect the Ingress Controller and load balancer choice.

Depends on:
* Platform Ingress

Key questions for Defuzz:
* Do the nodes / pods need to be routable?
* How does load balancing work?



