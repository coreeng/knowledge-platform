+++
title = "Canary Deployments"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation

Testing out features or just new releases on a subset of production traffic before fully rolling out new releases.

Very high confidence of not affecting all customers if enough verification has been done on the canary.

Cheaper or more resource efficient than blue green.

## Requirements

P2P pipeline shape supporting canary deployments

Mechanisms agreed on and implemented for how a canary is promoted to a full production release e.g.

* Time
* Alerts
* Continuous testing

## Questions 

