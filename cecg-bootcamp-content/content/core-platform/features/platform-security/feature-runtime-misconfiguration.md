+++
title = "Runtime Misconfiguration Detection"
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b></b>"
+++

## Motivation

* Build security into the platform so that applications can adhere to security requirements by being on the platform
* Agree on Kubernetes and container best practice and have these practices enforced


## Requirements

* Agree and document the misconfiguration rules:
  * Kubernetes resource usage covering at least: Deployments, Ingress, Services
  * Docker image best practices 
* Cover Kubernetes CIS benchmark
* Cover Docker CIS benchmark

## Questions

What tool should be used to enforce the policies?

Can existing implementations of CIS benchmarks be used?

What are the policies?


