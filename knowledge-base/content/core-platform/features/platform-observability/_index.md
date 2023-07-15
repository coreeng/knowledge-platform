+++
title = " "
date = 2022-01-15T11:50:10+02:00
weight = 10
chapter = false
pre = "<b>Platform Observability</b>"
+++

# Platform Observability
Ensure the quality of our running platforms, visibility for debugging and insights into platform usage.

## Details

### Goals of Observability

#### Ensure Service Level Objectives (SLOs)
* Ingress traffic latency
* Pod scheduling latency
* Container churn rates

#### Alert on Error States
* Internal/External DNS
* Unresponsive nodes
* Internal/external network connectivity

#### Debugging Problems
* Audit logs
* Event logs
* Container logs
* Container metrics
* Application metrics
* Cluster metrics
* Network flow logs (L4)
* Application traffic tracing (L7)
* Aggregation Dashboards
* Debug containers

#### Historical Analysis
* Long term queries
* Reports of resource usage

### Observability Tooling
* Prometheus stack in every cluster
  * cadvisor (container stats)
  * node-exporter (node stats)
  * metricbeat (kubernetes events) 
* Log shipping / ELK stack in every cluster
* Tenant facing Prometheus for federation
* Thanos / cortex / etc. for HA and long term storage

### Draft - Use cases

#### Use Case 1:
* Tools:
  * Prometheus
  * Grafana(Monitoring + Alerts)
  * Alert Manager
  * Service Now
* Tools runs in each clusters
  * Federates from platform prometheus to tenant prometheus
* Has on-call rotations
* Using Thanos for long-time persistency
* Using Kibana and ES for logging - logs available in Grafana 
* There are a lot of shared resources like ES indexes, so tenants may impact eachother with what they do
 
##### Alerts
* Focused on the platform -started out by having the tenant resources usages, but caused too mcuh noise
* Total node memory and CPU available
* Unschedulable pods 

##### Debug:
* Logs and metrics
* Fallback to AWS console
