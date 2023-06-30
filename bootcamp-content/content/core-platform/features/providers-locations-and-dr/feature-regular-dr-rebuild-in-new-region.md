+++
title = "Regular DR rebuild in new region"
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b></b>"
+++

## Motivation
Alternative to multi region

Regulatory requirements 

## Requirements
* Regular rebuilding of the whole platform in a different region
* External and internal Ingress to the platform can be switched to the new region with minimal end user impace

## Additional Information
We prefer to do multi region setups so the other region is always available.

Tenants can then do active passive or active active.

In that scenario we aim to reduce blast radius by updating a single region at a time.



