+++
title = "Tactical Registries"
date = 2022-01-15T11:50:10+02:00
weight = 4
chapter = false
pre = "<b></b>"
+++

### Motivation

* Remove burden on delivery teams to provision registries.
* Allow early onboarding of delivery teams

### Requirements

* Registries required for tenants before strategic multi provider solution 
* Delivery team authentication and authorisation to registries 
  * Including any promotion required between environments
* Platform team authentication and authorisation to registries
  * Including any promotion required between environments

### Additional Information

* May be skipped due to strategic registry solution 
* Key enabler for many security features if all images come from a single registry per environment
* Key enabler for a promotion mechanism between environments for Continuous Delivery.
* Can be built on top of cloud provided services such as ECR
* Depends on:
  * Base Infrastructure (if they are provisioned via the Platform P2P, not if they are independent)


