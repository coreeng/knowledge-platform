+++
title = "Encryption in transit"
date = 2022-01-15T11:50:10+02:00
weight = 7
chapter = false
pre = "<b></b>"
+++

## Motivation
Encryption in transit without any tenant effort. Make security easy.

## Requirements
* All tenant network paths encrypted by default e.g. mTLS

## Additional Information
Encrypt the data before transmitting it and decrypting the data on arrival between authenticated endpoints.
Can be implemented transparently to the tenants with various CNIs and service meshes. 



