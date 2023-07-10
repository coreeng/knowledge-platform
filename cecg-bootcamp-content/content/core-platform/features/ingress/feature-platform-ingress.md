+++
title = "Platform Ingress"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation

Typically there will be requirements that mandate a custom ingress solution such as:

* Scalability in number of Ingresses
* Segregated ingress controllers
* Path re-writing?
* More advanced routing
* Enriched access logs
* App level metrics: 
  * Success/Failure
  * Latency metrics (depends on Ingress Controller choice)
* Contribute to SLIs, SLOs, Error Budgets
* Zero downtime rolling updates
* Fine grained control over connection draining etc

## Requirements

* Scalable enough for MVP tenant load * 2
* Support segregated ingress
* Typical Requirements 
  * HTTP Ingress
  * TCP Ingress
  * UDP Ingress
  * Host based routing
  * Header based routing
  * Access logs
  * Per ingress metrics

## Questions

Is segregated ingress a requirement? Should it be designed in from day 1?

* Typically this is the case, even for External vs Internal Traffic

Which Ingress Controller should be used? Two we recommend:

* [Traefik](https://traefik.io/)
* [Contour](https://projectcontour.io/)
  * Specifically L7, so if there are use cases for TCP/UDP, need an alternative solution 

If the Ingress controller has any dashboards/metrics - how is access managed?

Can each tenant see each other's network traffic and metrics?

Does the ingress controller have support for ACME? Is it needed?

## Additional Information
Two parts to this:

* Internal ingress controller
* External to the cluster loadbalancers fronting the cluster. These can be a mix of internet facing and private LBs

As part of this feature we need select an Ingress Controller based on requirements. Cloud ingress controllers typically don’t work for more than a handful of exposed services.

What is delivered? Exposing of services using a shared set of ingress controllers

#### Examples

##### NLBs to Ingress Controller pods on AWS exaple


{{< figure src="/images/reference/platform/core-platform-aws-usecase-one.png" >}}

[1]

AWS [Network Load Balancers (NLBs)](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/introduction.html) with target pools that contain the nodes that ingress controller is running on. The NLB sends traffic directly to the Ingress Controller via a HostPort that is set on the deployment. 

The advantages of this set up are:
* **Scalability**: Network load balancers are implemented as part of the SDN fabric in AWS. They are network routers, essentially, rather than a collection of managed VM instances that proxy connections. While it is a black box to us, we can assume they are performing something akin to ECMP routing to load balance at the IP layer (layer 3). As a result, they can handle as much load as the networking fabric tolerates, rather than relying on a scale up of VMs. As a result, we can instantly handle very high loads (up to a million RPS in testing) instantly.
* **Ease of use**: As far as network load balancing goes, AWS NLBs are fairly easy to use. Feed pods receive the original client IP as-is, and the destination IP is already modified to match the target node’s IP. In traditional network load balancing, usually each node has to be aware of the virtual IP (VIP) itself, so it can set up routing to handle the routed traffic. AWS takes care of this for us and “magically” remaps the target IP to that of the node - and also remaps the source IP back to the VIP for return traffic.

The drawbacks are:
* **TLS termination**: TLS termination is handled by the feed component itself. This means we have additional complexity in managing the correct set of TLS protocols and settings to use, in addition to the computational overhead of having to terminate TLS. At large scale, this would become a severe bottleneck if we had to handle many TLS connections. This is largely mitigated because we use Akamai to terminate TLS connections for customers, so the actual number of connections we need to handle is fairly low in proportion to the number of customers.
* **Security**: It is generally considered better to use a managed service for TLS termination, due to various TLS schemes being vulnerable. Since we have to do it ourselves, we have to ensure we stay up to date on security related to TLS. Mozilla SSL Configuration Generator  can be used for this.
* **Scalability**: While the NLB provides insane scalability, we are limited in our set up to a single feed pod per node. So we can’t scale horizontally as much as would be possible otherwise - it is limited to the number of nodes in the cluster. So far, this hasn’t been an issue for us.
* **Availability**: Availability during rolling updates, node reboots, and ingress updates is much trickier in this set up, compared to if we were using a L7/L4 terminating proxy. We had to invest significant engineering effort to create a graceful termination solution that worked with haproxy+nginx. Nginx alone does not gracefully terminate and drops all connections immediately upon reload - which is not something that was acceptable to our use case. In essence, we have to implement connection draining ourselves which is usually something one gets for free from a managed proxying load balancer.

* Cross zone load balancing and client IP preservation are mutually exclusive with NLBs. If you enable them both, rare connection resets can occur.
* Hairpin traffic fails with internal NLBs unless client IP preservation is disabled.

#### GCP Platform Ingress Example

{{< figure src="/images/reference/platform/core-platform-gcp-usecase-one.png" >}}
[2]

This is a very similar set up to AWS. We use a Network Load Balancer with a virtual IP (VIP), which in GCP is implemented via [Maglev](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/44824.pdf) (the GCP software load balancer) and [Andromeda](https://www.usenix.org/system/files/conference/nsdi18/nsdi18-dalton.pdf) (the GCP networking fabric). The main difference is we rely on the Kubernetes `Service` with `type: LoadBalancer` to handle the complexities of VIP management for us, as this isn't handled by GCP itself (unlike AWS NLBs).

Each node adds the VIP to its local routing table. Packets arrive at each node’s NodePort with the original source IP/port, and destination IP as the VIP. Only the destination port is different, as it gets rewritten by the networking fabric to be the same as the NodePort. The kubernetes proxier (Cilium in our case) handles load balancing to one of the local pods in a form of DNAT - the destination IP is rewritten to the destination pod, and then translated back on return packets. From the node itself, everything is direct return back to the client IP directly, bypassing any network mangling.

TLS is terminated directly by Traefik, similar to the AWS set up.

The advantages of this approach are:
* **Scalability**: Similar to the AWS use case, network load balancers don’t need to scale up like a traditional load balancer, so the full capacity is immediately available as needed. This means we can hit very high requests per second, limited only by the number of backend Traefik pods. Unlike in AWS, we can also have multiple Traefik pods per node, so we have a much higher limit for horizontal scaling.
* **Ease of use**: Kubernetes takes care of all the complexities of managing the myriad of GCP resources needed to orchestrate NLBs, and also takes care of managing the local VIP routing. This means all we need to do is create the appropriate Service and almost all the work is done for us.
* **Availability**: Traefik supports hot reloads of ingress unlike nginx. This means we don’t need a lot of workarounds and effort put towards making config reloads non-disruptive.
* **Simplicity**: Traefik and Kubernetes does most of the infrastructure work for us, with little configuration required. The Traefik set up is very straightforward compared to our AWS set up. It also supports Ingress and the annotations we need out of the box (strip-path, etc.).
* **Performance**: In performance testing, the throughput and latency metrics are very similar to the AWS solution - so we haven’t lost anything in this area.

The drawbacks of this approach are:
* **Graceful termination**: Kubernetes Service with externalTrafficPolicy: local has known problems with graceful termination. There is KEP-1699 in progress to fix this by modifying the cloud provider to use the EndpointSlice to maintain the list of viable pods during termination - so terminating pods can receive new connections. As a workaround, we’re currently using externalTrafficPolicy: cluster. The drawbacks of this are we lose the client IP (it is rewritten via SNAT to be the node IP), and there is potentially an additional hop between nodes as packets arriving on the NodePort will be redirected to any traefik pod in the cluster. It’s been deemed acceptable in our use case, since it is similar to the AWS solution which also doesn’t preserve the client IP for internal traffic.
* **TLS Termination/Security**: Similar to the AWS approach, we are terminating TLS ourselves. It would be preferable to use a managed service for termination. However, we largely mitigate this by relying on Akamai to handle potentially hostile traffic - we only ever receive connections from trusted sources (either internal or Akamai’s own IPs).
* **Potential for Misbalancing**: Nodes all have equal weight when using the built in Kubernetes Service support. As a result, traefik pods may not receive equal amounts of traffic, since connections will be load balanced across nodes rather than pods. Given enough pods we expect this to not be a huge issue since Kubernetes will try to schedule them across nodes equally.






