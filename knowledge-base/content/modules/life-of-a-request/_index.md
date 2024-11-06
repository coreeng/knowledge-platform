+++
title = "Networking: Life of a request"
date = 2024-11-06T12:13:10+02:00
weight = 7
chapter = false
+++

The goal of this learning module is to provide the engineers with the knowledge and skills needed to be able to solve networking issues with Kubernetes. We will demonstrate how a HTTP request finds its way into the application. We will identify potential networking issues and discuss how to debug them.

The request starts its journey from a web browser (or curl). It will get into the Kubernetes cluster where it will be handled by the application. The application will make a downstream call to another application inside the cluster which in turn will make a call to a third-party service outside the cluster.

![The journey of the request](/images/loar/0-1.png)
_Figure 0-1. The journey of the request_

This module is split into 3 sections, each describing a part of the journey:

* [Part 1 - Sending a request from the outside](part1)

  A web browser is sending a request to the Web App inside the Kubernetes cluster.

* [Part 2 - Downstream call inside the cluster](part2)

  The Web App is calling the User API (inside the cluster) to get user details.

* [Part 3 - Calling an external third-party service](part3)

* User API is calling a 3rd party service outside the cluster to validate the user token.

There is also the [Appendix](appendix) section containing extended details about specific topics that were intentionally skipped during the journey description to keep the reader focused.

#### Reference system architecture {#reference}

Now let’s take a deeper look into the architecture of the system that will handle the request.

![Reference system architecture](/images/loar/0-2.png)

_Figure 0-2. Reference system architecture_

We will focus mainly on **Kubernetes**, leaving Cloud Provider implementations out of scope.

The request will go through the Cloud Provider Network infrastructure (e.g. Cloud Load Balancer) and will be forwarded into the Kubernetes cluster where it will be routed according to the Ingress configuration.

The Ingress capabilities are implemented by Ingress Controllers, like Traefik of NGINX. The incoming traffic is balanced by several instances of Ingress Controller.

Our applications are stateless, they are deployed as multiple Pods behind a Service.

Web App is a backend application that serves HTML pages and provides REST API endpoints. It is deployed into the **frontend** namespace. The application exposes an HTTP endpoint which we want to make accessible from the outside world. To generate a response it needs to retrieve user details from the User API which is located in another namespace.

User API is a backend application that exposes a REST API endpoint returning a list of users. It is deployed into the **backend** namespace. The application exposes an HTTP endpoint which we don’t want to be accessible from the outside world (internet). The application returns user data based on the provided user token. The token needs to be validated against the third-party OAuth Service that is deployed outside of the cluster.

To get more details about the deployment check the [Appendix A](appendix/#a) section. 


