+++
title = "What is Service Discovery?"
weight = 2
chapter = false
+++

#### Overview

In part 2 we are touching the DNS area again as it plays an important role in Kubernetes Service Discovery.

Kubernetes automatically assigns domain names to services so they can be reached by name instead of by IP address, which may change.

Each service gets a Fully Qualified Domain Name (FQDN) in the following form:
```
<service-name>.<namespace>.svc.cluster.local
```

Where:

- **service-name** is the name of the service.
- **namespace** is the Kubernetes namespace in which the service is running.
- **svc** indicates that this DNS name is specifically for a service resource within Kubernetes.
- **cluster.local** is the default Kubernetes cluster domain, representing the local DNS zone within the cluster.

In our case the URL for the **user-api** service will be:

```
http://user-api.backend.svc.cluster.local:8080
```

Kubernetes supports several formats for referring to services, depending on the context and level of specificity required. For example:

- **service-name** (e.g., **user-api**): The simplest reference. Kubernetes will resolve this name to the service in the same namespace as the requesting pod.
- **service-name.namespace** (e.g., **user-api.backend**): Used to explicitly reference a service in a different namespace.
- **service-name.namespace.svc** (e.g., **user-api.backend.svc**): Useful for distinguishing between services and other potential resource types.
- **service-name.namespace.svc.cluster.local**: The full domain, specifying both the service and its location within the cluster’s internal DNS.

The DNS addon in Kubernetes (usually kube-dns or CoreDNS) keeps track of these records, enabling pods to discover each other by name, making networking more resilient to changes.

This makes it easier to update services without having to reconfigure applications that depend on them. Pods can communicate with other services using their DNS names, which are stable even if the ClusterIP changes.

#### Resolver Configuration

If we get inside the **web-app** pod running in the **frontend** namespace, and look at the **/etc/resolv.conf** file on that pod, we will see:

```shell
$ cat /etc/resolv.conf  

nameserver 10.96.0.10  
search web-app.svc.cluster.local svc.cluster.local cluster.local  
options ndots:5
```

**nameserver** defines a DNS server that the system will query for name resolution.

**search** defines a list of domain names to append to a hostname during resolution. This is useful for resolving short, unqualified hostnames. If you try to resolve a hostname that doesn’t end with  a dot (.), the resolver will append each of the listed domains and attempt to resolve them.

We have the **user-api** service running in the **backend** namespace. If we try to resolve a non-FQDN like **user-api.backend**, we will get the response, because the DNS resolver will search for the following options:

- user-api.backend.web-app.svc.cluster.local.
- **user-api.backend.svc.cluster.local.**
- user-api.backend.cluster.local.
- user-api.backend.

The second variant `user-api.backend.svc.cluster.local.` will match FQDN.

**options ndots** specifies the minimum number of dots in a name for it to be considered an FQDN. If a name has fewer dots, the resolver will append the domains listed in search.

The default **ndots** value in many Kubernetes clusters is set to **5**. This means that any hostname with fewer than five dots will be treated as a non-FQDN, and the DNS resolver will append the configured search paths to resolve it.

In Kubernetes, service names are often used without full qualifications, such as **user-api** instead of **user-api.backend.svc.cluster.local**. The **ndots** setting can affect how quickly or accurately services are discovered within the cluster.


#### ExternalDNS

[ExternalDNS](https://github.com/kubernetes-sigs/external-dns) is a Kubernetes add-on that automates the management of DNS records for Kubernetes services. It allows you to dynamically create DNS records for services or Ingress resources in Kubernetes, making it easier to manage external access to applications running in your cluster.

![Example of external-dns annotation for Ingress](/images/loar/2-2.png)
_Figure 2-2. Example of external-dns annotation for Ingress_

ExternalDNS watches your Kubernetes cluster for changes to certain type of resources, e.g.: Ingresses and Services (of **type=LoadBalancer**). When an Ingress or Service is created, updated, or deleted, ExternalDNS uses the respective DNS cloud provider’s API to create, update, or remove DNS records.

ExternalDNS can be configured using annotations on Kubernetes resources to specify how DNS records should be managed.   
In the below example **external-dns.alpha.kubernetes.io/hostname** specifies the hostname for a given service:

```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/hostname: myapp.cecg.io
    external-dns.alpha.kubernetes.io/target: cecg.io
  name: cecg-training-platform
spec:
  ingressClassName: traefik
  rules:
    - host: myapp.cecg.io
      http:
        paths:
          - backend:
              service:
                name: myapp
                port:
                  number: 80
            path: /
            pathType: ImplementationSpecific
```

ExternalDNS will request a DNS provider (e.g. Cloud DNS) to create a record for this domain.

![Example of a type A record created in CloudDNS for myapp.cecg.io](/images/loar/2-3.png)
_Figure 2-3. Example of a type A record created in CloudDNS for myapp.cecg.io_

ExternalDNS simplifies the management of DNS records for Kubernetes services, automating the creation and updating of DNS entries as your infrastructure evolves. It’s especially useful for dynamic environments where services and their endpoints may change frequently.

#### Troubleshooting

##### Install troubleshooting pod

The following command creates a pod in the **default** namespace.

```
$ kubectl apply -f https://k8s.io/examples/admin/dns/dnsutils.yaml

pod/dnsutils created
```

Once that Pod is running, you can exec **nslookup** in that environment.

```
$ kubectl get pods dnsutils  

NAME       READY   STATUS    RESTARTS   AGE  
dnsutils   1/1     Running   0          2m56s
```

If you see something like the following, DNS is working correctly.

```
$ kubectl exec -i -t dnsutils -- nslookup kubernetes.default

Server:	10.96.0.10  
Address:	10.96.0.10#53

Name:	kubernetes.default.svc.cluster.local  
Address: 10.96.0.1
```

If the **nslookup** command fails, we need to investigate deeper.

##### Check local DNS configuration

Take a look inside the **/etc/resolv.conf** file.

```
$ kubectl exec -ti dnsutils -- cat /etc/resolv.conf

nameserver 10.96.0.10  
search frontend.svc.cluster.local svc.cluster.local cluster.local eu-central-1.compute.internal  
options ndots:5
```

Verify that the search path and name server are set up like the following (note that search path may vary for different cloud providers).

Try resolving the host with **nslookup**:

```
$ kubectl exec -ti dnsutils -- nslookup user-api.backend

Server:	10.96.0.10  
Address:	10.96.0.10#53

Name:	user-api.backend.svc.cluster.local  
Address: 10.99.47.181
```

Errors such as the following indicate a problem with the CoreDNS (or kube-dns) add-on or with associated Services

```
$ kubectl exec -ti dnsutils -- nslookup user-api.backend  

Server:	10.96.0.10  
Address:	10.96.0.10#53

*** Can't find user-api.backend.svc.cluster.local: No answer
```

##### Check if the DNS pod is running

Use the kubectl get pods command to verify that the DNS pod is running.

```
$ kubectl get pods --namespace=kube-system -l k8s-app=kube-dns 

NAME                       READY   STATUS    RESTARTS      AGE  
coredns-6f6b679f8f-2qmml   1/1     Running   2 (17h ago)   27h  
coredns-6f6b679f8f-ljn7d   1/1     Running   2 (17h ago)   27h
```

The value for label **k8s-app** is **kube-dns** for both CoreDNS and kube-dns deployments.

##### Check for errors in the DNS pod

Use the **kubectl logs** command to see logs for the DNS containers.

```
$ kubectl logs --namespace=kube-system -l k8s-app=kube-dns
```

See if there are any suspicious or unexpected messages in the logs.

#### Further Reading

- [Customising DNS Service | Kubernetes](https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/)
- [Debugging DNS Resolution | Kubernetes](https://kubernetes.io/docs/tasks/administer-cluster/dns-debugging-resolution/) 
