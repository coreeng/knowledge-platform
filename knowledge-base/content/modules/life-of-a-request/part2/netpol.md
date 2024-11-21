+++
title = "What are Network Policies?"
weight = 3
chapter = false
+++

#### Overview

The **web-app** successfully resolved the IP address for the **user-api** service and is going to send an HTTP request. The **user-api** service is deployed into a different namespace - **backend**. We need to make sure that such communication is allowed.

Network policies in Kubernetes are crucial for managing the communication between pods within a cluster. They provide a way to control how groups of pods can communicate with each other and with other network endpoints.

We have two sets of pods: **web-app** and **user-api**. We want to restrict the **web-app** pods to only communicate with the **user-api** pods. Also, we don’t want **web-app** to be able to communicate with 3rd party APIs directly.

![Network Policy for Web App](/images/loar/2-4.png)
_Figure 2-4. Network Policy for Web App_


#### Default Deny

Enabling a **Default Deny** policy in Kubernetes is an important aspect of securing a cluster by controlling network traffic between pods. By default, Kubernetes allows all traffic between pods in the same namespace, but with a Default Deny policy, you explicitly deny all traffic unless otherwise specified.

In the following example, we create a default deny NetworkPolicy that denies all **outcoming** calls from the **frontend** namespace pods.

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: default-deny
  namespace: frontend
spec:
  podSelector: {}
  policyTypes:
    - Egress
```

#### Allowing DNS Traffic

When you apply the default deny policy in the **frontend** namespace, you get your DNS request blocked as well. So, DNS resolution is not working for **frontend** pods.

```
$ kubectl exec -ti dnsutils -n frontend -- nslookup cecg.io  

;; connection timed out; no servers could be reached
```

We need to explicitly allow UDP traffic to port **53** to **kube-dns**. 

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-dns-access
  namespace: frontend
spec:
  podSelector:
    matchLabels: {}
  policyTypes:
  - Egress
  egress:
  - to:
    - namespaceSelector:
        matchLabels:
          kubernetes.io/metadata.name: kube-system
      podSelector:
        matchLabels:
          k8s-app: kube-dns
    ports:
    - protocol: UDP
      port: 53
```

After that the DNS resolution is working correctly:

```
$ kubectl exec -ti dnsutils -n frontend -- nslookup cecg.io  
Server:		10.96.0.10  
Address:	10.96.0.10\#53

Non-authoritative answer:  
Name:	cecg.io  
Address: 75.2.60.5
```

#### Allowing Traffic To Pods

Get into a **web-app** pod and try to reach a **user-api** pod:

```
$ kubectl exec -it web-app-6698c689d-qwhh7 -n frontend -c webapp -- curl –-max-time 2 http://user-api.backend:8080/user/123  

curl: (28) Connection timed out after 2002 milliseconds
```

The communication is blocked. We need to explicitly allow the **web-app** pods to communicate with the **user-api** pods by creating a Network Policy.

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-web-app-to-call-user-api
  namespace: frontend
spec:
  podSelector:
    matchLabels:
      app: web-app
  policyTypes:
    - Egress
  egress:
    - to:
      - namespaceSelector:
          matchLabels:
            kubernetes.io/metadata.name: backend 
        podSelector:
          matchLabels:
            app: user-api
```

You can test the policy by attempting to access the **user-api** pods from the **web-app**.

```
$ kubectl exec -it web-app-6698c689d-qwhh7 -n frontend -c webapp -- curl –max-time 2 http://user-api.backend:8080/user/123  
{"id":"123","name":"User"}
```

#### Troubleshooting

Debugging Kubernetes network policies can be challenging, especially since they directly control network traffic within the cluster.

##### Understand the rules

Review your network policies carefully, as a small configuration mistake can block unintended traffic.

Ensure you’re familiar with key properties in the policy, such as **podSelector**, **namespaceSelector**, **policyTypes**, **ingress**, and **egress** rules.

Note that network policies are **additive**; if you have multiple policies, Kubernetes will apply them together, and traffic will only be allowed if at least one policy permits it.

Make changes to policies in small steps, testing connectivity after each change. This makes it easier to identify which part of the policy is causing issues.

Use **describe** command to inspect specifics of a policy:

```
$ kubectl describe networkpolicy default-deny -n frontend

Name:         default-deny
Namespace:    frontend
Created on:   2024-10-31 13:39:51 +0200 EET
Labels:       <none>
Annotations:  <none>
Spec:
  PodSelector:     <none> (Allowing the specific traffic to all pods in this namespace)
  Not affecting ingress traffic
  Allowing egress traffic:
    <none> (Selected pods are isolated for egress connectivity)
  Policy Types: Egress
```

##### Use debugging pods

Test network connections between pods. Deploy pods specifically for testing purposes, such as:

- [netshoot: a Docker + Kubernetes network troubleshooting swiss-army container](https://github.com/nicolaka/netshoot): A network debugging pod that includes tools like **curl**, **ping**, **nc**, **ip**, **tcpdump**, and **nslookup**.
- [BusyBox](https://busybox.net/): A lightweight debugging pod with networking utilities.

Use these pods to try pinging or using curl to test access between namespaces or within the same namespace:

```
$ kubectl run tmp-shell --rm -i --tty --image nicolaka/netshoot
```

##### Run checks inside the pod

If your policies should allow traffic between pods, but they aren’t connecting, you can use **kubectl exec** to test connectivity:

```
$ kubectl exec -it <source-pod> -n <namespace> -- curl http://<target-pod>.<target-namespace>.svc.cluster.local
```

Run commands like **ping**, **curl**, or **telnet** to test network access from one pod to another.

##### Check CNI logs {#check-cni-logs}

Most network plugins (e.g., Calico, Cilium, Weave, or Flannel) provide logs or diagnostic commands that can show policy-related events.

For example, if you’re using Calico you can get the logs from Calico’s pods to see if there are policy enforcement errors:

```
$ kubectl logs -n kube-system -l k8s-app=calico-node 

Defaulted container "calico-node" out of: calico-node, upgrade-ipam (init), install-cni (init), mount-bpffs (init)
2024-11-01 13:30:55.077 [INFO][83] monitor-addresses/autodetection_methods.go 103: Using autodetected IPv4 address on interface eth0: 192.168.49.2/24
2024-11-01 13:31:32.770 [INFO][80] felix/summary.go 100: Summarising 11 dataplane reconciliation loops over 1m4.5s: avg=11ms longest=38ms ()
2024-11-01 13:31:55.083 [INFO][83] monitor-addresses/autodetection_methods.go 103: Using autodetected IPv4 address on interface eth0: 192.168.49.2/24
```

##### Use simulation tools {#use-simulation-tools}

Consider using [Network Policy Editor for Kubernetes](https://editor.networkpolicy.io/) , which can help visualise network policies to better understand them.

![Screenshot of Network Policy Editor](/images/loar/2-5.png)
_Figure 2-5. Screenshot of Network Policy Editor_

#### Further Reading {#further-reading-5}

- [Declare Network Policy | Kubernetes](https://kubernetes.io/docs/tasks/administer-cluster/declare-network-policy/)
- [Kubernetes — Debugging NetworkPolicy (Part 1\) | by Paul Dally | FAUN — Developer Community](https://faun.pub/debugging-networkpolicy-part-1-249921cdba37)
- [Network Policy Editor for Kubernetes](https://editor.networkpolicy.io/)
- [Example recipes for Kubernetes Network Policies that you can just copy paste](https://github.com/ahmetb/kubernetes-network-policy-recipes) 