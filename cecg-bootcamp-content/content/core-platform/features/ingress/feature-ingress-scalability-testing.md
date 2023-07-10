+++
title = "Ingress Scalability Testing"
date = 2022-01-15T11:50:10+02:00
weight = 10
chapter = false
pre = "<b></b>"
+++

## Motivation
Have confidence in the Ingress solution can scale beyond current workloads. Where scalability can be:
* Single connection throughput
* Number of connections
* Number of distinct Ingress Resources
* Frequency of Ingress Resource updates.

## Requirements
Regular ingress scalability tests covering X, Y, & Z. (picked during defuzz)

## Additional Information
Key questions for Defuzz:
* What scalability concerns are there? Which should be tested?
* How will the test run?
* How will it interact with autoscaling of the ingress controller
* What external LB is used and how does its scaling affect the results?
* What level does the test hit? Cross region LB? Single region? 
* Where does the load come from? Does it cause any bottlebecks? 

What will be delivered? Automated regular Ingress scalability report for the platform.



