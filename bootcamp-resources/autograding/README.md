# Autograding bootcamp modules based on acceptance criteria
 
This directory holds the autograding functionality based on acceptance criteria for the [bootcamp modules](../../knowledge-base/content/bootcamp/modules).

## Overview 

At the moment we do autograding using BDD testing in Go, using [GoDog](https://github.com/cucumber/godog).

## Structure

Each autograding module has its own directory within the `autograding` folder, which contains:
    - the Go BDD tests packaged in a docker image
    - the helm charts for installing the necessary infrastructure
 
## How to run the autograding locally

- `make install-pushgateway`
- `make MODULE=<module-name> autograde`

To see the metrics for the autograding tests:
- `make port-forward-pushagateway`
- go to http://localhost:9091

Note: at the moment, when building locally, the image is pushed to the minikube repo. Have a look at the `build` task. 
If you have a different setup make sure you build and push the image to the respective repo.

## How to create your own autograding module

- Create a new directory under `autograding`
- Create a Dockerfile that packages your autograding functionality. Have a look at the examples in `autograding`, at the moment
  we have acceptance tests used for autograding some bootcamp standard modules. 
- Create the helm charts needed for installation of the module
- Create a deployment pipeline for validating and publishing the image and the helm charts.
- Add your modules in the SUPPORTED_MODULES in the `Makefile` and test the autograding by running `make MODULE=<your-new-module> autograde`
- Add a task in the skeleton app to install the helm charts that run the autograding jobs