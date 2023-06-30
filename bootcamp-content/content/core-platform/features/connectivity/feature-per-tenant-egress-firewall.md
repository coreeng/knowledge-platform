+++
title = "Per Tenant Egress Firewall"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation
Allow other areas to open up the full CIDR range of the platform and enforce egress firewalls in the cluster.

Prevents connectivity and reachability being setup more than once i.e. once one tenant requests access, they all get it, but it is default deny by default.

## Requirements
* Per tenant, context aware, firewall rules for egress out of the platform.
* To enable connectivity to other systems that need to open up the full CIDR of the platform

## Additional Information
Can be implemented with CIlium.

Egress control is controlled by a mix of cluster wide network policies and network policies.

Global whitelist to allow to known destinations as cluster wide policies.

A common requirement is that a third party stakeholder e.g. Cyber, want to be able to dynamically block CIDR ranges. A custom solution can be used that takes in requests (e.g. via a queue) that then updates cluster wide network policies. 

To prevent users from creating policies that allow traffic to external sources a policy can be used. A more advanced workflow can be implemented where that policy is checked against allowed external URLs.

How are allowed external URLs allowed? Can be in policy configuration or as a custom resource in the cluster.


