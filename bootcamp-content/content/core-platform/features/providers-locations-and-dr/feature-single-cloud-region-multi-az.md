+++
title = "Single Cloud Region, Multi AZ"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation
Tolerate Cloud provider AZ outages.

## Requirements
* Node pools across many AZs
* Tests for AZ outages
* Ability to deploy node pool changes incrementally
* Node pools setup in a compatible way with the cluster autoscaler
* Cluster autoscaler deployed

## Additional Information
The minimum availability we recommend. Core Platform running across many AZs in a single Cloud Region. Node pool epic in connected kubernetes likely set this up. Use this epic to ensure proper deployment is in place and it is setup correctly with the cluster autoscaler.



