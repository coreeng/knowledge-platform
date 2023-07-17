+++
title = "Platform SLI & SLO definitions"
date = 2022-06-09T17:18:00+01:00
weight = 5
chapter = false
pre = "<b></b>"
+++

Here is an example of SLI/SLO definition that form a mini platform engagement. It appears to be sufficient for that particular engagement but YMMV. That being said you are solid if you:

* Have a clear boundary of what the platform manages, and what the tenants manage (what AWS calls the "shared responsibility model"). What the platform manages should be covered by SLOs & SLIs.
* Use the user journey as the north star, and define the SLOs & SLIs based on how the tenants interact with the platform.

## Context and Problem Statement

We want to define SLOs & SLIs for the platform and derive alerts based on it. By doing so:

* Across different teams we can have a consistent, measurable expectation of what the platform delivers from SLI & SLO definitions.
* When the SLOs is at risk or not being met, alerts can be triggered to notify the platform team to take action.

## Decision Drivers

* The SLOs & SLIs must be agreed with the business
* The SLOs & SLIs must be measurable
* The alert definition must be based on SLOs & SLIs
* The alert must be actionable

## Principles

* Having the right amount of SLIs & SLOs - too many causes attention fatigues, whereas too few causes oversight.
* Exercising in a structured manner - drive the expectation using SLOs, measuring using SLIs, and implementing using metrics & alerts.
* Perfect is the enemy of good - implement the SLOs & SLIs good enough and improve over time.
* It’s not a close-door exercise - make sure it’s agreed with the business.
* Alert must be actionable - if the alert is not actionable, it’s not an alert, it’s a notification.

## Out of Scope

* Tenants application monitoring & alert - although it’s crucial to have happy customers (and sometimes tenant app failure is merely a syndrome of platform outage), it’s unreasonable and unfair to measure platform healthiness based on tenant app’s healthiness. That being said a guideline for measuring tenant app’s SLIs & SLOs are needed but it will be treated as a separate matter.
* We cannot guarantee the data freshness of tenant's monitoring metrics, logs and traces appears in the Grafana Cloud, but rather in a best-effort QoS, since it is out of our control.

## SLOs & SLIs definitions based on the tenant user journey

Reliability can be characterised via availability, throughput, correctness, durability, latency, coverage, quality, freshness. In this section we will use the user journey as the north star, and define the SLOs & SLIs based on the user journey using the categories above.

Note that all the SLOs and SLIs are measured in non-maintenance periods.

### Control plane

* As the tenant of the platform I want to deploy and onboard app to the platform via the platform control plane. Large majority of the API requests to the control plane are successful, assuming the request is valid. (from availability perspective)
* As the tenant of the platform I want computing resources requested by my app to be scheduled and available in a timely manner assuming 1) the app is functional 2) the resource requested is reasonable (from latency perspective)

| Category | SLI      | SLO | Reference |
| ----------- | ----------- | ----------- | ----------- |
| Availability | The proportion of HTTP requests processed successfully by the platform k8s API server, measured from the k8s API server scrapping endpoint. _Any HTTP status other than 5xx is considered successful._ | x% over a 30 day rolling window | https://cloud.google.com/kubernetes-engine/sla |
| Availability | Consecutive minutes where the API server is not accessible by the monitoring system. _The API server is considered to be accessible when the API server endpoint can be hit from a health check probe running inside the cluster_ | consecutive minutes <= X        | https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubeapidown/ |
| Latency | Startup latency of schedulable stateless pods, excluding time to pull images and run init containers, measured from pod creation timestamp to when all its containers are reported as started and observed via watch, measured as 99th percentile over last 5 minutes<sup>[1]</sup> | 99th percentile per cluster-day <= 5s | https://github.com/kubernetes/community/blob/master/sig-scalability/slos/pod_startup_latency.md |

**Notes**:
- [1] This might be tricky to measure. Another option is to measure "startup latency of schedulable pods, measured from pod creation timestamp to when the pod is marked as `scheduled`, measured as 99th percentile over the last 5 mintes.

### Data plane

* As the tenant of the platform I want to make sure that there's enough capacity on the platform to schedule my workload

| Category | SLI                                               | SLO                              | Reference |
| ----------- |---------------------------------------------------|----------------------------------| ----------- |
| Availability | The proportion of cpu available on the cluster    | 1 node worth of cpu available    |  |
| Availability | The proportion of memory available on the cluster | 1 node worth of memory available |  |

**Notes**:
* Tenants are expected to request enough resources to sustain a single node outage as part of their deployment
* Keeping 1 node's worth of resource as a buffer is only required for (quick) autoscaling, but we keep it in our case to provide a minimum zonal redundancy given the small number of tenants we have (and so nodes)

### Data plane networking

* As the tenant of the platform I want to make sure that my application can reach out to the services of on-prem DC and over the internet to 3rd parties.

| Category | SLI                                                                                                                                                                                                                                                                    | SLO | Reference |
| ----------- |------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------| ----------- | ----------- |
| Availability | The proportion of successful http requests to on-prem web service made by the blackbox health check probes, measured from the metrics endpoint provided on the probes[1]</sup>. _Any responses come back from the on-prem web service is considered to be successful._ | x% over a 30 day rolling window | |
| Availability | The proportion of successful http requests to the internet service made by the blackbox health check probes, measured from the metrics endpoint provided on the probes. _Any responses come back from the internet service is considered to be successful._            | x% over a 30 day rolling window | |
| Availability | The proportion of successful http requests to <INSERT_AN_ON_PREM_COPONENT_HERE> by the blackbox health check probes, measured from the metrics endpoint provided on the probes. _Any responses come back from Vault is considered to be successful._                                     | x% over a 30 day rolling window | |

**Notes**:
- [1] Ideally running a http(s) request to a FQDN thus we check both TCP & UDP end 2 end. The unideal side is any external outages that are out of our control might affect the QoS we measure internally as a platform. From implementation perspective this can be achieved by the [prometheus blackbox exporter](https://github.com/prometheus/blackbox_exporter).

### Load balancing

* As the tenant of the platform I want my application to be reliably accessible via the internet through `platform.$CORP.com` domain name and on-prem DC securely via HTTPS through a load balancer. (from availability perspective)
* As the tenant of the platform I want the load balancer used by my application with acceptable latency to ensure responsiveness of my application. (from latency perspective)
* As the tenant of the platform if I use the managed certificate from the load balancer, I want to make sure the certificate is valid and not expired. (from availability perspective)

| Category | SLI                                                                                                                                                                                                                                                                              | SLO | Reference |
| ----------- |----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------| ----------- | ----------- |
| Availability | The proportion of successful http requests to the internal / external load balancer's `/ping` path<sup>[1]</sup> made by the black box health check probes, measured from the metrics endpoint provided on the probes. _Any HTTP status other than 5xx is considered successful._ | x% over a 30 day rolling window |  |
| Latency | The proportion of sufficiently fast requests hitting `/ping` path measured from the internal / external load balancer. "Sufficiently fast" is defined as <=100ms                                                                                                                 | 95 of the requests <= 100ms overa 30 day rolling window. |  |
| Availability | Expiry date of the TLS certificate for both internal / external load balancer. The expiry date is defined by the `Validity/Not After` section of the platform-managed TLS certificate on the ingress                                                                             | X day before certificate expire |  |

**Notes**:
- [1] The ping path must have a backend service deployed to ensure the http requests are tested e2e.

### Observability platform

* As the tenant of the platform I want to be able to have observability of the application from the Argus platform. (from availability perspective)

| Category | SLI                                                                                                                              | SLO | Reference |
| ----------- |----------------------------------------------------------------------------------------------------------------------------------| ----------- | ----------- |
| Availability | The proportion of times when the grafana agent is available, measured by the up query on the Grafana Cloud prometheus database . | x% over a 30 day rolling window |  |


## Decision outcome

As it stands we primarily use SLO-based alerts. This helps us to narrow the scope of our alerts to the symptons that genuinely impact the service reliability that are experienced by our customers (tenants & end users in our cases). 
By decoupling symptoms from what and why, traditional system monitoring alerts (cpu, memory, disk USE on a node-by-node basis) are largely eliminated. 
There are a few benefits of this approach:

* It helps operators to pin-point where the investigation _should_ start (where the customer experiences are actually impacted) with minimum noises.
* It makes the alerts more actionable vs traditional system monitoring where CPU, OOM errors tell you nothing about whether end users are impacted, neither it is easy to act upon.
* It largely reduced the alerting fatigue, as the number of alerts are reduced significantly.
* It generally surfaces the issues earlier than traditional system monitoring alerts. For example your app might appear to be looking fine from the traditional system monitoring perspective, but alerts show excessive SLO error burn rate (e.g. 10% error budget burned over 5 minutes, meaning 50min before SLO breaches). After alerts you drill down to the issue and pin-point the root cause to an OOM of the app. For this scenario traditional system alerts would have surfaced the issue much later, when the OOM flooded the entire fleet.

That being said system/low-level monitoring metrics are still useful for observability and issue drill down in general. They need to be collected for observability/warning purpose, but not necessarily alerting.