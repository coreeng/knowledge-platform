+++
title = "Core Platform Environments"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation

Have enough core platform environments to implement a P2P for the platform that can safely test, stage, then promote 
changes to tenant facing environments.
Treat any environment used by an end user aka tenant as a production environment.

## Requirements

For each core platform the environments can change based on requirements. The standard ones used by CECG are:
* Sandbox: 
  * Non-tenanted environment
  * Dev account for the platform team, for features that can't be developed on the Dev environment
  * Keeps the Dev environment stable for tenants.
  * May contain multiple instances of the platform, for CI (PR validation)
* Pre-Dev: 
  * Non-tenanted environment
  * Stages platform changes.
* Dev:
  * Tenanted environment
* Pre-Prod:
  * Tenanted environment
  * Stages tenant changes
  * Prod like
* Prod:
  * Tenanted environment
  * Production usage

## Additional Information

{{< figure src="/images/reference/platform/environments.png" >}}

#### Account management

Our recommended approach for mapping accounts to environment is as follows:
1. One account for the **Production** environment. This would likely have the strictest access permissions.
1. One account for **each staging**  environment. For infrastructure, this would be **pre-dev**, for tenants this would be **pre-prod**. Some client might request other environments, see [Platform P2P](/core-platform/features/platform-path-to-prod/).
1. One account for **all dev** environments. We make use of logical separation in the dev environment for testing: functional tests, non-functional tests, integration tests, and showcase. Some clients might request a physical separation.
1. One account for **infrastructure development aka sandbox**. Ideally, we would offer infrastructure developers isolated environments for specific feature development. We do this by provisioning different instances of the platform in the same account.
1. For a multi-region platform, the default approach would be to keep them in the same account per environment.
1. This is considering that each region is simply an additional instance created for resiliency and availability, and they share the same configuration.  If this is not the case, then each region would probably merit its own account.
   1. To reduce the blast radius for each region, the P2P should allow staged rollouts to each region.
1. For a heavily used platform, having multiple regions in the same account can lead to a noisy neighbour scenario (see service quotas and API request limits management below). At this point, splitting them into different accounts would be worth it.
1. For heavily used platforms, even if only single region, we could potentially encounter a noisy neighbour scenario. If this is the case, we could consider creating multiple accounts based on organisational units. This could easily result on operational complexity and management overhead, so it needs to be carefully planned.

#### Justification

* Environment isolation. Each environment requires distinct levels of operation, security and compliance. Account separation is the most powerful and simple way to achieve isolation of resources.
* Security:
  * Each environment can have a different security profile.
  * It facilitate the principle of least privilege access.
  * Simplifies IAM management.
  * Cross environment access (if really needed) needs to be granted explicitly.
* Limit the scope of impact of un/intentional adverse changes
  * Resources in one account are isolated from resources in the other accounts.
  * A compromised environment will not affect other environments.
* Cost management
  * Identify where costs are coming from.
  * Use consolidation of billing accounts for overall costs.
* Service quotas and API requests limits management
  * Accounts have quotas and limits for the different services.
    * Hard limits can’t be increased; Soft limits can be increased.
    * Request throttling
  * Isolating usage per environment type will simplify its management.
  * Usage in an environment will not affect other environments.




