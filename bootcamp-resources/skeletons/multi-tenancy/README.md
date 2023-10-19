# Skeleton for multi tenancy app

This represents the starting point for the multi tenancy bootcamp module.
The skeleton has an interface defined in the `Makefile`, some of the tasks need to be implemented.

## Pre-requisites

- `make`
- `helm`

## Make tasks 

The following tasks are provided:
- `make bootstrap` -> starts up a minikube cluster
- `make autograde` -> runs the autograding job in the cluster

The following tasks need to be implemented:
- `make run` -> this task should run all the functionality that implements the module