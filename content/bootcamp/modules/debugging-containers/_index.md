+++
title = "Debugging for Linux, Containers, and Kubernetes"
date = 2022-12-22T05:07:10+02:00
weight = 5
chapter = false
+++

## Introduction 

**What should you {{< colour green understand >}} at the end of the week?**

* Root cause analysis in a Linux / Containers / Kubernetes environment
* Badly designed applications
  * Not handling signals correctly: pods get killed because they didn’t gracefully shut down
* How Cilium network policies fit in
* Memory constraints
* CPU Constraints

**What should you be {{< colour green "able to do" >}} by the end of the week?**

* Debug and resolve the following issues in Kubernetes
  * Kubernetes probe issues
  * Signal handling & graceful shutdown
  * Application level misconfiguration
  * Networking issues
  * Resource related issues



