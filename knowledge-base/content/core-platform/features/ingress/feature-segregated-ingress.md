+++
title = "Segregated Ingress"
date = 2022-01-15T11:50:10+02:00
weight = 4
chapter = false
pre = "<b></b>"
+++

## Motivation
There are many reasons why a single set of Ingress Controllers i.e. one Deployment of an Ingress Controller for all tenants won't meet all requirements. The most common are:
* Scalability
* Ability to deploy separately to reduce blast radius
* Different settings per tenant for the ingress controllers
* Different ingress controllers for internal vs external traffic
* Dealing with ad-hoc vendor products that mandate another ingress controller
  * Normally we can work around this and use a different ingress controller

## Requirements
* Support multiple deployments of the ingress controller with different ingress classes

## Additional Information
Key questions for defuzz:
* Do we want separate classes for internal vs external?
* Do we want separate classes for different tenants?


