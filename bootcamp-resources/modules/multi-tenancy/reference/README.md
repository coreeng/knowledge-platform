# Bootcamp - Multi Tenancy Reference App

## Environment Setup

### Pre-requisites

- `make`
- `kubectl`

###  Boostrap the cluster 

`make bootstrap`

Note: If the minikube cluster is already running you will be prompted for input to delete or keep the existing cluster.

## How to onboard a new tenant

At the moment the tenant onboarding functionality supports the following features:
- creation of hierarchical namespaces with RBAC functionality for a new tenant
- provision of the monitoring stack if opted in
- updates of the existing per-tenant resources

Not supported yet:
- deletion of the existing per-tenant resources

In order to validate the onboarding manifests:
- run `make validate-onboarding-manifests`

In order to onboard a new tenant:
- add your tenant configuration in [tenant config](/bootcamp-resources/modules/multi-tenancy/reference/reference/onboard/tenant_config.cue)
- `make run`

## How to run some tests locally

### Rbac

`make test-tenant-rbac`

### Network isolation
After onboarding, the following networking rules apply:
- inter namespaces pods cannot communicate with each other
- intra namespace pods can communicate with each other

In order to test this network isolation initial stage run:
- `make test-network-isolation`

In order to create some per-team network policies: 
```
cd reference/teams/team-a
make apply-network-policies
```
This should allow connectivity from team-a's app namespaces to the monitoring namespace. 
You can retest that with the same make task from the main directory:
- `make test-network-isolation`

Note: at the moment no reference app is deployed, we are testing connectivity from a netcat test image to the monitoring namespace for simplicity reasons.


