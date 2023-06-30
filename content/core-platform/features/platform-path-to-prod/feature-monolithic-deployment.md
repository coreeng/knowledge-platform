+++
title = "Monolithic Deployment"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

### Motivation
Agree with the team if a monolithic or decoupled platform deployed is best for the initial phase of the project.

### Requirements
* Versioning strategy agreed for the platform, with an ADR.
* To be able to clearly identify which version is in which environment 
* Each version should have to pass through the lower environments, environments can’t be skipped 
* The only manual action is triggering a deployment, no manual interaction with the platform

### Additional Information

Advantages:
* Easy to understand.
* Easy to deploy a whole new version as it is a single life cycle.

Downsides:
* Longer deployments
* Larger blast radius

