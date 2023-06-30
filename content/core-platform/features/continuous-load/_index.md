+++
title = " "
date = 2022-01-15T11:50:10+02:00
weight = 15
chapter = false
pre = "<b>Continous Load</b>"
+++

# Continuous Load

Continuous load gives you a black box mechanism of testing an environment e2e in a very low cost to implement way, giving you the ability as a platform team to have visibility of client level failures hopefully before your clients do.  You can also easily through CL ensure that all your points of connectivity are working and have precise visibility on failures at any level (direct ingress, Cloud Provide LB,  External CDN Layer etc).
Ultimately it’s a quick way of giving your platform an e2e availability score, with some really good built in root causing.

Features to define:

* Metrics, Dashboarding and Alerting: Metric collection to show historic performance of all the tested network paths
* Pod to pod validation: Continuously test pod to pod communication, by-passing service VIPs
* Pod to service validation: Continuously test pod to service communication via the service VIP
* Ingress validation: Continuously test traffic though external ingress mechanism
* Single node validation: Continuously test pod to pod communication for pods on the same node
* Inter-AZ validation: Continuously test pod to pod communication across AZs
* Inter-region validation