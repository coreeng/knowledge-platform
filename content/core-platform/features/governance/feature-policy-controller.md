+++
title = "Policy Controller"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation
Have a policy controller that can be used to implement governance features.

## Requirements
Pick and install a policy controller such as [JSPolicy](https://www.jspolicy.com/), [OPA](https://open-policy-agent.github.io/gatekeeper/website/docs/), [Kyverno](https://kyverno.io/) . Used internally by the platform team initially, then used to implement the features described below.

## Additional Information
There are many policy controllers in Kubernetes that are implemented using Mutating and Validating webhooks under the covers to provide a nice way to control / modify the Kubernetes resources as they are created. 



