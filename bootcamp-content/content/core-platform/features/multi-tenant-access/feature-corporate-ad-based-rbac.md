+++
title = "Corporate AD based RBAC"
date = 2022-01-15T11:50:10+02:00
weight = 2
chapter = false
pre = "<b></b>"
+++

## Motivation

- Use an organisation's standard group management to control access in the core platform.
- Delegate to the tenants the membership of their groups.

## Requirements

* Kubernetes access based on corporate AD groups
* Tenants control the users in the AD groups

## Additional Information

* Map Kubernetes roles to corporate AD groups. Requires Corporate SSO integration
* Decide if there is a group per environment or one for all.
* How does it work for CI? Same groups or different?

Authorization is built on top of Authentication, typically implemented via the [SSO Feature](/core-platform/features/connected-kubernetes/feature-sso-integration)

Once a user has been identified via authenticating, then they need to be authorized.

This is done via tenant onboarding creating a set of ClusterRoles, ClusterRoleBindings, Roles, and RoleBindings.

Assuming that [Hierarchical Namespaces](https://github.com/kubernetes-sigs/hierarchical-namespaces) are used, then a Role and RoleBidning can be created inside their parent namespace that is propagated to all sub namespaces.

The end to end journey of SSO, Tenant Kubernetes Access and mapping to AD groups is depicted below:

{{< figure src="/images/reference/platform/e2e-sso-journey.png" >}}

Exactly what the granularity of the groups and roles are will depend on the client’s RBAC requirements.

## Questions

* Do they need separate groups for read only and read write?
* How granular is the access? Per application? Per tenant? Per department?
* How does authentication work for agents such as CI/CD tooling?

