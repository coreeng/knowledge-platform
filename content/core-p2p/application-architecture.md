+++
title = "Application Architecture"
date = 2022-12-27T06:32:10+02:00
weight = 5
chapter = false
+++

## Multiple Instances + Stateless Architecture 
Highly available applications require multiple instances, multiple instances are made more easily via having a stateless application architecture along with load balancing in front of the application instances.

### What is stateless application architecture?
Any state that is required beyond a single request is kept out side of the application e.g. in database or event log. 

Requests from the same user can go to different instances of an application. Allowing instances to be replaced without impact to the user.

## Graceful Shutdown
Common problem: Containers Ignore SIGTERM. Then Kubernetes times out and sends a SIGKILL. 

1. Container receives SIGTERM
1. Readiness probe to fail 
1. Stop accepting new connections
1. For existing connections, finish off existing requests, but don’t accept any new ones
1. Close any database connection pools 
1. Shut down
