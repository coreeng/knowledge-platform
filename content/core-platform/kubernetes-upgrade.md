+++
title = "Kubernetes Upgrade"
date = 2022-12-28T12:50:10+02:00
weight = 5
chapter = false
pre = "<b>3.4 </b>"
+++

## General

To allow for smooth upgrades between versions, Kubernetes allows API servers, nodes, and clients to be within a certain number of releases of each other. Their version skew policy has all of the details but in summary:
* Control plane versions need to be with in 1 minor version
* Nodes can be up to 2 minor versions behind the API servers. This means you can skip upgrading nodes every other release if you happen to be upgrading API servers 2 or more minor versions.
* Kubectl can be within 1 minor version of the API servers (either older or newer). This is important for CI images, administrators or even tenants that use kubectl to interact with the cluster

## Planning the upgrade
* Research the latest state of the art regarding Kubernetes upgrades and update this process if required
* Identify current and target versions
  * We aim to keep up to date but if there is multiple jumps some work can be saved by having nodes move up 2 versions
* If the current and target versions are more than 2 minor versions apart, create an Epic for each 2 minor version upgrade cycle (e.g. moving from 1.16 to 1.22 would result in 3 epics (1.16 → 1.18, 1.18 → 1.20, and 1.20 → 1.22). The reason for this breakdown is due to the Kubernetes' version skew policy, which is described above.
* Each upgrade epic should be broken down into the following stories:
  * ChangeLog analysis;  it’s important to note any changes that will affect tenants or administrators. Our experience has shown the following common themes:
    * API deprecations and removals
    * Command line flag deprecations and removals
    * Metrics renames and removals.
  * Judicious use of GitHub’s search feature to search across your entire organisation is strongly advised!
  * If any ChangeLog entries are identified as affecting components, tenants or administrators, add a story for each one
  * Upgrade control plane to current version + 1
  * Upgrade control plane to current version + 2 (if required)
  * Upgrade nodes to current version + 2 (or + 1)
  * Upgrade CI agents to include a compatible version of kubectl + update onboarding documentation on how to install the correct version of kubectl

## Case study 1: GKE cluster upgrade guide

### Overview

A cluster's control plane and nodes are upgraded separately.
The control plane needs to be upgraded first but should not be more than 2 minor versions newer than the nodes.
This is because the [Kubernetes version and version skew support policy](https://kubernetes.io/docs/setup/release/version-skew-policy/)
guarantees that control planes are compatible with nodes up to two minor versions older than the control plane.
For example, Kubernetes 1.23 control planes are compatible with Kubernetes 1.21 nodes.

Once the control plane is upgraded, node pools will be [upgraded automatically](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-upgrades#upgrading_automatically) separately - which is also the default option.
Nodes are [surge upgraded](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-upgrades#surge) one zone at at time and one node at a time.
When a node is upgraded, it is cordoned to ensure no more pods are scheduled onto it then drained.
Evicted pods are automatically rescheduled to other nodes.

### Before the upgrade

The steps required to upgrade the cluster will depend on the gap between the current and target versions.
Upgrading patch versions should require very few changes, but upgrading minor versions might bring some incompatibilities.

Before doing an upgrade, you'll need to understand the impact of any API changes and then communicate the impact to the tenants
1. Check out the [Kubernetes changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/). 
   This gives you a lengthy description of all the changes that have been made to Kubernetes since the last release. 
   It's useful to skim through, but IMO not designed to be read in detail.
2. Look for APIs changes and deprecated features as this may require changes for both platform and tenant services. For instance tenant teams might require to
   * change their deployment spec if they have used a feature that has moved from alpha to beta.
   * Move away from deprecated features that will be removed in the target release.  To understand the impact you can run the following query in the log explorer in the target GCP project. It shows the usage of deprecated APIs that are used by non-Google-managed user-agent. **Any usage of removed APIs usage must be eliminated before upgrading the cluster**.
     ```
     resource.type="k8s_cluster"
     labels."k8s.io/removed-release"="1.25"
     protoPayload.authenticationInfo.principalEmail:("system:serviceaccount" OR "@")
     protoPayload.authenticationInfo.principalEmail!~("system:serviceaccount:kube-system:")
     ```
3. Look for deprecated or removed APIs that may no longer be compatible with your helm charts (**NB**: Feel free to skip this part if you are **not** using helm to deploy platform components). You can use the [helm-mapkubeapis](https://github.com/helm/helm-mapkubeapis) plugin to find out. Here is a small script that can help you with that:
    ```bash
    helm_charts=$(helm list -A --output json | jq -r '.[] | .name + "@" + .namespace')
    for item in ${helm_charts[0]}; do
      IFS=@ read chart namespace <<< ${item}
      echo -e "\n==== Chart: $chart, Namespace: $namespace"
      helm mapkubeapis -n $namespace $chart --dry-run
    done
    ```
4. That being said some of the changes can still go undetected. The [kubernetes deprecations information](https://cloud.google.com/kubernetes-engine/docs/deprecations#deprecations-information) is also a good source of information. 
5. If any upgrades need tenant team's effort, you should communicate with the team clearly and agree a time-frame for the upgrade, you can use the template at [communicate with tenants](###tenant-comm-template) as a starting point:
   * For non-live upgrade you can do it at via IM system (e.g. slack or MS Teams).
   * For prod upgrade ideally you want to communicate with the tenant teams via email DL.
6. If the changes are non-trivial you will need to agree a time-frame with the tenant teams to make their changes before you can proceed with the upgrade.
7. Make sure that on your workstation you have `kubectl` tool that is within the [supported version skew range](https://kubernetes.io/releases/version-skew-policy/#kubectl) against target release. You will need to [upgrade your kubectl](https://kubernetes.io/docs/tasks/tools/) if it falls too behind the target release.

### Upgrade

#### Upgrading patch versions

These upgrades are automatically done by GKE.

#### Upgrading minor versions

These upgrades should be controlled and staged across environments starting from lower environments.

You can upgrade the gke cluster of a specific environment by bumping the gke module's `gke_kubernetes_version` in the `infra/envs/${env}/gke/main.tf` file.

For example, to upgrade the cluster to version from `1.24` to `1.25` in the prod environment:

```diff
--- a/infra/envs/da-platform-prod/gke/main.tf
+++ b/infra/envs/da-platform-prod/gke/main.tf
@@ -44,7 +44,7 @@ provider "google-beta" {

 module "gke" {
   source                 = "../../../modules/gke"
-  gke_kubernetes_version = "1.24"
+  gke_kubernetes_version = "1.25"

   gcp_region   = local.gcp_region
   gcp_project  = local.gcp_project
```

Checklist before/after upgrading the cluster - feel free to attach this to the change request for a comprehensive list of TODOs.

* Make sure that the upgrade has been tested thoroughly in the lower environments.
* Raise change request in Sparks (**applicable to prod only**).
* Communicate with the tenants teams to agree a time-frame for the upgrade to avoid unecessary disruption.
* Silence the control-plane related alerts (**applicable to prod only**), since certain control plane components will be unavailable during the upgrade. Since this is customer specific we won't go into details of how to do this.
* Apply the version bump, and roll the change up to the desired environment - either through a CI system (highly recommended) or manually.
* Run e2e acceptance tests to ensure that the upgrade has not broken anything major.
* Un-silence the alert you just created in your monitoring system before the upgrade (**applicable to prod only**).
* No alerts should be triggered during the upgrade.
* Update your helm charts to remove deprecated or removed Kubernetes API from your Helm metadata
  You can use the [helm-mapkubeapis](https://github.com/helm/helm-mapkubeapis) plugin to do that as shown below:
    ```bash
    helm_charts=$(helm list -A --output json | jq -r '.[] | .name + "@" + .namespace')
    for item in ${helm_charts[0]}; do
      IFS=@ read chart namespace <<< ${item}
      echo -e "\n==== Chart: $chart, Namespace: $namespace"
      helm mapkubeapis -n $namespace $chart
    done
    ```  
* Communicate with the tenants teams that the upgrade has been completed.
* Remediate any issues that may have been caused by the upgrade.

### Tenant comm template

You can use the following slack message template to communicate with tenant teams for non-live upgrade:

```markdown
Hey Team,

Just wanted to give everyone a heads up that we will be upgrading our Kubernetes cluster from version ${current_version} to ${target_version} in ${env} on ${date}. The upgrade is part of our ongoing efforts to keep our infra up-to-date and maintain a secure and reliable platform-as-a-service for you.

The upgrade is scheduled to take place on ${date_and_time}, and we anticipate that it will take approx. ${duration} to complete. During this time, the Kubernetes API server and other control plane components will be unavailable. That being said you are **not** expected to experience interuptions of your workloads running on the cluster.

To minimise the impact of the upgrade, we recommend you take the following steps:

* Replace API X with API Y (with detailed instructions if it is not straight forward).
* etc
* Let us know if you are planning for any big launches or expecting any traffic spikes during the upgrade - we can always reschedule the upgrade if needed.


You might need to upgrade your kubectl if [the version is a bit of behind](https://kubernetes.io/releases/version-skew-policy/), which you will get warning or error from the command line. In that case you can follow the [official guide](https://kubernetes.io/docs/tasks/tools/) for upgrade.

_OPTIONAL: You can mention the features that will be supported out-of-box by the new version_

Thanks for your understanding and cooperation!

```
You can use the following email template to communicate with the tenant teams for live upgrade:

```markdown
Hi team,

I'm writing to inform you that we will be performing a scheduled upgrade of our Kubernetes cluster from version ${current_version} to ${target_version} on ${date}.

This upgrade is part of our ongoing efforts to keep our platform up-to-date and maintain a secure and reliable platform-as-a-service for you.

THe upgrade is scheduled to take place on ${date_and_time}, and we anticipate that it will take approx. ${duration} to complete. During this time, the Kubernetes API server and other control plane components will be unavailable. That being said you are **not** expected to experience interuptions of your workloads running on the cluster.

To minimise the impact of the upgrade, we recommend you take the following steps:

* Replace API X with API Y
* etc
* Let us know if you are planning for any big launches or expecting any traffic spikes during the upgrade - we can always reschedule the upgrade if needed.

Here are some of the new features that will be supported out-of-box by the new version:
* _OPTIONAL: You can mention the features that will be supported out-of-box by the new version_

We will be monitoring the upgrade process closely and will keep you updated on any changes or issues that arise during the upgrade.

If you have any questions or concerns about the upgrade, please do not hesitate to reach out to us. We are here to help and ensure that the upgrade process goes as smoothly as possible.

Thank you for your cooperation and understanding as we work to maintain and improve our Kubernetes platform.
```