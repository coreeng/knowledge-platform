+++
title = "Corporate DNS Reliability"
date = 2022-01-15T11:50:10+02:00
weight = 5
chapter = false
pre = "<b></b>"
+++

## Motivation
Instability of DNS requests that come from corporate DNS servers. Failed lookups. High latency.

## Requirements
Highly available access to corporate DNS e.g. local caching, improving reliability

## Additional Information
This has become an issue in various platforms we have experience with. Prioritised after issues have been discovered.

Included as part of Ingress for now. Maybe better placed in Connected Kubernetes.

One solution we’ve used previously is to deploy unbound. We’re moving away from that and instead using Core DNS on every node.


