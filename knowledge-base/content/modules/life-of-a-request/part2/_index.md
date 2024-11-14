+++
title = "Part 2 - Downstream call inside the cluster"
weight = 2
chapter = false
+++

The **web-app** successfully accepted the incoming request and started processing it. To generate a response it needs to get the data from the **user-api** application running inside the cluster.

![Downstream call inside the cluster](/images/loar/2-1.png) 
_Figure 2-1. Downstream call inside the cluster_

We are going to make a call to the User API using its ClusterIP Service domain name.

- [What is Service Discovery?](discovery)
- [What are Network Policies?](netpol)
