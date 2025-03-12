+++
title = "What is NAT?"
weight = 3
chapter = false
+++

#### Overview

NAT (Network Address Translation) is a method used in networking to map multiple IP addresses to a single or a few IP addresses. In Kubernetes, NAT is often used to manage the traffic flow between pods (containers within Kubernetes) and external systems. Kubernetes also uses NAT to ensure that workloads inside a cluster can access external resources securely and efficiently.

Kubernetes uses SNAT (Source Network Address Translation) when pods connect to external IPs. This allows outgoing traffic to be translated from the pod’s IP to the node’s IP, making it routable to the internet.

**NAT Gateway**

In cloud environments, configuring NAT is usually done through your cloud provider’s NAT gateway options. For example:

- **AWS**: Use an **AWS NAT Gateway** in a public subnet, and route traffic from private subnets (where your Kubernetes nodes are) to this gateway.  
- **GCP**: Configure a **Google Cloud NAT** for your VPC, which handles NAT for private IP addresses.  
- **Azure**: Use an **Azure NAT Gateway** associated with your subnet to enable internet access.

![NAT Gateway](/images/loar/3-3.png)
_Figure 3-3. NAT Gateway_

**IP Whitelisting**

IP whitelisting is a security feature that allows only specified IP addresses to access certain resources or systems. It’s a way of enforcing access control by permitting only “trusted” IPs, usually by creating a list of these approved addresses in a network firewall or an application.

Administrators define a list of trusted IP addresses or ranges (the “whitelist”) that are allowed access. Any IP address not on this list is denied access.

When IP whitelisting is enabled on systems behind NAT, access is often based on the public IP address provided by NAT. Only the NAT device’s public IP is whitelisted.

#### Troubleshooting

##### Confirm IP whitelisting requirements

Determine the IPs or IP ranges allowed by the service you’re trying to access.

Verify if whitelisting is based on the public IP of the Kubernetes node or the IP range of the cluster (for cloud-based services).

##### Check outbound IP address

For testing, you can deploy a simple curl pod and attempt to reach an external IP or DNS address, confirming that outbound NAT is correctly configured.

```
$ kubectl run curl-test --image=busybox -it --rm --command -- curl https://api.ipify.org  

212.50.113.226
```

Compare this IP with the whitelisted IP list. If they don’t match, you may need to adjust your NAT setup or ask the third party administrator to update their IP list.

