+++
title = " "
date = 2022-01-15T11:50:10+02:00
weight = 5
chapter = false
pre = "<b>Kubernetes Network</b>"
+++

# Kubernetes Network

Based on the requirements choose the CNI and the features that are to be built on top of the CNI e.g. 
* [Default Deny](./feature-default-deny/)
* [IPV6](./feature-ipv6/)
* [Tenant Managed Firewalls](./feature-tenant-managed-firewalls/)


## Networking Model
See [Services, Load Balancing, and Networking](https://kubernetes.io/docs/concepts/services-networking/).

Understanding the Kubernetes networking model is important for almost any user of a cluster. If only one thing is learned about networking in Kubernetes, this would be it.

* Every Pod has its own routable IP within the cluster.
* Containers on a pod share the same IP - that is, they live in the same network namespace. They can communicate to each other via localhost.
* NAT is explicitly forbidden for pod to pod communication within a cluster.

### Services
See Services, Load Balancing, and Networking .

The fundamental unit of network composition in Kubernetes is the Service.

* A Service represents a group of pod IPs.
* A Service has its own routable IP, referenced usually as a virtual IP (VIP) since it isn’t assigned to a single resource.
* When a pod wants to interact with a group of other pods, such as the replicas in a deployment, it will typically use the Service IP rather than having to figure out the individual pod IPs itself.
* Kubernetes will load balance requests to all the pods that belong to a Service.
* In addition, kubernetes DNS will map the **service.namespace** hostname to the service IP. This provides basic service lookup for clients. They only need to know the service and its namespace rather than having to worry about IPs.
* The Service VIP is supposed to be stable - it should never change dynamically. This means we don’t need to worry about stale DNS caching and the like.

There are a few caveats to be aware of with Services:
* Load balancing is implementation specific. In its most basic form, using kube-proxy, it performs a random allocation in iptables. During rolling deployments, this can lead to imbalances in connections. This can be mitigated somewhat by not maintaining long lived (>30s) connections in clients.
* There are many different service types, some can even serve external traffic via the LoadBalancer service type. Typically though, we use them for internal load balancing only and rely on Ingress for external traffic.
* Services have many optional features that have been added over the years, like topology aware traffic routing. However a lot of these features only work on kube-proxy and not on other implementations. It is largely implementation specific at this point how your traffic will be handled. 

### Ingress
See [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/) .

Ingress is the primary mechanism for supporting external layer 7 (http) traffic into a cluster.

* Conceptually it defines a reverse proxy mapping - virtual host and path to target Service in the cluster.
* It is heavily implementation specific and the basic Ingress spec is quite limited. This has led to many divergent implementations, even completely different resources. Very often there will be a large number of annotations on an Ingress to dictate its behaviour.
* It is likely in the future Ingress will be superseded by the Gateway API. See [Introduction - Kubernetes Gateway API](https://gateway-api.sigs.k8s.io/) . This is largely inspired by Istio and Contour’s work and is managed by SIG-NETWORK.
* There is no de facto ingress implementation. We have many options to choose from, including managed ingress provided by GKE/EKS.

### DNS

See [DNS for Services and Pods](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/).

Cluster DNS is typically provided by either kube-dns or CoreDNS. There is a [DNS spec](https://github.com/kubernetes/dns/blob/master/docs/specification.md) which covers exactly what DNS should provide for Services.

For services:

* Service IPs get mapped to `<service>.<ns>.svc.cluster.local.`
* Each port of a Service gets a SRV record, like `_https._tcp.kubernetes.default.svc.cluster.local.`

And pods:

* Each pod gets <pod-ip-address>.<ns>.pod.cluster.local.
* If a pod has a service, it also gets <pod-ip-address>.<service>.<ns>.svc.cluster.local.
* It’s also possible to assign DNS hostnames to pods, [using headless services](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-hostname-and-subdomain-fields). 

DNS Lookups
For performing lookups, pods are loaded with a set of search domains in their `/etc/resolv.conf`. The search domains start with `<ns>.svc.cluster.local`, and become more general. These can be further [configured.](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-dns-config)

This makes client configuration a lot simpler. To refer to a local service, a pod can just use the service name like `curl https://service:8443/stats`, and the search domain will take care of fully resolving it. Similarly, you can do this with cross namespace requests, like `curl https://service.another-namespace:8443/stats`.

A drawback of the search domain process is it can generate a substantial amount of redundant DNS traffic, as search domains get looked up in parallel in glibc. This was a huge problem in early Kubernetes clusters - which has been mostly solved by using node local DNS caching.

Still, to get the most performance, you can disable search domain lookups via a pod’s dnsPolicy field. Another trick is to use FQDNs (a domain name ending in a full stop, like `www.example.com`. ) which will bypass any search domains.

#### CNI
See [network-plugins](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/.)

CNI is the backend interface for implementing network configuration. It is required for any Kubernetes cluster, although typically its configuration will be hidden from you in managed clusters.

Its main task is deciding how to assign IPs to pods. There are many different implementations. Generally, if deploying to the cloud, it’s best to use the cloud CNI implementation, like aws-cni. This will give pods cloud native IPs that can be routed directly. For on premise clusters, generally CNI will be handled by an overlay network, such as calico - which will usually perform some sort of NAT/tunneling to traverse nodes.

It also does other things, like implementing NetworkPolicies.

#### Service Proxier
Kubernetes Services are implemented using a proxier. By default. this is kube-proxy which provides an iptables based implementation. (It also has an IPVS implementation, but it is largely unmaintained/deprecated at this point - with [some known issues](https://github.com/kubernetes/kubernetes/issues/72236)).

There are alternative implementations, such as kube-router and cilium, which can do much more sophisticated load balancing between pods, and a myriad of other features.

Generally these proxiers are all operating at the IP layer.

#### Service Meshes
Service meshes generally operate at the application layer (http, grpc, etc.), and generally supplant Services. They are purely optional, but many companies have come to rely on them to provide extended features beyond what vanilla Services can do. Examples are [linkerd](https://linkerd.io/) and [istio](https://istio.io/).

Service meshes usually provide:
* Built in application request tracing.
* Layer 7 load balancing/smart load balancing, which is request based rather than connection based.
* Sophisticated observability dashboards.
* Circuit breakers and other layer 7 features.

The main drawback is they generally do this by introducing one or two reverse proxies (like [envoy](https://www.envoyproxy.io/)) into every single hop to provide the sophisticated feature set. This has some downsides: increased latency and the proxies themselves become sources of failure.

#### External Load Balancers
For ingress, we generally need to configure external load balancers which will front our in-cluster ingress reverse-proxies. These will have an external IP (or corporate-wide private IP), and is what clients will contact for external traffic.

We generally have a choice between two types of load balancers: network load balancers (IP layer), and application load balancers (HTTP layer).

##### Network Load Balancers
Pros:
* Significantly better scalability over application load balancers, as these are usually implemented at the network/routing layer rather than with proxies.

Cons:
* Generally won’t terminate TLS, so you have to terminate TLS in your cluster (1).
* Connections are pass through and not pooled - so if you have 1 million external connections, your cluster proxies need to handle 1 million connections too.

1: AWS NLBs do support terminating TLS through some magic, but it costs significantly more and supports far less simultaneous connections. It is likely this is just routing through something like an ALB. See https://aws.amazon.com/elasticloadbalancing/pricing/?nc=sn&loc=3. An LCU on an NLB provides 100k active TCP connections. With a TLS listener, that drops to 3000! active connections.
