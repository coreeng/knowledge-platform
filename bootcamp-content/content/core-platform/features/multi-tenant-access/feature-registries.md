+++
title = "Multi-Tenant Registries"
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b></b>"
+++

## Motivation

* Provide isolated registries for tenants 
* Follow principle of one tenant can't affect another (e.g. by pushing to a shared registry)

## Requirements

* Each tenant should get their own registry that they can push to from CI and pull from their namespaces
* One tenant should not be able to push images to another tenant's registry
* (Optional) one tenant should not be able to pull images from another tenant's registry
* Platform to be only able to deploy from these registries

## Questions

How will the tenants authenticate?
Will the authentication mechanism be cloud provider agnostic?
Should pull be authenticated?