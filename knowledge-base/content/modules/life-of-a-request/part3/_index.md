+++
title = "Part 3 - Calling an external third-party service"
weight = 3
chapter = false
+++

The request from the **web-app** successfully reached the **user-api** pod. In order to generate a response, the **user-api** pod needs to make a call to the external third-party OAuth Service to get the ID of the user.

**Note:** We don’t really have a real OAuth Service for the purpose of this learning module, so we will be using a mock service [https://httpbin.org/uuid](https://httpbin.org/uuid) to generate UUIDs.

![User API is calling external OAuth Service](/images/loar/3-1.png)
Figure 3-1. User API is calling external OAuth Service

### WIP
