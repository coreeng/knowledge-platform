+++
title = " "
date = 2022-01-15T11:50:10+02:00
weight = 6
chapter = false
pre = "<b>Ingress</b>"
+++

# Ingress

Ingress allows services in the platform to be exposed outside of the platform over a variety of protocols e.g. HTTP, HTTP(s), TCP, UDP. 

Depending on requirements this may be all the way to the Internet or just internally within the organisation's network.

A suitable ingress (into the cluster) solution to meet the requirements of the initial tenants + a year’s worth of growth.

### Platform Ingress

Our preferred solution is to use a limited number of external load balancers e.g. AWS NLBs/ALBs and to use an internal ingress / reverse-proxy controller where you can have 1 routable IP address mapped to multiple ingresses.

AWS and GCP both provide an Ingress Controllers that map Ingress/Service Loadbalancer types to external load balancers. This approach works if there are a small number of applications to be exposed externally but not practical for multi tenanted platform with 100s of applications. The downsides are:
* **Cost Overhead**: External load balancers, whether cloud-managed or instantiated through physical network appliances on-premises, can be expensive. For Kubernetes clusters with many externally exposed applications, these costs will quickly add up. 
* **Operational Complexity**: Load balancers require various resources (IP addresses, DNS, certificates, etc.) that can be difficult to manage without accompanying automation. Many large enterprises aren’t at the point where they are willing to use technology like LetsEncrypt. 
* **Features**: In cluster Ingress Controllers typically have significantly more features (e.g. app level metrics, enriched logs, path re-writing) than Ingress Controllers that configure external load balancers 
* **Lock-in**: Exposing ingress features from an ingress controller that can be deployed across cloud provider / on prem clusters helps prevent lock in. 
* **Scalability**: Managed layer 7 load balancers tend to have issues scaling quickly in response to demand. This is very evident with the AWS ELBs/ALBs for example, which are a set of managed VMs that scale in response to demand. These can take 5-10 mins to scale up to larger loads, which is problematic when serving spikey traffic.


