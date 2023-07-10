+++
title = "Service Catalog"
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b></b>"
+++

## Motivation

Discover any service running on the core platform, including who owns it

## Requirements

* All tenant applications in the service catalog
* Automatically add any new onboarded application with no extra effort for the tenant
* All platform components in the service catalog
* Access for all tenants via same SSO mechanism as the core platform
* If [application blueprints](./feature-application-blueprints) has not been implemented yet provide at least one template, even if it is just a catalog info yaml

## Questions

How do things get into the service catalog?

* Should the only way to deploy to a core platform be via a service template that is automatically added?
* Can components be added retrospectively?

How does this interact / compare with an existing configuration management database (CMDB) such as service now? 



