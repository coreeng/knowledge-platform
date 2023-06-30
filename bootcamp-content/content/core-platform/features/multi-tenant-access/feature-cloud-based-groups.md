+++
title = "Cloud based Auth and Group management"
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b></b>"
+++

## Motivation

* A Core Platform with no external dependencies from the Cloud Provider
* Enable easy migration to SSO and AD Based RBAC 

## Requirements

* Authenticate via the cloud provider 
* Kubernetes access based on user and group information captured in the client 
* Groups for tenants automatically created when tenants onboard
* If possible for the cloud provider:
  * Group membership managed by the tenant teams 

## Questions

* Are separate groups required for read only and read write?
* How granular is the access? Per application? Per tenant? Per department? This should be defined in [multi-tenant access](./feature-tenant-kubernetes-access).
* How does authentication work for agents such as CI/CD tooling?
  * If this is a single cloud provided Core Platform, the CI/CD tooling should also run in the cluster e.g. Tekton, solving this issue 


## Additional Information

Any solution should take into account how easy it would be to migrate to [SSO](../../connected-kubernetes/feature-sso-integration) along with [AD based RBAC](../feature-corporate-ad-based-rbac)

### GCP

GKE authentication uses the `gke-gcloud-auth-plugin` `gcloud` plugin. These must be installed.
Users will authenticate with their GCP user.
RBAC can then be managed with Google Groups or IAM.

#### RBAC with Google Groups (preferred)

Every user needs: `container.clusterViewer`, nothing else should be added to RBAC.

`ClusterRole`s and `Role`s are created that are agnostic of Google Cloud.
The responsibility lies with the [Multi-Tenant Kubernetes Access](./feature-tenant-kubernetes-access) feature.

The Cloud coupling is limited to the `RoleBinding`s.

Example role:

```
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: platform-accelerator 
  name: pod-reader
rules:
- apiGroups: [""] # "" indicates the core API group
  resources: ["pods"]
  verbs: ["get", "watch", "list"]
```

Example role binding:


```
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: pod-reader-binding
  namespace: platform-accelerator 
subjects:
- kind: Group
  name: platform-accelerator-admin@cecg.io 
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```

### How can the tenant manage membership of their own group?

- Groups for each tenant are created with the onboarding automation
- The end tenant is given manage access to manage just their google group 
- Owner is the platform engineering team 