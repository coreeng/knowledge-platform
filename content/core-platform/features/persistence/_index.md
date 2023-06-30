+++
title = " "
date = 2022-01-15T11:50:10+02:00
weight = 17
chapter = false
pre = "<b>Persistence</b>"
+++

# Persistence 

The Core Platform's main purpose is stateless workloads.
So where does persistence live? The answer is normally in a separate cloud account, managed separately.
This won't work for every use case and the Core Platform can be augmented ranging from:

* Zero opinions on persistence. The only feature is connectivity to the persistance tier.
* Paired cloud account provisioning: Provision a cloud account for the tenant and enable connectivity between the tenant's namespaces and a network running in the cloud account
* Templates for use in the sister cloud account for key managed services
* Managed service provisioning in sister cloud account for key managed services 
* Fully managed persistence running inside the Core Platform 