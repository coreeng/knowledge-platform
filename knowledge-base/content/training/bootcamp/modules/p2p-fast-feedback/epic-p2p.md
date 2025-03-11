+++
title = "Consolidated P2P"
weight = 6
chapter = false
+++

## Motivation

Understand how the same P2P can be used for application that use difference tech stacks.

## Requirements

* Share the same P2P between the Java and Go reference applications 
  * Have both cloned into your own private repository, or one cloned twice
* Custom steps for each application should be in the Make file rather than in the pipeline definition so we achieve a single action used across different apps with specific implementation living in the app itself

## Questions / Defuzz / Decisions

The approach for this should be the "Automatically generated as a platform engineering service" as described in consolidated P2P.

#### Tekton

If using tekton create a separate repo that can generate the pipelines for two applications.
Name the repo `bootcamp-p2p-<github handle>`

#### GitHub Actions

Pull out the P2P in a common action used by both repos.

#### Q: Which two repos? 
A: You can copy the same reference application and assign different names or copy both the java and golang application


## Deliverables (For Epic)

- ... 
