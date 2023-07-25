+++
title = "Uptime Dashboard"
date = 2023-07-24T14:56:10+01:00
weight = 3
chapter = false
pre = "<b></b>"
+++

## Motivation

* To have confidence in the kubernetes setup, a dashboard can show the health of deployed synthetic load 

* Visualise a history of activity so that patterns that form may be discovered and action may be taken pre-emptively

* Understand availability over differing time periods to ensure that SLOs are able to be met

## Requirements

* Dashboard that shows Synthetic Load Health

* Metrics To Display:
  * Latency
  * Response Codes (to measure availability) 


## Additional Information

Helm charts for [continuous-load](https://github.com/coreeng/continuous-load)