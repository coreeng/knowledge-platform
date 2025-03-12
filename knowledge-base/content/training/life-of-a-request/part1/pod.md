+++
title = "What is inside a Pod?"
weight = 7
chapter = false
+++

#### Overview

The request had a very long journey so far, it managed to get into the Kubrenetes cluster and was routed by Ingress and Service into the **web-app** Pod.

The **web-app** Pod consists of 2 containers: **nginx** on port **80** and **web-app** on port **8080**.

![ Inside the Pod](/images/loar/1-15.png)
_Figure 1-15. Inside the Pod_

NGINX is commonly used as a reverse proxy, which is a server that sits between client devices (like web browsers) and the backend servers. It is also used to serve static content like images and files. A reverse proxy can help with load balancing, security, and caching, making it a popular choice for serving high-traffic websites.

See [Appendix L](../../appendix/l) for more details about the Nginx configuration.

Backend web application is written in Go, it generates basic HTML layout and provides REST API endpoints.

Below is the **deployment.yaml** showing the containers and their ports:

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web-app
  namespace: frontend
  labels:
    app: web-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: web-app
  template:
    metadata:
      labels:
        app: web-app
    spec:
      containers:
      - name: nginx
        image: nginx:latest
        ports:
        - containerPort: 80
      - name: web-app
        image: ghcr.io/lameaux/mox:latest
        ports:
        - containerPort: 8080
```

You can investigate the internals of the pods using **kubectl** command:

```
$ kubectl -n frontend get pods

NAME                      READY   STATUS    RESTARTS   AGE
web-app-7d759ffd6c-thnsr   2/2     Running   0          2m6s
web-app-7d759ffd6c-wpkzr   2/2     Running   0          3m
```

If the pod is running, but the application does not work as expected you can execute commands inside the container to debug the issues.

Usually you would like to start an interactive shell in a container, so that you can run commands:

```
$ kubectl exec -n frontend -it web-app-7d759ffd6c-thnsr -- /bin/sh

Defaulted container "nginx" out of: nginx, web-app
#
```

In case of multiple containers inside the pod, you need to specify the container explicitly

```
$ kubectl exec -n frontend -it web-app-7d759ffd6c-thnsr -c web-app -- /bin/sh

#
```

Your container may be built using a minimal Docker image (e.g. **distroless**) that contains only the application and its runtime dependencies, without a package manager, shell, or other debugging tools. Since these images are designed to be as small and secure as possible, they don’t include debugging utilities like bash, curl, or strace.

Distroless has a dedicated “debug” variant for each image, such as **gcr.io/distroless/base:debug**. These images include a minimal shell (busybox) and some basic utilities.

Another approach is to use a sidecar container. You can run a full-featured image (like Ubuntu or Alpine) alongside your distroless container. This sidecar container can have all the tools you need to debug the main container.

#### Troubleshooting

##### Check ports

Run **nc** inside a **web-app** container to check if port 8080 accepts connections: 

```
$ nc -zv localhost 8080

localhost [127.0.0.1] 8080 (?) open
```

##### Port Forwarding

You can also use **kubectl port-forward** to forward a local port to a Kubernetes Service, which can help you verify if pods are accepting connections on a specific port:

```
$ kubectl -n frontend port-forward service/frontend 8080:8080

Forwarding from 127.0.0.1:8080 -> 8080
Forwarding from [::1]:8080 -> 8080
Handling connection for 8080
```

This will map the container port to a local machine port, which you can then test.

```
$ curl localhost:8080/hello

404 page not found
```

##### Check open connections

Running **netstat** inside a **web-app** container displays a wide range of network and interface statistics.

**netstat -l** will show you all the sockets that are currently open.

In this case, all pod ports 80 (http nginx), 8080 (http web-app) and 9090 (metrics) appear to be active:

```
$ netstat -l

Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State
tcp        0      0 0.0.0.0:http            0.0.0.0:*               LISTEN
tcp        0      0 :::http-alt             :::*                    LISTEN
tcp        0      0 :::9090                 :::*                    LISTEN
tcp        0      0 :::http                 :::*                    LISTEN
```

##### Scan open ports

Nmap (Network Mapper) is a powerful and versatile open-source tool used for network discovery and security auditing. It allows you to scan and analyse networks by identifying hosts, services, open ports, and operating systems on a network.

```
$ nmap localhost

Starting Nmap 7.93 ( https://nmap.org ) at 2024-10-15 10:00 UTC
Nmap scan report for localhost (127.0.0.1)
Host is up (0.0000020s latency).
Other addresses for localhost (not scanned): ::1
Not shown: 997 closed tcp ports (reset)
PORT     STATE SERVICE
80/tcp   open  http
8080/tcp open  http-proxy
9090/tcp open  zeus-admin

Nmap done: 1 IP address (1 host up) scanned in 0.06 seconds
```

Nmap done: 1 IP address (1 host up) scanned in 0.06 seconds

Although we are running the tool from one of the containers, we actually see the ports that are open in the pod: **80** is for **nginx**, **8080** is for **web-app**.

See [Appendix J](../../appendix/j) for more details about **nmap**.
