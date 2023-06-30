+++
title = "Functional Testing"
date = 2022-12-27T06:32:10+02:00
weight = 2
chapter = false
+++

Functional testing is an overloaded term. Our full term is "Stubbed Functional Testing". We define it as:

* No manual actions: automated and automatically run. It is quite common to hear tests are "automated" but when digging deeper there are test scripts manually run. Functional tests should be automatically run as part of the P2P for every commit.
* Stubbed aka decoupled tests: All external dependencies are stubbed out so that the tests are independent. 

Ideally the tests can be run:

* Locally by a developer 
* On a built agent
* Against a deployed version of the application

We typically use [BDD](https://dannorth.net/introducing-bdd/) type tests for Functional Tests.
