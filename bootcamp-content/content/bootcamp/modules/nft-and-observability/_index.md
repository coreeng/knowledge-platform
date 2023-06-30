+++
title = "NFT and Observability"
date = 2022-12-22T05:07:10+02:00
weight = 2
chapter = false 
+++

## Introduction 

**Theme:** NFT & Observability

**What should you {{< colour green understand >}} at the end of the module?**

* The later stages of a mature Path to Production 
  * [Promotion: Gates, steps](/core-p2p/promotion)
  * [Extended testing](/core-p2p/testing-strategy/)
  * Promotion based on time in environments, alerts, monitoring etc
* Monitoring and alerting
  * Metrics, aggregated logs, tracing
  * The four golden metrics
  * Three pillars of observability
* Grafana, prometheus, and alert manager
* Monitoring as code
* Right sizing for migration to Cloud 
  * Vertical scaling + resource limits
  * Horizontal scaling
* NFT test types
  * Smoke 
  * Soak
  * Peak load
  * Resiliency 


**What should you be {{< colour green "able to do" >}} by the end of the week?**

* Install prometheus, grafana, and alert manager to minikube
* Add metrics to your applications and use them to:
  * Analyse performance issues
* Defining dashboards, alerts in their service repo
  * Unit test alerts using Promtool 
* Write NFT tests
* Prove linear via horizontal scalability

### FAQs
**Why do the CPU / Memory dashboards update so slowly?**

By default prometheus is configured to scrape the resource statistics from [cadvisor](https://github.com/google/cadvisor) every 60s. This can be reduced but increases the resources required by prometheus. 

**What is cadvisor?**
[cAdvisor](https://github.com/google/cadvisor) (Container Advisor) provides container users an understanding of the resource usage and performance characteristics of their running containers. It is run as a daemon that collects, aggregates, processes, and exports information about running containers.

**Where is cadvisor installed?**
Cadvisor is embedded in [kubelet](https://kubernetes.io/docs/command-line-tools-reference/kubelet/#:~:text=The%20kubelet%20is%20the%20primary,object%20that%20describes%20a%20pod.) which is the agent that runs on every node of a Kubernetes cluster. It doesn’t need to be separately deployed.

