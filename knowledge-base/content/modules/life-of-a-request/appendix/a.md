+++
title = "A - Exploring the cluster with kubectl"
weight = 1
chapter = false
+++

Kubernetes provides a command line tool for communicating with a Kubernetes cluster's control plane, using the Kubernetes API. This tool is named **kubectl**. See [https://kubernetes.io/docs/reference/kubectl/](https://kubernetes.io/docs/reference/kubectl/)

Below is the output of **kubectl** that demonstrates pods and a service deployed to the **frontend** and **backend** namespaces.

```
$ kubectl get pod -n frontend

NAME                        READY   STATUS    RESTARTS   AGE  
web-app-657b784bd-f7wfm     1/1     Running   0          1h  
web-app-657b784bd-zr9pp     1/1     Running   0          1h
```

```
$ kubectl get service -n frontend 

NAME        TYPE        CLUSTER-IP        EXTERNAL-IP   PORT(S)     AGE  
web-app     ClusterIP   192.168.193.215   <none>        8080/TCP    1h
```

```
$ kubectl get pod -n backend  

NAME                        READY   STATUS    RESTARTS   AGE  
user-api-788999d99-nk4f9    1/1     Running   0          1h  
user-api-788999d99-rtmkm    1/1     Running   0          1h
```

```
$ kubectl get service -n backend 

NAME        TYPE        CLUSTER-IP        EXTERNAL-IP   PORT(S)   AGE  
user-api    ClusterIP   192.168.193.216   <none>        8080/TCP   1h
```
