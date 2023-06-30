+++
title = "CNI"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation

CNIs enabled by default for managed Kubernetes services such as EKS/GKE have the following limitations:
 * Don’t support network policies
 * Have scalability limits as they are based on iptables rather than ebpf

## Requirements

Typical requirements for a CNI for a multi tenanted:

* Is implemented at the eBPF layer for low overhead
* Supports standard network policy
* Supports cluster wide network policies for egress firewall
* Supports FQDN policies

## Questions 

Which CNI
Our recommended CNI for most use cases is [Cilium](https://cilium.io/)
Include tasks for observability e.g. if it is Cilium also deploy hubble

On GCP we recommend using DataplaneV2

