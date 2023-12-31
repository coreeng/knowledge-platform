+++
title = "Consolidated P2P"
date = 2022-12-24T04:32:10+02:00
weight = 1
chapter = false
+++

A consolidated P2P at a high level it means all delivery teams use the same steps to deploy a service to production.

{{% notice note %}}
A common **Path to Production** (P2P) is a contract between a delivery team and operations/change management for the steps a service takes to get into production.
{{% /notice %}}

These steps are agreed across engineering teams, infrastructure teams, change management & security.
It gives a high level of confidence that every service meets the minimum level of testing and staging before going to production.

## Maturity

The maturity of a consolidated P2P, without up front buy in, typically follows:

1. [Agreed / documented](#agreed--documented): A team, department, or whole organisation has agreed and documented the steps of a P2P.
2. [Templated](#templated): The agreed P2P is captured as a template that can be used by delivery teams.
3. [Automatically generated by a Platform Engineering function](#automatically-generated-as-a-platform-engineering-service): The agreed P2P is automatically generated by tenant configuration.

### Agreed / Documented

Initially this could be within a single team that manages many services.
The first step is at this level agreeing that a common P2P should be used.
It could just be documented that this is what should be implemented.

**Pros**

* Starts the conversation about the need for a common P2P
* Sets expectations for delivery teams on what they should be doing in their P2P

**Cons**

* Hard to measure adoption
* Lots of work for every team to adopt
* Hard to measure conformance
* As the documented pipeline changes, the teams that have followed it will go out of sync. Losing many of the benefits of a common P2P.

### Templated

Pipeline generation can be templated with tools such as [cookie cutter](https://cookiecutter.readthedocs.io/en/stable/).

**Pros**

* Very low investment to get going
* Speeds up delivery teams getting a quality P2P right away

**Cons**

* As teams generate pipelines with different versions of the template, they slowly go out of sync. Losing many of the benefits of a common P2P.
* No clear owner of the templates

### Automatically generated as a platform engineering service

Based on configuration, that ideally lives in the delivery team's repository, a platform service generates the agreed P2P.

**Pros**

* Pipelines are kept up to date as the P2P evolves as the platform service updates pipelines
  * Much better story for Security and Change Management to adopt Continuous Delivery if there is assurances every team are using the agreed, and up to date, P2P.
* Very low effort for the delivery team
 

**Cons**

* Significant buy in needed from engineering