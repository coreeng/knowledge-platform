+++
title = "Runtime Vulnerability Detection"
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b></b>"
+++

## Motivation

* Build security into the platform so that applications can adhere to security requirements by being deployed to the platform
* Detect vulnerabilities that were discovered after an application is deployed
* For instances where build time vulnerability detection isn't in place

## Requirements

* Every image in-use is regularly scanned
* Report is easily extractable for current vulnerabilities for in-use images 
* New vulnerabilities can be alerted on and that alert includes metadata: who owns the image
* Policy for how long vulnerable images can remain in-use
* Exception process for policy conformance 

## Questions

Can the solution be Cloud Agnostic?

How often should images be scanned? Is it configurable?

What is the alert mechanism?

For alerting, if an image is used by multiple teams, is that 1 alert or many?

What tool will be used?

What central reporting needs to take place (for enterprise)?

## Additional Information

Normally depends on [Policy Controller](../governance/feature-policy-controller)
Depends on every use of an image having the right metadata which can be enforced/implemented by multi-tenant access.

Scanning: To a vulnerability scan
Checking: Validating a scan has taken place and it conforms to a policy

