+++
title = "Paired Cloud Account"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation

Provide a seamless way for a tenant to provision persistent services in a cloud account that can be accessed from the Core Platform
Avoid a lengthy account provisioning and connecting to the Core Platform process

## Requirements

* Cloud Account provisioned per tenant that only they have access to
* Cloud Account has network connectivity to and from the tenant's namespaces
* Kubernetes service accounts can assume IAM roles in the paired cloud account for passwordless authentication
* Cloud Account lifecyle is fully managed by the tenant apart from the base network connectivity

## Questions

Does the platform team have access to provision cloud accounts?

Should the cloud accounts be able to contain arbitrary cloud resources or only specific services with blue prints?
