+++
title = "Tenant DNS"
date = 2022-01-15T11:50:10+02:00
weight = 7
chapter = false
pre = "<b></b>"
+++

## Motivation
DNS strategy for services exposed outside of the cluster

## Requirements
* ADRs and strategy agreed for:
  * Region-specific DNS 
  * Environment specific DNS 
  * Global DNS
* Tenant ingress resources get resolvable DNS names 
* Per region DNS name: ability to target service in a single cluster
* Per provider DNS name: ability to target a service in a single provider, e.g. any AWS region
* Global DNS name: agnostic of any particular provider
* Some times there is a requirement for custom, tenant (or product) specific domain names. Can the platform handle this?  

## Additional Information
Typically involves using external-dns for the per region DNS name.

The per cluster DNS entry aka region DNS can be managed by the external-dns controller. 

Key questions for defuzz:

* Where do they need to be resolvable from?

Depends on: 
* Platform Ingress
* Cluster DNS




