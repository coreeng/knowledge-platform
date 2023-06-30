+++
title = "Service Mesh"
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b></b>"
+++

## Motivation

Out of the box observability, including:

* Service interaction 
* Latencies between services
* Success/failure rates between services

Out of the box encryption e.g. mTLS (can also be satisfied at the [CNI layer](../kubernetes-network/feature-cni))

## Requirements

* Justification, including what functionality is needed above that of the CNI, for Service Mesh captured as an ADR
* Out of the box observability for:
  * Service interaction
  * Inter-service latency
  * Inter-service success/errors
* ADR for chosen strategy for mapping Gateways to tenants

## Questions 

How does the service mesh interact with [Platform Ingress](../ingress/feature-platform-ingress)?
The two should be designed together.

* Are gateways used instead of Ingress Controllers?
* How are gateways exposed externally (i.e. to the Internet)?
* How are gateways exposed internally (i.e. private in the corporation)?
* Are different gateways used for internal vs external traffic?

Is a service mesh actually required? Or can the required functionality be provided by the CNI?

Which service mesh?

* Industry leader is considered Istio
* Cilium service mesh - considered the eBPF leader
* Linkerd - for ease of use

Can the service mesh co-exist with applications that don't use it? If so, how?

## Additional Information

* If a service mesh is to be used to provider high level features observability and transparent encryption, deploy it early to avoid disruption

### Istio

#### Additional Requirements for Istio

* Highly available prometheus for metrics
  * What should the retention be?
* Prometheus automatically scrape control plane metrics
* Prometheus automatically scrape data plane metrics
* Grafana configured with dashboards if already exists, or a new deployment of grafana just for Istio with the Istio dashboards installed

#### Additional Questions for Istio

What will the [deployment model](https://istio.io/latest/docs/ops/deployment/deployment-models/) be?

* How many clusters are there? How many could there be potentially when additional regions and cloud providers are added?
* Should a mesh be isolated to a cluster or span clusters in the same environment?
* How is the mesh control plane made HA?
* How are Istio upgrades managed?

Should Istio gateways be used for egress?

What UI will be used?

* [Kiali](https://istio.io/latest/docs/ops/integrations/kiali/)?

How will the UI be authenticated and authorised?

How will Istio and addons be installed?

Should the Istio or (soon to be standard) Gateway APIs be used?

For the metrics backend:

* Should Istio have its own instance of prometheus / grafana?
* What are the deployment's limits? How many tenant services can it handle?
* How can the metrics backend be scaled?

#### Istio on GKE

* Native support has been deprecated with Anthos being the replacement
* To use OS Istio with a private GKE cluster the master firewall rule needs [updated for the webhook](https://istio.io/latest/docs/setup/platform-setup/gke/)

## Deliverables

Potential deliverables based on a typical deployment

* [ ] ADR for deployment model
* [ ] ADR for metrics backend deployment, scalability and reliability 
* [ ] Istio P2P 
* [ ] Kiali P2P



