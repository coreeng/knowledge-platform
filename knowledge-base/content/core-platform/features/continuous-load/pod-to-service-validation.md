+++
title = "Pod to Service Validation"
date = 2023-07-24T14:56:10+01:00
weight = 3
chapter = false
pre = "<b></b>"
+++

## Motivation

* Have confidence in the kubernetes setup by having a continuously running service that is under a continuous load

* Be aware of any problems before clients are

## Requirements

* To deploy a test service and place under synthetic load that tests the availability of that service

* To collect metrics on the health of the deployed service under synthetic load 

* The following should be configurable:
  * Throughput - Min 3TPS, so we have enough datapoints per hour to calculate all the percentiles
  * Additional endpoints to be hit by the injector - ingress, edge, but should always hit the target
  * Thresholds and latency

## Additional Information

Helm charts for [continuous-load](https://github.com/coreeng/continuous-load)