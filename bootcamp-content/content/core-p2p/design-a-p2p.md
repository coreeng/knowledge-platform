+++
title = "Designing a P2P"
date = 2022-12-24T04:32:10+02:00
weight = 1
chapter = false
+++


It is common when discussing CI/CD that it ends up as a tooling discussion. We aim to avoid this. 
First agree what are the principles and requirements for a P2P and let an engineering team pick tooling to implement them.

Not every customer is ready for a common P2P. All of the principles and requirements in our common P2P 
can be used for a single P2P then used as an example for other services in the department.

## Reference P2P

Here is a common P2P that has been used at many different organisations.

### Fast Feedback

The first part of the P2P is about one thing: fast feedback to the delivery engineers.
Do as much as possible, as early as possible, and as fast as possible.

For definitions of key terms such as Quality Gates and Steps see [Promotion](../promotion).

The delivery team need to fill in the steps as defined by the common P2P. Each step can be implemented entirely differently 
by each team e.g. they could:
* Use different build tools
* Use different testing tools

{{< figure src="/images/reference/p2p/fast-feedback-pipeline.png" title="Fast Feedback Pipeline" >}}

* **Quality Gate: Local** (a.k.a Build): Run on every commit
  * **Step: Local verification**
    * Localised tests
    * Built on fresh build agents or in a container
    * Produce version, immutable artifact(s)
   
* **Quality Gate: Fast Feedback Deployed Testing:** Run tests against a deployed application
  * **Step: Stubbed Functional Testing:**
    * Application tested in isolation
    * Downstream dependencies stubbed
    * Gives confidence before integrated tests
    * Facilitates negative testing
    * Tests the full application stack: no issues with deployment manifests etc
  * **Step: Stubbed NFT:**
    * Catch any major problems early
    * Validate configuration ready for extended run
    * Much easier to gain confidence in non-integrated environments
    * Can test for graceful degradation through priming
    * Gain confidence on linear scalability
    * Used to have achieved right-sizing
  * **Step: Integrated Functional Testing:** 
    * Fast Integration off main branch
    * HA Environment Treated Like Prod
    * Non-Disruptive Deployments / Graceful Termination
    * Monitoring and Alerting Setup Early
    * Tests Own Test Data Lifecycle
    * Anyone Can Start A Test At Anytime
    * Additional Centralized Test Suites Run
    * Second Hardest Environment To Get Right

### Extended

Many businesses require long, aka "extended" testing, before a release to production.

* **Quality Gate: Extended** 
  * **Step: Extended Stubbed Functional Testing:**
    * Longer load and peak tests
    * Longer resiliency tests
    * Soak tests
  * **Step: Extended Integration Testing:**

{{< figure src="/images/reference/p2p/extended-tests-pipeline.png" title="Extended Tests Pipeline" >}}

### Deployment models 

See [Deployment Models](../deployment-models).

Not all promotion is directly related to testing. Time spent in a previous environment without any alerts for example
can lead to a higher level of confidence before a full production release.
A common version of this is canarying where a release is tried out on a subset of users or a subset of traffic.
In multi AZ/Region setups a release can be run for some time in a subset of regions before a full rollout.

{{< figure src="/images/reference/p2p/canary-and-production-deployemnt-pipeline.png" title="Canary and Production Deployment Pipeline" >}}

### Common Discussion points
* Which parts of the pipeline are run on PR validation? 
  * If lightweight environments exist ideally all stubbed testing can run
  * Some users might just want to run local testing
* What is the deployable?
  * We want it to be versioned and immutable, it is typically a container image
  * Does it include the tests versioned with it?
* What is the promotion mechanism? See [promotion](../promotion)

## Build a Common P2P

Assuming no existing Common P2P you need to design it with your initial users.
The biggest pitfall is to design it in isolation and try and force it on users. Instead
work with a set of users, find out the biggest problems, and aim to solve them with a MVP pipeline.


{{<mermaid align="left">}}
graph LR;
  A(Definition) --> B(Build It: MVP with first tenants)
  B --> C(Launch & Adopt)
  C --> D(Evolve)
{{< /mermaid>}}
