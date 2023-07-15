+++
title = "Core Pipelines"
date = 2022-12-24T04:32:10+02:00
weight = 4
chapter = false
pre = ""
+++

We strive for a common path to production for all applications deployed to central platforms. 
It is the key enabler for **Continuous Delivery**.

Core P2P is a large scoped roadmap item of a developer platform. It is pulled into its own top level
category due to its large scope and importance.
It also can be **optionally** build on top of Multi Tenant Access or be separate e.g. on its own infrastructure or using a managed service.

### What is a common P2P?

At a high level it means all delivery teams use the same steps to deploy a service to production.

{{% notice note %}}
A common **Path to Production** (P2P) is a contract between a delivery team and operations/change management for the steps a service takes to get into production.
{{% /notice %}}

These steps are agreed across engineering teams, infrastructure teams, change management & security.
It gives a high level of confidence that every service meets the minimum level of testing and staging before going to production.

### Is this possible with teams using different technologies?

Yes! If the right level of abstraction is chosen e.g. we want every team to do stubbed functional testing, but aren't opinionated about
how they do it. We believe in a light touch approach to a standard P2P. 

### What are the benefits for a delivery team?

* Saved time on CI/CD infrastructure setup
* Pre-agreement with Change Management for automated change management process, one of the key blockers to CD
* Collective bargaining power: delivery teams working together to improve technical practices such as automated testing

### What's the right level of abstraction for the steps in a P2P?

This is a fine art! Too vague and it becomes useless, too specific and it becomes unusable by both teams.
See how to [design a P2P](./design-a-p2p) for guidance.

## Principles

### Deliver services, not artifacts
Delivery teams should deliver operable services, with all the required monitoring and observability built in from the start.

A [Core Platform](/core-platform) enables delivery teams to focus on business logic rather than infrastructure concerns but even if working in an environment without a Core Platform a team should still deliver a service.

### Decouple P2P Platform team and delivery teams
Onboarding onto a common P2P should be self-service. New teams, with whatever financial approval, should be able to have all the infrastructure required for P2P provisioned without waiting on an external team.

### Service Quality made visible
For P2P, service quality made visible is about:

* Automated testing
* Alert based promotion

Our standard [types of tests](/core-p2p/testing-strategy) are typically:

|Fast Feedback|Longer Verification|
|-------------|-------------------|
|{{< list "Stubbed Functional Tests" "Stubbed NFT Tests" "Limited Integrated Tests" >}}| {{< list "Extended NFT / Soak Tests" >}} |

Standard doesn’t mean prescribing specific tools. The P2P should have a place for these tests to run and potentially tools for displaying results but the tool choice and implementation should be up to the delivery team.

This leads to our next principle, flexible light touch approach.

### Flexible light touch approach
Resistance to a common P2P normally stems from product teams wanting autonomy in tooling and pipeline shapes.

For tooling the common P2P can describe the steps, then delegate to scripts in the product repo meaning any tools can be used to implement the steps.

Having common shapes is required for the P2P and [carrots](https://en.wikipedia.org/wiki/Carrot_and_stick) such as automated change control can be used to make product teams want to adopt the common P2P.

### Automatically generated pipelines
Fully automated pipeline creation from code (repeatability, consistency, reduced effort). A key metric is how long does it take for a new service to do the sprint 0 tasks of:

* Create all environments (easy with a self-service core platform)
* Have all pipelines for deploying to all environments

A common P2P should allow self-service generation of pipelines for new services in minutes.

### Encapsulated service repositories with a single main branch always built into an immutable, releasable artifact
Include everything that is required to run and test a service in a single repository.

Including living documentation defining at least:

* Developer onboarding
  * Building and testing locally
* Environments
  * How to deploy
  * Usage of environments. Who uses what for what.
* Tests
  * All stubbed tests (functional, NFT)

### Security shifted left
A key carrot to security stakeholders is that a common P2P will include library and image security scans.

A carrot for development teams is that they get these for “free”.

### Lightweight environments
Well-defined lightweight environments with stubbed dependencies. The “lightest weight” would allow even PRs to spin up a temporary stubbed environment. A middle ground would be a fixed set of “lockable” environments for developers and PR validation. The minimum is all the environments required for the P2P testing phases.
