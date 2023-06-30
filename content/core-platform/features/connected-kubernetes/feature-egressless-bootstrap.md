+++
title = "Egressless Bootstrap"
date = 2022-01-15T11:50:10+02:00
weight = 9
chapter = false
pre = "<b></b>"
+++

### Motivation

Be able to start the kubernetes cluster without node egress to the Internet. Very common at Banks.

### Requirements

* All required images are synced to an internal registry 
* All image names are re-written to the internal name e.g. by a mutating web hook

