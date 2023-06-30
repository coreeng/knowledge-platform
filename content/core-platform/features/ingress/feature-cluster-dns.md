+++
title = "Cluster DNS"
date = 2022-01-15T11:50:10+02:00
weight = 6
chapter = false
pre = "<b></b>"
+++

## Motivation
Clearly documented DNS strategy for platform.

## Requirements
* Highly available internal DNS to the cluster

## Additional Information
At some client we run dnsmasq to cache cluster DNS

Typically implemented with Core DNS. EKS/GKE come with default deployments of this.

A potential future solution is to run Core DNS on every node and remove dnsmasq. 


