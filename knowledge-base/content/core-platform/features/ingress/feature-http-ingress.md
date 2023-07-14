+++
title = "HTTP Ingress"
date = 2022-01-15T11:50:10+02:00
weight = 8
chapter = false
pre = "<b></b>"
+++

## Motivation
Support expositing HTTP services outside of the platform

## Requirements
* The Ingress mechanism must support HTTP and/or HTTP(s)
* TBD based on defuzz

## Additional Information
Typically comes out of the box with any Platform Ingress solution.

Any additional requirements can be met under this epic e.g.
* Particular cert requirements
* Additional HTTP routing requirements
* Any additional observability specific to HTTP

Key Questions for Defuzz:
* What are the certificate requirements? Per app? Per tenant? Wildcard certs?


