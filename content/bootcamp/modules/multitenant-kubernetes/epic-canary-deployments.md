+++
title = "Canary Deployments"
weight = 4
chapter = false
+++

## Motivation

* Learn about a key deployment model: Canary
* Apply knowledge about Kubernetes Ingress, Services, Deployments & Labels
* Understand how to deploy a custom alert manager / prometheus.

## Requirements

* The canary deployment is captured in code with no manual actions
* Application has at least 3 replicas at all times
* To deploy a new version of an application:
    * Deploy 1 new replica with the new version
    * Run smoke tests, just targeting that 1 replica, before receiving traffic through Ingress
    * Enable that replica to receive traffic through Ingress
    * Metrics checked to show that the traffic going to the new replica is healthy
    * Replace the remaining replicas with the new version if the new replica is healthy
    * Don't continue the deployment if the new replica is not healthy

## Questions / Defuzz / Decisions

* What tools will you use to do the various Kubernetes updates? Helm? Kustomize?
  * Think about the user interface
  * Create an [ADR](/core-engineer/adrs/) with your decision
* How do you deploy replicas without them receiving production traffic?
* How will you deploy prometheus with custom configuration?
    * Maybe a folder in your repo with a script that deploys the helm chart and overrides configuration

## Deliverables (For Epic)

- Publish a custom metric that can be used for the canary check. Do via prometheus compatible library.
- Change the prometheus configuration to scrape the reference application. Have this alert defined as code.
- Validate the health of the canary metrics via an alert manager alert.
- Successfully deploy the canary pod carrying the new version
- Successfully roll out the canary version to other pods when deemed healthy based on the custom metric