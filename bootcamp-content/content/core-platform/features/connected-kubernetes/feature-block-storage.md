+++
title = "Block Storage"
date = 2022-01-15T11:50:10+02:00
weight = 6
chapter = false
pre = "<b></b>"
+++

## Motivation
Deploy services such as prometheus that require persistent storage.

## Requirements
* Persistent block storage that persists beyond the pod lifecycle e.g EBS, GCP Zonal Storage
* Ability to deploy single AZ deployments if block storage is tied to an AZ
* Ability to deploy to full region if block storage is region

## Additional Information
* May be offered as a tenant feature or just for use by the platform team initially. Break into separate sub-epics if they will be delivered separately. 
* Depends on:
  * Node Pools (as once across multiple AZs, implementation may change).

