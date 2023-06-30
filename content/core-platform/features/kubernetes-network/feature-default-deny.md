+++
title = "Default Deny"
date = 2022-01-15T11:50:10+02:00
weight = 3
chapter = false
pre = "<b></b>"
+++

## Motivation

* Show security stakeholders how firewalling can be implemented at the kubernetes layer
* Stick to the principle of least privilege access applied to internal networking
  * In a multi-tenant environment the tenants can not impact each-other accidentally and should work as it they were completely isolated. 
    That’s true for resources as well as network. For example, if a tenant pod is compromised, it should not be able to maliciously attack other pods/tenants, which can be achieved with default deny, as you’ll need to whitelist all the connectivity needed for any given application.

## Requirements

* Default connectivity outside, and between namespaces
* Tenants can open up inbound connectivity to their namespaces

## Additional Information
* Depends on Tenant Managed Firewalls
* Much easier to do initially rather than retrofitting!
* Which parts are defaulted?
  * Inside the cluster aka pod IPs, service IPs:
    * Egress is open
    * Ingress is denied
  * Outside the cluster:
    * Ingress is only via Ingress LBs/ICs
    * Egress has a whitelist either globally or per “tenant”
      * Internal to the corporation
        * Everyone doesn’t have open up to standard internal services running outside of the cluster
        * Other outbound FQDNs or CIDR ranges should be restricted per tenant via “Per Tenant Egress Firewall”. If default deny is implemented before “Per Tenant Egress Firewall” then likely this will just have to open for now. 
      * External to the Internet
        * Some companies have locked down outbound connectivity i.e. via some gateway / proxy. If that’s the case then it might make sense to just open up to that, and let any other rules go there. 
* Who manages what?
  * Inside the cluster aka pod IPs, service IPs:
    * Managed by the teams via standard network policy
  * Outside the cluster:
    * Initially by the platform team, until “Per Tenant Egress Firewall”
      * Implemented by a combination of Cluster wide network policies and policies that prevent tenants from doing things.



