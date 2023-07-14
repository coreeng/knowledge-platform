+++
title = "TCP Ingress"
date = 2022-01-15T11:50:10+02:00
weight = 8
chapter = false
pre = "<b></b>"
+++

## Motivation
Tenant applications require TCP ingress i.e. not over HTTP e.g. vendor products with custom protocols.

## Requirements
Raw TCP Ingress, any arbitrary TCP port / protocol.

## Additional Information
Should be known before Platform Ingress is implemented as it will affect the Ingress Controller and load balancer choice.

Depends on:

* [Platform Ingress](./feature-platform-ingress)

Key questions for Defuzz:

* Do the nodes / pods need to be routable?
* Can the TCP ingress be load balanced?



