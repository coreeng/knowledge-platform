+++
title = "Core Platform in the Cloud"
weight = 5
chapter = false
+++

##  Motivation

* Understand the main components of a core platform in a cloud environment

## Requirements

Implement the following Core Platform features.

* [Connected Kubernetes: Basic Clusters](/core-platform/features/connected-kubernetes/feature-basic-cluster/)
* [Connected Kubernetes: Node Pools](/core-platform/features/connected-kubernetes/feature-node-pools/)
* [Kubernetes Network: CNI](/core-platform/features/kubernetes-network/feature-cni/)
* [Ingress: Platform Ingress](/core-platform/features/ingress/feature-platform-ingress/)
* An end-to-end test (preferably Golang) of Ingress that:
  * Deploys an application
  * Exposes the service via ingress
  * Makes a successful request to the application over HTTP(s)
  * Run after every deployment

All features should be deployed automatically with no manual actions in GCP console or the Kubernetes API.

Additional or more specific requirements than are what on the Core Platform knowledge base:

### Basic Clusters

Control Plane CIDR: 10.128.0.0/28
Node CIDR: 10.0.0.0/24
Min size: 1
Max size: 3
In the VPC created in IaC Setup

### Node Pools

Spread across multiple AZs
Standard GKE or EKS node image

###  CNI

GCP: DataplaneV2
AWS: Cilium

### Platform Ingress

Use the Contour Ingress controller

## Additional Information

Build on top of IaC Setup.