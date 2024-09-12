+++
title = "Deployed Load Testing"
weight = 5
chapter = false
+++

## Motivation

* Learn how NFT load generation is often harder than scaling your application
* Learn the value of capturing load test metrics in the monitoring stack

## Requirements

* Ability to run k6 in a distributed way in Kubernetes, validated in your local kubernetes setup
* Deployed load testing running in the P2P 
* All stats from load generation available in Grafana

## Additional Information


Use the following the [guide for running k6 on Kubernetes](https://k6.io/blog/running-distributed-tests-on-k8s/). 
However, rather than using their test script and yaml files, use the one in the nft/ramp-up folder in the reference applications.

You can find Grafana dashboard for load testing stats in the nft/dashboards folder in the reference applications

## Questions / Defuzz / Decisions
...

## Deliverables 

* ... 
