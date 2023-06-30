+++
title = "Counter Resets"
weight = 7
chapter = false
+++

## Motivation

Learn by doing: adding features with proper tests for the P2P.

## Requirements

* Do this for one of the reference applications: Java or Go
* Reset counter. Including:
  * Unit tests
  * Functional tests
  * Non-functional tests

The following HTTP endpoints are available:
* `GET /counter/{name}`  get the current counter of that name. if it doesn't exist, return 0.
* `PUT /counter/{name}` - increments the counter of the name. if it doesn't exist, create one and then increment it.
  * If this were a POST, should it have the same behaviour? What’s the difference between PUT and POST?


You need to implement:
* `DELETE /counter/{name}` - delete the counter for that name


## Questions / Defuzz / Decisions

If this were a POST, should it have the same behaviour? What’s the difference between PUT and POST?

## Deliverables (For Epic)

- ...
