+++
title = "Basic Clusters"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

### Motivation

* The starting point of a Core Platform
* Get basic clusters running in all environments that all other features can be built on top of.
* Show progress to the business early.

### Requirements

* Kubernetes API available in all environments for other features to be built on
* A basic node pool so that platform pods and services can be deployed

### Additional Information

* Can use default node image, any CNI, just to get other team members unblocked.
* Larger decisions such as Corporate IP range for nodes can be decided in [Network Connectivity](/core-platform/features/connected-kubernetes/feature-network-connectivity/)

### Questions / Defuzz / Decisions

* Should a managed cloud offering be used? E.g. EKS/GKE?
* Does the control plane need to be private or public?
* Should nodes be public or private?
* Should nodes be configured with Internet connectivity?
* Does the installation need to be offline compatible or can images for Kubernetes system services be pulled from the Internet?
 

