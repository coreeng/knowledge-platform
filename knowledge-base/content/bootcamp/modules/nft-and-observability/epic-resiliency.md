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

* Set resource limits for your application 
* Support 500 virtual users 95th Percentile is < 20ms with a single replica 
* Understand the breaking point of the application with these resource limits. Requests start failing or latency rising. 
* Horizontally scale to reach 1500 virtual users
* Fail gracefully if the database is slow
* Fail gracefully if the database is down

## Questions / Defuzz / Decisions
...

## Deliverables (For Epic)

* ...