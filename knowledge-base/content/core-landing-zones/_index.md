+++
title = "Core Landing Zones"
weight = 101
chapter = false
pre = ""
+++

A cloud landing zone is a cloud adoption framework, by using that organizations can perform large-scale cloud migration in an efficient and streamlined manner 

The concept is the same across Cloud Provider but implemented differently.

* [AWS]https://docs.aws.amazon.com/prescriptive-guidance/latest/migration-aws-environment/understanding-landing-zones.html) using a multi account framework
* [GCP](https://cloud.google.com/architecture/landing-zones) using a resource hierarchy and organisation policy
* [Azure](https://learn.microsoft.com/en-us/azure/cloud-adoption-framework/ready/landing-zone/) using Management Groups and subscriptions


## Principles

To design and build scalable landing zones we adopt the following principles:

* **Self service**: Users with the appropriate approvals can onboard via a self service API 
* **Tenant isolation**: one tenant should not be able to adversely affect another  
* **Environment isolation**: one environment should not be able to affect another 
* **Simple for users**: Make the right things easy to do 
* **Everything as code**: Full infrastructure lifecycle management 
* **Dev is prod**: Treat workload development environments as production 

## Scope and features 

A landing zone can vary hugely in scope depending on the use case. Ranging through:

* Small scale isolated Cloud usage in a single region
* Large scale isolated Cloud usage in a single region
* Large scale isolated Cloud usage in multiple regions
* Cross Cloud usage
* Hybrid Cloud usage

Where scale is the number of teams, engineers and the size of the workloads.

Another axis of scope is the variety of the Cloud usage. A small number of application
types will require limited Cloud expertise to set the appropriate security controls and policies.
The more types of workloads that run in the Cloud the more expertise is required at the Landing Zone
level to set appropriate controls.

We categorise the features of a Cloud Landing zone into:

* [Access Management and Identity](./identity/)
* [Disaster Recovery, Regions & Hybrid](./regions/)
* [Security, Governance & Audit](./security-controls/)
* [Billing & Financial Reporting](./finops/)
* [Golden Images](./images/)
* [Networking](./networking/)
* [Landing Zone P2P](./p2p/)
* [User Interface](./developer-experience/)

## Requirements

As most enterprises typically end up in multiple clouds specify requirements in a Cloud Agnostic way
then go into how to implement those requirements in the major cloud providers.





