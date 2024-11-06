+++
title = "What is a Service?"
weight = 6
chapter = false
+++

#### Overview

In Kubernetes, a Service is an abstraction that defines a logical set of Pods and a policy by which they can be accessed. The main purpose of a Service is to provide stable networking and load balancing for a group of Pods, which are ephemeral and can have changing IP addresses.

**Stable Networking**

Pods in Kubernetes have dynamic IPs that can change when they are created, restarted, or scaled. A Service provides a static IP (called a ClusterIP) that allows other services or users to access the Pods without worrying about their changing IP addresses.

**Load Balancing**

A Service can distribute incoming network traffic across multiple Pods. It acts as a built-in load balancer, ensuring that traffic is evenly spread among the available Pods in the service.

**Discovery**

Services in Kubernetes make Pods discoverable by assigning a DNS name to the Service. Other components can resolve this DNS name to access the Pods behind the Service.

**Decoupling**

By using a Service, Pods that need to communicate with each other are decoupled from the specific details of how they are deployed. Clients communicate with the Service rather than individual Pods, making the system more resilient and scalable.

**Exposing Pods**

Services can expose Pods internally within the cluster (using a ClusterIP) or externally to the internet (using a NodePort or LoadBalancer) depending on the service type.

#### ClusterIP

The ClusterIP is the default type of Service in Kubernetes. It exposes the service on an internal IP address, only accessible from within the Kubernetes cluster.

It’s ideal for internal communication between services within the same Kubernetes cluster, like when a front-end service needs to communicate with a back-end service.

It cannot be accessed directly from outside the cluster. However, within the cluster, services can be accessed using the internal ClusterIP address or through DNS.

Example:

```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: my-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: ClusterIP
```

This Service selects Pods with the label **app=my-app** and forwards requests on port **80** to port **8080** on those Pods.

#### NodePort

A NodePort service exposes the service on a static port on each worker node in the cluster. When a client sends a request to **\<NodeIP\>:\<NodePort\>**, Kubernetes routes the request to the appropriate Pod behind the service.

Useful for exposing a service to external traffic without the need for a cloud load balancer. This type of service can be used for testing or for non-production environments.

NodePort services expose the application on all nodes in the cluster, making it accessible via any node’s IP address, on the assigned port. This opens up the possibility of unintended external access. If not properly restricted, attackers can scan and discover the open NodePort, which may expose the application to external threats such as DDoS attacks or brute force attempts.

It can be accessed externally, but only by specifying the node’s IP address and the assigned port.

The port range for NodePort is typically between 30000 and 32767\.

Example:

```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: my-app
  type: NodePort
  ports:
    - port: 80
      targetPort: 8080
      nodePort: 30036  # Optional, if omitted, Kubernetes assigns a random port in the range

```

#### LoadBalancer

The LoadBalancer service type exposes the service externally using a cloud provider’s load balancer (e.g., AWS ELB, Google Cloud LB, Azure LB). When you define a Service of type LoadBalancer, Kubernetes makes API calls to the cloud provider to create an external load balancer. This external load balancer will route traffic to the service’s exposed IP and port.

Ideal for exposing a service to the internet. It is commonly used for production environments where external clients need to access a service.

By default, the cloud provider assigns a dynamic external IP to the service. If you want to assign a static IP to a LoadBalancer service, you need to reserve a static IP address in your cloud provider and configure your service to use it.

Only available when Kubernetes is running in a cloud environment (e.g., AWS, GCP, Azure) that supports this feature.

Example:

```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: my-app
  type: LoadBalancer
  ports:
    - port: 80
      targetPort: 8080
```

#### ExternalName

An ExternalName service maps a Kubernetes service to a DNS name external to the cluster. This type of service returns a CNAME record with the value of the externalName field.

Useful when you want to access an external service (like an external database or API) by assigning it a service name within the cluster. It does not route traffic through the Kubernetes network, but simply acts as a DNS alias.

This service type does not create a ClusterIP or any networking routes, and simply resolves to an external DNS name.

Example:

```
apiVersion: v1
kind: Service
metadata:
  name: my-external-service
spec:
  type: ExternalName
  externalName: external-service.cecg.io
```

#### Headless

A Headless Service is a variant of the ClusterIP service but without a stable IP address. Instead of load-balancing across the Pods, it returns the IP addresses of individual Pods. Clients can then perform service discovery and route traffic to individual Pods.

Headless services are useful for stateful applications where clients need to address individual Pods directly (like databases, message queues, etc.).

Clients will get the IP addresses of all the Pods behind the service instead of a single service IP.

Example:

```
apiVersion: v1
kind: Service
metadata:
  name: my-headless-service
spec:
  clusterIP: None  # Indicates headless service
  selector:
    app: my-app
  ports:
    - port: 80
```

#### Troubleshooting

If a Kubernetes Service is not functioning correctly, it often comes down to configuration issues, selector mismatches, or connectivity problems.

##### Check Service

Ensure the service is created and running properly:

```
$ kubectl get service -n frontend -o wide

NAME		TYPE 		CLUSTER-IP 	EXTERNAL-IP	PORT(S) 
AGE 	SELECTOR
web-app	ClusterIP  192.168.147.96  <none>        	8080/TCP   	1h	app.kubernetes.io/name=web-app
```

Look for the service type (ClusterIP, NodePort, LoadBalancer), the assigned IP, and the port mappings.

##### Verify Pods

Ensure the Service selector matches the labels on the Pods you want to expose. If the selector is incorrect, the Service will have no endpoints.

```
$ kubectl get pods -n frontend --show-labels

NAMESPACE                            NAME                                                       READY   STATUS      RESTARTS       AGE     LABELS
frontend                 web-app-545cf85c75-mmskx                                   1/1     Running     0              1h      app.kubernetes.io/name=web-app,pod-template-hash=545cf85c75
```

Confirm that the labels on your Pods match the **spec.selector** field in the Service.

Verify the Pods backing the service are running and healthy.

```
$ kubectl get pods -n frontend --selector=app.kubernetes.io/name=web-app

NAMESPACE                            NAME                                                       READY   STATUS      RESTARTS       AGE     LABELS
frontend                 web-app-545cf85c75-mmskx                                   1/1     Running     0              1h      app.kubernetes.io/name=web-app,pod-template-hash=545cf85c75
```

Ensure the Pods are in the Running state, and that the service selector matches the Pods’ labels.

##### Check Endpoints

Verify that the service has assigned endpoints, i.e., the IP addresses and ports of the Pods the service routes traffic to.

```
$ kubectl get endpoints web-app -n frontend

NAME       ENDPOINTS          AGE
web-app   	10.100.7.31:8080   1h
```

If the output shows no endpoints, the Service cannot find any matching Pods.

##### Validate Network Connectivity

If you’re unable to connect to the Service from within the cluster, you can test connectivity using a temporary Pod:

```
$ kubectl run mycurlpod --image=curlimages/curl -i --tty -- sh
```

Then, try to **curl** to **\<service-name\>.\<namespace\>.svc.cluster.local:\<port\>**:

```
$ curl web-app.frontend.svc.cluster.local:8080
```

##### DNS Issues

If a Service is not resolving correctly through DNS, you can test domain resolution using a temporary Pod:

```
$ kubectl run mydnspod --image=tutum/dnsutils -i --tty -- sh
```

Then try resolving **\<service-name\>.\<namespace\>.svc.cluster.local** with **nslookup**:

```
$ nslookup web-app.frontend.svc.cluster.local

Server:	192.168.0.10
Address:	192.168.0.10#53

Name:	web-app.frontend.svc.cluster.local
Address: 192.168.224.76
```

This helps verify that the DNS service is functioning correctly in the cluster.

Check the logs of Kube-DNS or CoreDNS or for errors:

```
$ kubectl logs kube-dns-6dd99549bd-6smjv -n kube-system
```

Restart the Kube-DNS or CoreDNS Pods:

```
$ kubectl rollout restart -n kube-system deployment/coredns
```

##### Check Network Policies

Network policies in Kubernetes can restrict traffic between Pods and Services. Ensure that there are no policies preventing traffic between Pods and Services.

```
$ kubectl get networkpolicies -n <namespace>
```

If a network policy exists, verify that it allows traffic to and from the correct Pods and namespaces.

##### Inspect Logs and Events

Check the logs and events for more detailed error messages:

```
$ kubectl logs <pod-name>
$ kubectl get events --sort-by='.metadata.creationTimestamp'
```
