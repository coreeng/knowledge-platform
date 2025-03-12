+++
title = "What is an Ingress?"
weight = 5
chapter = false
+++

#### Overview

The journey of the request goes on. We successfully established a secure HTTPS connection with the target host and sent an HTTP request body. According to the reference architecture (see [Reference system architecture](../../#reference)) the entrypoint into our Kubernetes cluster is represented by Ingress.

Kubernetes Ingress is a native API resource that manages external access to services within a Kubernetes cluster, typically for HTTP and HTTPS traffic. It acts as an entry point for external users to access the applications running inside the cluster by defining rules for routing HTTP/HTTPS requests to the appropriate services based on path, hostname, or other conditions.

The Ingress resource is not a standalone solution; it requires an Ingress Controller (such as NGINX, Traefik, or HAProxy) to be installed in the cluster. The controller translates the Ingress rules into configurations that define how the network traffic is handled.

An Ingress resource is associated with a set of rules to manage external traffic. The Ingress controller interprets these rules and ensures that the correct routing is configured. It typically performs tasks like:

- Routing requests based on hostnames, paths, or other criteria to specific services.
- TLS/SSL termination to secure traffic.
- Load balancing for services that have multiple pods.

![Ingress terminates TLS and routes HTTP traffic](/images/loar/1-14.png)
_Figure 1-14. Ingress terminates TLS and routes HTTP traffic_

#### Troubleshooting

##### Check Ingress Resource

List all Ingress resources and check their status:

```
$ kubectl get ingress -n frontend

NAMESPACE	        NAME		CLASS		    HOSTS	ADDRESS PORTS   AGE
web-app	frontend	ingress	    myapp.cecg.io			        80      1h
```

View the detailed configuration:

```
$ kubectl describe ingress frontend -n frontend

Name:             web-app
Namespace:        frontend
Address:
Ingress Class:    ingress
Default backend:  <default>
Rules:
  Host		Path  Backends
  ----		----  --------
  myapp.cecg.io
			/   	web-app:80 ()

Annotations:	
external-dns.alpha.kubernetes.io/hostname: myapp.cecg.io
external-dns.alpha.kubernetes.io/target: cecg.io
Events:	<none>
```

Ensure the hostname, paths, backend services, and annotations are correct.

##### Verify Ingress Controller

Ingress requires a controller to function, such as NGINX Ingress Controller, HAProxy, or Traefik.

Check if the Ingress controller pods are running:

```
$ kubectl get pod -n traefik

NAME                            READY   STATUS    RESTARTS   AGE
traefik-746768f765-2f825        1/1     Running   0          1h
traefik-746768f765-7c9x7        1/1     Running   0          1h
traefik-746768f765-lbbzs        1/1     Running   0          1h
```

Inspect the logs of the Ingress controller for errors using `kubectl logs -n <ingress-namespace> <controller-pod-name>`.

Ensure that the controller matches the version or type of ingress resource you’re using (e.g., NGINX annotations won’t work with Traefik).

#####  Examine Ingress Annotations

Different ingress controllers support different annotations. Check the annotations carefully.

For NGINX Ingress, you might need to adjust annotations such as **nginx.ingress.kubernetes.io/rewrite-target, nginx.ingress.kubernetes.io/ssl-redirect**, etc.

Run `kubectl describe ingress <ingress-name>` to confirm if the right annotations are applied.

#####  TLS Troubleshooting

If TLS termination is involved, check the certificate configuration.

Run `kubectl describe secret -n <ingress-namespace>` to check if the TLS secret is correctly configured.

##### Check for Errors in Logs

Inspect the logs of your Ingress controller (e.g., NGINX Ingress):

```
kubectl logs -n <ingress-namespace> <ingress-controller-pod-name>
```

Look for connection issues, 404 or 502 errors, TLS errors, or backend service connection failures.
