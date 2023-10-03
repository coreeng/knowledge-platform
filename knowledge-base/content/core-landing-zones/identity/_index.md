+++
title = "Access Management & Identity"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
+++

## Identity Federation

A cloud provider is unlikely to be an Orgsanisation's [Identity Provider](https://en.wikipedia.org/wiki/Identity_provider) (IdP). Typical identity providers are:

* Azure AD
* Google Workspace
* Ping
* Okta
* Cyber Ark

All the major cloud providers support importing identity from external IdPs for users and groups.

Access in the Cloud should be based of groups in the IdP.

### MFA 

Access to cloud resources by human's should be via [multi factor authentication](https://en.wikipedia.org/wiki/Multi-factor_authentication).

## Temporary Elevated Access

* How do users gain temporary elevated access? 
* What is the workflow?
* Who are the approvers?
* How is it audited?
* What tooling is involved? e.g.
  * CyberArk
  * Beyond Trust
  * OKTA

##  Cloud IAM setup

* What roles in the cloud are needed for an initial role out?
* How do they map to groups in the IdP?