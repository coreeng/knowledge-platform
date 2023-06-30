+++
title = "Cluster Config Management"
date = 2022-01-15T11:50:10+02:00
weight = 9
chapter = false
pre = "<b></b>"
+++

### Motivation

Keep resources in sync across many Kubernetes clusters across all environments.

### Requirements

* Keep resources in clusters up to date across clusters in the same environment, possibly running in different providers.
* Ability to select which namespaces and clusters receive which configuration.
* All resources versioned
* Ability to see what version of resources are in each cluster

## Additional Information

Popular tools to implement this:

* Flux
* config-sync from Anthos



