+++
title = "Node Pools"
date = 2022-01-15T11:50:10+02:00
weight = 3
chapter = false
pre = "<b></b>"
+++

## Motivation

Provide VMs for single region, single provider.
Gather base requirements around disaster recovery and availability

### Requirements

* Node pools covering multiple AZs
* Strategy for node patches
* Decision on whether a custom AMI needs to be used
* Future proofed for the Cluster Autoscaler
* Ability to deploy a new node image in a staggered way e.g. 
  * canary a node image
  * per AZ rolling deployments

## Questions / Defuzz / Decisions

Should there be separate node pools per AZ to ease integration with the Cluster Autoscaler?

* This is because if you have a single node pool and then a request for a particular AZ (e.g. it depends on zonal EBS) then the autoscaler will not create a node in the right AZ if there is capacity. Whereas if you have a EKS node pool per AZ, it will work

What type of staggered node releases can be done?

* Canary node image
* Per AZ rolling deployments


## Additional Information

Depends on [Basic Clusters](/core-platform/features/connected-kubernetes/feature-basic-cluster/)

