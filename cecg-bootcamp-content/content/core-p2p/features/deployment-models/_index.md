+++
title = " "
date = 2022-01-15T11:50:10+02:00
weight = 9
chapter = false
pre = "<b>Deployment Models</b>"
+++

# Deployment Models 

Deployment models can be implemented into the P2P, the Core Platform (e.g. with higher level CRDs representing applications), or left to the user.

P2P can offer varying levels of coupling with deployment:


* Fully decoupled: the P2P only offers the connectivity / authentication to the target deployment infrastructure. 
* Flexible: tooling is provided (e.g. CLI, shared build steps, recipes) that can optionally be used.
* Fully coupled: there are deployment steps that deploy to a specific target infrastructure

Another question is push vs pull deployments. 

* Push: CI/CD tooling push deployments to target infrastructure 
* Pull: CD is implemented in the target infrastructure and pulls immutable versions created by CI.