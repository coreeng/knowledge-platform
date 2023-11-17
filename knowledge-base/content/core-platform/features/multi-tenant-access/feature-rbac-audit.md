+++
title = "RBAC Audit"
date = 2022-01-15T11:50:10+02:00
weight = 6
chapter = false
pre = "<b></b>"
+++

## Motivation

* Address lack of confidence in who can do what in clusters.

Reports for team owners to understand who can do what in a cluster.

## Requirements
* Query OIDC provider and Kubernetes API to report on who can do what in a cluster

## Additional Information
A common problem at clients is getting to the point where it is unclear who can do what in the clustes.

It is possible to engineer a tool that enables you to see end to end what principals/subjects can do in given Kubernetes cluster and why/how they can achieve it.

Example: User `alice@example.com` can `get secrets` in namespace `app-stream` via membership in group `team-stream` which is bound via `roleBinding team-admin` in namespace `app-stream` to `clusterRole team-admin` which contains that rule.

### Problem Definition 
Kubernetes API does not make it easy to see what one principal can do in a cluster. Its permissions model is built in reverse order, which means that starting point is a resource from which we can partially backtrack who has access to it.

This is, of course, a well known problem, so there are a few projects that are trying to make inspections easier like `kubectl auth can-i` or the plugin [who-can](https://github.com/aquasecurity/kubectl-who-can). Unfortunately, these tools are not trying to escalate their permissions and thus are a good first step, but not enough for full scope audit of RBAC.

Role based access also means that groups are preferred over individual users as it makes it tedious to manage users in Kubernetes. This usually means that cluster administrators set up integration with an OIDC provider where group membership is managed. As a result it is not possible to see who is member of a group. To make inspection even harder groups are usually referenced via its ID (usually in the form of UUID) which is not easily identifiable by human auditor.

Impact of this design leads to:
* missing ability to audit what one person/principal/subject can do in cluster
* misconfiguration of rules goes undetected and enables unwanted privilege escalation
* leftover bindings referring to Cluster(Roles) that does not exist
* leftover subject referred in bindings
* principals not removed from groups in OIDC as relation between group membership (in OIDC) is disconnected from permissions granted to that group

### Relationship between entities

{{< figure src="/images/reference/platform/k8s-rbac.png" >}}

Example of all Kubernetes RBAC objects we refer in this project:

Role 
```
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: secret-reader
  namespace: example
rules:
- apiGroups:
  - ''
  resources:
  - 'secrets'
  verbs:
  - 'get'
```

ClusterRole
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: admin
rules:
- apiGroups:
  - '*'
  resources:
  - '*'
  verbs:
  - '*'
```

ClusterRoleBinding for ClusterRoles with User kind of Subject
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-crb
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: admin
subjects:
- name: alice@example.com
  apiGroup: rbac.authorization.k8s.io
  kind: User
```

RoleBinding for ClusterRole with ServiceAccount and User kind of Subjects

```
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: admin-rb
  namespace: example
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: admin
subjects:
- name: default
  kind: ServiceAccount
  namespace: example
- name: alice@example.com
  apiGroup: rbac.authorization.k8s.io
  kind: User
```

RoleBinding for Role with Group kind of Subject
```
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: secret-reader-rb
  namespace: example
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: secret-reader
subjects:
- name: bcb9875c-190d-43a6-985b-422e53f7beef
  apiGroup: rbac.authorization.k8s.io
  kind: Group
```

ServiceAccount
```
apiVersion: v1
kind: ServiceAccount
metadata:
  name: default
  namespace: example
```

Azure Group `bcb9875c-190d-43a6-985b-422e53f7beef` and Azure User `alice@example.com`
are resources managed in Azure Active Directory.

> Subject as a resource `kind` does not exist. Referring some resource eg `Role` in
> `RoleBinding` is not checked by API and that `Role` may not exist. Same goes for
> `ClusterRole`s and all kinds of Subjects.

To be able to see which rules (from Roles and ClusterRoles) apply to which users
we need to resolve all these relations.


### Data Collection
There are 2 systems we need to query:

#### Kubernetes API
1. Read all RoleBindings and ClusterRoleBindings to create list of all Subjects
mentioned in them. This enables us to see all Roles (and their namespaced rules)
and ClusterRoles (and their rules) that apply to each Subject.
1. Try to escalate permissions using impersonate, bind, escalate or any other
possible route.

As we mapping bindings and escalations we build trace (list of hops) from Subject to each rule.

Example of resolved Group Subject using resources mentioned above:

```
{
    "kind": "Group",
    "name": "bcb9875c-190d-43a6-985b-422e53f7beef",
    "rules": [
        {
            "verbs": [
                "get"
            ],
            "apiGroups": [
                ""
            ],
            "resources": [
                "secrets"
            ],
            "namespace": "example",
            "trace": [
                {
                    "kind": "Role",
                    "name": "secret-reader",
                    "namespace": "example"
                },
                {
                    "kind": "RoleBinding",
                    "name": "secret-reader-rb",
                    "namespace": "example"
                },
            ]
        },
    ]
}
```

> Notice that rule now contains additional field `namespace` as it applies only there.

Once we resolve and walk all escalations paths we have complete picture of what each `Subject` can do in Kubernetes cluster.

#### OIDC provider (Azure)
Last step of collection is to find out details of Azure Groups as UUID reference in `Group` kind of `Subject` is not helpful for human auditor and we do not know who is member of this group.

After we receive these detail from Azure, we can extend our Kubernetes data with new `User` kind of subject and enrich current data.

Example assuming all above and that Azure Group with UUID `bcb9875c-190d-43a6-985b-422e53f7beef` has display name `Dragons` and has only one member with display name `Bob Dylan` and mail `bob@example.com`

```
{
    "kind": "Group",
    "name": "bcb9875c-190d-43a6-985b-422e53f7beef",
    "displayName": "Dragons",
    "rules": [
        {
            "verbs": [
                "get"
            ],
            "apiGroups": [
                ""
            ],
            "resources": [
                "secrets"
            ],
            "namespace": "example",
            "trace": [
                {
                    "kind": "Role",
                    "name": "secret-reader",
                    "namespace": "example"
                },
                {
                    "kind": "RoleBinding",
                    "name": "secret-reader-rb",
                    "namespace": "example"
                },
            ]
        },
    ]
}
```

and new `User` kind of `Subject`

```
{
    "kind": "User",
    "name": "bob@example.com",
    "displayName": "Bob Dylan",
    "rules": [
        {
            "verbs": [
                "get"
            ],
            "apiGroups": [
                ""
            ],
            "resources": [
                "secrets"
            ],
            "namespace": "example",
            "trace": [
                {
                    "kind": "Role",
                    "name": "secret-reader",
                    "namespace": "example"
                },
                {
                    "kind": "RoleBinding",
                    "name": "secret-reader-rb",
                    "namespace": "example"
                },
                {
                    "kind": "Group",
                    "name": "bcb9875c-190d-43a6-985b-422e53f7beef",
                    "displayName": "Dragons"
                }
            ]
        },
    ]
}
```

### Data Presentation
Once we have all data about each principal/subject we have can present data in various way.

#### Generate (wiki) page where each person can review their (team) access
Example:
Bob Dylan (bob@example.com) 

| Namespace | Verbs | APIGroups | Resources | ResourceNames | Trace |
|-----------|-------|-----------|-----------|---------------|-------|
| example | get | "" | secrets | | {{< list "kind: Role, name: secret-reader, namespace: example" "kind: RoleBinding, name: secret-reader-rb, namespace: example" "kind: Group, name: bcb9875c-190d-43a6-985b-422e53f7beef, displayName: Dragons"  >}}|


We can also group these per team if we find a way how to automatically detect a team. Then we can flag anomalies/differences between team members.

#### Raise alerts for predefined rules

We can have some config where we define various rules aiming at sensitive/powerful permissions like

```
[
  {
    "verbs": ["*"],
    "apiGroups": ["*"],
    "resources": ["*"],
    "namespaces": ["kube-system", ""]
  },
  {
    "verbs": ["delete"],
    "apiGroups": [""],
    "resources": ["namespaces"],
  }
]
```

Our auditing tool loads rules and raises alerts if any of the User subjects have them.

### Good To Read
* `impersonation` [docs](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#user-impersonation)
* not well-known verbs `escalate` and `bind` are a bit described in [Escalating Away](https://raesene.github.io/blog/2020/12/12/Escalating_Away/) and [Kubernetes RBAC Security Pitfalls](https://certitude.consulting/blog/en/kubernetes-rbac-security-pitfalls/)
* nice [example](https://raesene.github.io/blog/2021/01/16/Getting-Into-A-Bind-with-Kubernetes/) of how to use `bind` for privilege escalation



