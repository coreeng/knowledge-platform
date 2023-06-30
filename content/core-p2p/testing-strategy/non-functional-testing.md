+++
title = "Non Functional Testing"
date = 2022-12-27T06:32:10+02:00
weight = 3
chapter = false
+++

The goal of non-functional testing is to:

* **Performance:** Validate your service meets the required performance characteristics
* **Breaking point:** Understand the breaking point and how the system behaves i.e. what if we get an expected peak? 
* **Resiliency:** understand how the system behaves under failure scenarios
* **Reliability:** validate the system continues to behave over longer periods

With that information we can:

* Set resource limits 
* Understand vertical and horizontal scaling characteristics of the service
* No its architectural limit i.e. at what point do we need to fundamentally change the service to meet a new level of load

 

## Types of NFT

### Load Testing
Load testing is a type of Performance Testing used to determine a system's behavior under both normal and peak conditions.

Load Testing is used to ensure that the application performs satisfactorily when many users access it at the same time.

### Stress Testing
Stress Testing is a type of load testing used to determine the limits of the system. The purpose of this test is to verify the stability and reliability of the system under extreme conditions.

Whereas basic load testing will verify the system under expected peak loads, stress testing will find the breaking point so we know how much free capacity we have before needing to scale or re-architect a service.

### Soak Testing
While load testing is primarily concerned with performance assessment, and stress testing is concerned with system stability under extreme conditions, soak testing is concerned with reliability over a longer period of time.

## Integrated vs stubbed NFT

At many organisations scale non-functional testing takes place in large integrated environments shared by many teams.
A typical workflow may be:
* Team A tests
