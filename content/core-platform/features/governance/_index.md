+++
title = " "
date = 2022-01-15T11:50:10+02:00
weight = 11
chapter = false
pre = "<b>Governance</b>"
+++

# Governance
Enforce platform, security, and other corporate governance function’s policies within the platform.

## Details
A big [carrot](https://en.wikipedia.org/wiki/Carrot_and_stick) for security and governance stakeholders is that many policies typically implemented by manual validation can be “guaranteed” for applications deployed to a core platform. Some examples:

* Ingress/TLS works the approved way for every application
* Access control is standardised
* Labelling / application ownership is standardised e.g. all applications are labelled and CIs created in service now

In additional as a company is adopting containers we can have a core platform enforce Container/K8s best practices.

### Policy Manger
Governance and security can be partly implemented with a Kubernetes Policy Manager. 

We use the following in production:

* [OPA/Gatekeeper](https://github.com/open-policy-agent/gatekeeper)
* [JSPolicy](https://www.jspolicy.com/)

Example policies we use in production:

* Corporate governance
  * Validating/setting required labels
  * Denying annotations e.g. if platform uses aws-lb controller that tenants aren’t meant to use
* Kubernetes best practice
  * Don’t use default namespace
  * Setting memory limits and requests
  * Setting cpu requests
  * Probes
* Docker best practice
  * No use of the latest tag
  * Read only root file system
* Tenant guard rails
  * Setting and or validating host/target DNS annotations
  * Setting/validating ingress class
  * Unique host names
* Privileged
  * Deny host mounts
  * Deny host networking
  * Deny privilege escalation
  
## Extensions

* A standard set of policies we recommend for JSPolicy/OPAGatekeeper
* A set of policies we use to do a production readiness check for new clients
