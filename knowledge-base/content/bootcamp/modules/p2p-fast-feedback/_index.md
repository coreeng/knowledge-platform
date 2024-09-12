+++
title = "P2P Fast Feedback"
weight = 1
chapter = false 
+++

## Introduction

**Theme:** Build a service with a Path to Prod (P2P) that provides fast feedback to the engineer

**What should you {{< colour green understand >}} at the end of the module?**

* Consolidated Path to Production a.k.a Paved Paths
    * [Core Path to Production]({{% ref "/roadmaps/core-p2p/" %}})
* Microservice structure Continuous Delivery 
    * [Service Encapsulation & Trunk Based Development]({{% ref "/core-compass/sdlc/delivery-unit-encapsulation" %}})
    * [Testing Strategy]({{% ref "/knowledge/foundational-practices/testing" %}})
* Basic intro to Containers and Kubernetes

**What should you be {{< colour green "able to do" >}} by the end of the module?**

* Build a service in Java with the Spring Boot framework or with GoLang and the gin web framework
    * Test it: unit tests, stubbed functional tests, stubbed non-functional tests
        * Applying the knowledge in [Testing Strategy]({{% ref "/knowledge/foundational-practices/testing" %}})
            * Focus on stubbed testing
    * Package it with Docker
    * Deploy it to Kubernetes

## FAQs

#### Why do we start with working on a service in a platform engineering bootcamp?

1. As Platform Engineers our goal is to speed up product delivery by abstracting over low level infrastructure. This allows us to emphasize with application engineers. 
2. A common product we build as Platform Engineers is a [consolidated P2P]({{% ref "/roadmaps/core-p2p/" %}}) for delivery teams
3. As Platform Engineers we solve problems with software and services 
