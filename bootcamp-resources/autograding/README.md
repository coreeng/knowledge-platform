# Autograding bootcamp modules based on acceptance criteria
 
This directory holds the autograding functionality based on acceptance criteria for the [bootcamp modules](../../knowledge-base/content/bootcamp/modules).

## Overview 

At the moment we do autograding using BDD testing in Go, using [GoDog](https://github.com/cucumber/godog).

## Structure

Each autograding module has its own directory within the `autograding` folder, which contains:
    - the Go BDD tests packaged in a docker image
    - the manifests that are needed to run the autograding in the cluster
    - a make file that represents the interface of each autograding module

## Results

At the moment, there are two ways to visualise the test results:
- by running the `autograde-<module>` task and visualising the results
- by looking at the metrics published in pushgateway 

## How to create your own autograding module

- Create a new directory under `autograding`
- Create a `Makefile` that has the following task
    - a "public" task with the prefix  `autograde-`. Public task means it can be displayed as available task when imported in 
      reference skeletons and a user types `make`. In order to make a task public add a comment after the task starting with `##`.
    - non-public convenience tasks that are useful for the lifecycle of your autograding module.
      If you want to see those tasks when you type `make`, add a comment that starts with a single `#`. 
    - all tasks should have a name specific to the autograding you are doing, this will prevent name clashes when multiple
      make files are included in the same skeleton app.  
- Create a Dockerfile that packages your autograding functionality. Have a look at the examples in `autograding`, at the moment
  we have acceptance tests used for autograding some bootcamp standard modules. 
- Create the manifests/infra you need for running the tests. Existing examples use kubernetes jobs.
- Create a deployment pipeline for validating and publishing your image. Have a look at how that is done in [GitHub actions](/.github/workflows).
- Include the `Makefile` in your skeleton app `Makefile` and run the `autograde-<module>` task from your skeleton app. 
  If you have multiple `autograde-<module>` jobs to be run, you can add an `autograde` job that runs them all in 
  the same time. For an example of how that is done have a look at the [Reference Go App Makefile](/bootcamp-resources/skeletons/reference-application-go/Makefile).
