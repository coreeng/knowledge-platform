+++
title = "Multiple Platform Lifecycles"
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b></b>"
+++

### Motivation

Agree with the team if a monolithic or decoupled platform deployed is best for the initial phase of the project.

### Requirements

* Platform deployed as multiple components, these should be defined e.g.
  * Base Infra
  * Each platform component has its own lifecycle.
* Clearly defined interfaces between the components 
* Ability to deploy the components independently i.e. avoid a distributed monolith

### Additional Information

Deploy the platform as many loosely coupled services. 
This allows quick changes, with smaller blast radius's for parts of the platform.
This is a decision that needs to be made (and can be revisited) rather than a feature.
The outcome should be team agreement and ADRs.


