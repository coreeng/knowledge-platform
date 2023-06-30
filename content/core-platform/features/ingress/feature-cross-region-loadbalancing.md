+++
title = "Cross Region Loadbalancing"
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b></b>"
+++

### Motivation

* Expose services active, active, across many regions
* Failover services between regions without the end user being affected as the load balancer they connect to is region agnostic.
* Provider agnostic loadbalancing.

### Requirements

* Loadbalancing across regions for the same application

### Additional Information
If the company’s LB/edge infrastructure does not support cross region load balancing then we need to provide that.

If an initial phase is a single region but multiple regions are planned then ingress into the cluster should be done in a way that won’t change when multiple regions are added.

