# CECG Knowledge Platform

The CECG Knowledge Platform is how we train our platform engineers and assist our clients to train their platform engineers.

It is made up of:

* Our, client agnostic, approach to platform engineering. Specifically around building [developer platforms](./knowledge-base/content/core-platform/) and [consoldiated Path to Production](./knowledge-base/content/core-p2p/) for the tenants of the developer platform.
* [Our Bootcamp](./knowledge-base/content/bootcamp/), a hands on, intensive training that all CECG engineers go through.
* [Custom Knowledge Platform](./custom-knowledge-platform/): The ability for you to host your own version of the knowledge platform, including custom bootcamp modules specific to your setup and custom knowledge articles.

## Running your own Knowledge Platform

To get the most out of the CECG Knowledge Platform we recommend [running your own version](./custom-knowledge-platform/README.md), combining our industry knowledge with your own knowledge specific to your setup.

Let's take Developer Platform Ingress as an example. The CECG knowledge base includes:
* Features a multi-tenant developer platform should support for Ingress
* Content in the bootcamp to teach platform engineers how Ingress works 

Then you can add a custom bootcamp module, or knowledge section, to document how your specific Ingress is setup e.g.
* What Ingress Controller is used and how it is configured
* Do tenants, or groups of tenants i.e departments, get their own Ingress Controllers?
* How does Ingress work outside of the cluster? NLBs? 
* Is there a CDN or GLB in front of a multi-region setup?

That way your platform engineers can learn industry best practices and quickly follow it up with how your specific setup works.

See [Running your own Knowledge Platform](./custom-knowledge-platform/README.md).