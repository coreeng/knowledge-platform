+++
title = "Ingress Custom Domain"
date = 2023-07-04T09:00:00+00:00
weight = 8
chapter = false
pre = "<b></b>"
+++

## Motivation

Allow tenants to onboard a service with its own custom domain onto the platform,
so that customers can continue to refer to the same URL.    

## Requirements


## Additional Information

Key Questions for Defuzz:

* Who manages the custom domain - platform or tenant? 
  * It can be tenant initially with a view to move the ownership to the platform

* Who manages the SSL certs rotation - platform or tenant?

* How can platform team ensure that a broken SSL certs for a custom domain doesn't affect other tenants?

* Assuming tenants manage their certs, can the platform provide an easy way for tenants to be alerted on certs expiry?

* How do platform engineer test this?

