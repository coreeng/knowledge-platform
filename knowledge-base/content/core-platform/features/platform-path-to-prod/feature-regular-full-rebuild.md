+++
title = "Regular Full Rebuild"
date = 2022-01-15T11:50:10+02:00
weight = 13
chapter = false
pre = "<b></b>"
+++

### Motivation

High confidence is required that a full rebuild of the core platform will always work, typically motivated by:

* Regulations e.g. TSR
* Having to provision new environments frequently e.g. an operating model that has many core platforms per env rather than a single multi-tenanted environment. 

### Requirements

The ability to fully rebuild the platform, hands off, given new base infrastructure.

### Additional Information

This is something we always aim for but can involve significant engineering work so the ROI (Return Of Investment) needs to be evaluated.

In our experience every IaC if run from scratch will not work as they have been incrementally built.


