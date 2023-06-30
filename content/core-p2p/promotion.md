+++
title = "Promotion"
date = 2022-12-27T06:32:10+02:00
weight = 4
chapter = false
+++

## Quality Gates
A quality gate classifies a type of verification done against a service to assure its quality. Example quality gates are:

* **Local:** Every type of test/scan that can be done locally
* **Fast Feedback Deployed Tests:** All types of quick deployed functional and non-functional 
* **Extended Deployed Tests:** Longer running extended tests such as soak and long peak load tests
* **Canary:** Validate with a subset of production traffic, tested via stability, monitoring, alerts rather than one off tests

## Steps
A step is a specific check required to pass a quality gate. For example:

To pass the **Local** quality gate the steps could be:

* **Step:** Unit tests
* **Step:** Image vulnerability scans
* **Step:** Non-deployed functional tests

## Environment Classes
 

The P2P defines standard environments and promotion is about when a service can enter those environments. Typical environment classes:

* **Dev**
  * Only used by the engineering team responsible for the service e.g. stubbed-functional-test, stubbed-nft, stubbed-showcase
* **Integrated**
  * Integrated with other engineering teams, must be kept reliable, typically quality gates Local and Fast Feedback Deployed Tests and possibly even Extended Deployed Tests
* **Production-Canary**
  * A subset of production traffic
* **Production**

## Promotion
After each quality gate, or a set of quality gates, a service can be promoted to the the next class of environment.

## Promotion Implementation
### Immutable artifact copy
TODO

### GitOps
TODO

### Binary Authorization
TODO
