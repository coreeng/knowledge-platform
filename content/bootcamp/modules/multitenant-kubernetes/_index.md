+++
title = "Multi Tenant Kubernetes"
date = 2022-12-22T05:07:10+02:00
weight = 3
chapter = false
+++

## Introduction 

**What should you {{< colour green understand >}} at the end of the week?**

* Kubernetes, as a user, in depth, in a multi tenanted environment 
  * Deployment models: Canary, Blue Green
    * Services, Ingress, DNS
  * Configuration via ConfigMaps, Secrets
  * Multi-tenancy & Tenant boundary: namespace scoping
    * Hierarchical Namespaces
    * RBAC
    * NetworkPolicy
  * Kubernetes users / kubeconf 
* Minimising environment configuration 
  * Templating configuration to minimise changes between environments
  * Service discovery 
  * Secrets 
* The difference between a CECG Core Platform, and plain Kubernetes

**What should you be {{< colour green "able to do" >}} by the end of the week?**

* How to implement Canary Deployments with Kubernetes
  * Combination of Ingress, Services & Deployments
* Implement default deny for your services 
* Multi tenancy
  * Understand Kubernetes RBAC + service accounts + limited access 
    * Namespace vs cluster scoped resources
* Kubernetes Command line tools
  * Helm
  * Kustomize 
