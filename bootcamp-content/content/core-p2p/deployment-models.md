+++
title = "Deployment Models"
date = 2022-12-27T06:32:10+02:00
weight = 6
chapter = false
+++

Deployment models can be built into a Core P2P or a Core Platform, or a combination.

They are documented here as even if a Core P2P or a Core Platform is not in place at your client we still want to aim for **undisruptive deployments.**

To speed up development velocity we aim to:

* Reduce **Lead Time for Changes**. Reducing the time for a commit to get to production gets features to customers more quickly.
* Increase **Deployment Frequency**. Smaller, more frequent changes, reduce the likelihood of an outage rather than “big bang releases”.

### Why do we aim for undisruptive deployments?
We want to move away from large releases to production and deploy more frequently. That requires undisruptive deployments, whereby end users of the software releasing aren’t aware that a deployment has happened. Undisriptive deployments require:

1. [Application Architecture | Graceful Shutdown](/core-p2p/application-architecture/#graceful-shutdown)
1. [Application Architecture | Multiple Instances + Stateless Architecture](/core-p2p/application-architecture/#multiple-instances--stateless-architecture)

### What is deployment frequency?
Deployment frequency looks at how often a client deploys code to production or releases to end-users. Mature clients have an on-demand deployment frequency while less mature clients have a deployment frequency of only once a month or once every six months. 

## Deployment Models
### Rolling
At a bare minimum we aim for rolling deployments. Assuming multiple instances of an application, a rolling deployment replaces one instance at a time, often with checks before moving on the the next one.

### Canary
[Canary](https://martinfowler.com/bliki/CanaryRelease.html) isn’t just a deployment mechanism but a more advanced way to do a full promotion to production. First releasing to a subset of end-users before doing a full production release.

### Blue Green
We normally recommend Canary over [blue green](https://martinfowler.com/bliki/BlueGreenDeployment.html) but there have been cases where we’ve implemented blue green deployments at clients. 
