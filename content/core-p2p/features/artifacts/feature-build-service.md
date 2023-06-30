+++
title = "Build Service"
date = 2022-01-15T11:50:10+02:00
weight = 1
chapter = false
pre = "<b></b>"
+++

## Motivation

Provide a reliable way to build container images and avoiding tenants any issue relating to building docker in docker or conflicting docker builds on shared hosts e.g. buildkit

## Requirements

* Tenants can build docker images from the P2P infrastructure
* Tenants can run docker containers in the jobs e.g. to run functional tests
* Tenants can run docker compose in their job

## Questions 

