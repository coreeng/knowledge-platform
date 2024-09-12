+++
title = "Resiliency and Capacity Testing"
weight = 6
chapter = false
+++

## Motivation

* Consider resiliency as a first class citizen
* Methodically set resource limits based on evidence (NFT)
* Understand the breaking point of your application
* Know the capacity of a single pod with those resource limits
* Know the horizontal scalability characteristics

## Requirements

* Test how `/hello` endpoint behaves under a load:
    * Set resource quota for the application namespace
    * Tweak resource limits for the reference application to support 500 virtual users 95th Percentile is < 50ms with a single replica
    * Understand the breaking point of the application with these resource limits. Requests start failing or latency rising. 
    * Horizontally scale to reach 1500 virtual users
* Handle slow downstream dependencies
    * Respond within 500ms even if the downstream is slow or down
    * You should return a 503 service unavailable status code 

## Additional Information

K6 load test runner and Prometheus may consume significant amount of memory during load testing, 
so you may need to increase resource limits for your local kubernetes cluster.

## Questions / Defuzz / Decisions
...

## Deliverables (For Epic)

* ...