+++
title = "Platform Services Provisioning"
date = 2022-01-15T11:50:10+02:00
weight = 7
chapter = false
pre = "<b></b>"
+++

### Motivation

Given a decoupled platform deployment, agree on a way to deploy platform services that are not part of the base infrastructure.

### Requirements

* Mechanism for deploying platform services such as ingress controllers, policy controllers with a lifecyle that includes testing and promotion.
* ADRs for how platform services are versioned
* ADRs for how platform services follow a P2P through all standard environments
* One platform service deployed to all environments as an example for all future platform services.

### Additional Information
This can either be done in isolation, or via doing a platform services such as:
* Policy Controller
* External-dns

Care should be taken not to increase the scope of this feature to include the whole platform form service.

What will be delivered? A single platform service and a mechanism to deploy future platform services.


