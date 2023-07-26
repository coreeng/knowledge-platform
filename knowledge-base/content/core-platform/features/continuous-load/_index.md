+++
title = " "
date = 2022-01-15T11:50:10+02:00
weight = 15
chapter = false
pre = "<b>Continous Load</b>"
+++

# Continuous Load

Continuous load gives you a black box mechanism of testing an environment e2e in a very low cost to implement way, giving you the ability as a platform team to have visibility of client level failures hopefully before your clients do.  You can also easily, through CL, ensure that all your points of connectivity are working and have precise visibility on failures at any level (direct ingress, Cloud Provide LB,  External CDN Layer etc).
Ultimately it’s a quick way of giving your platform an e2e availability score, with some really good built in root causing.

## Motivation

* Gain confidence of any k8s setup with synthetic traffic, dashboards and alerts

## Additional Information

Helm charts for [continuous-load](https://github.com/coreeng/continuous-load)

