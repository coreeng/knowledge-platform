# Acceptance Criteria

This module holds the acceptance criteria tests for the standard bootcamp modules. These standard modules are:

* P2P-fast-feedback
* NFT & Observability
* Multi-tenant Kubernetes
* Platform Engineering with Golang
* Debugging for Linux, Containers and Kubernetes
* Cloud, IaC and Managed K8s

## Structure

In our project structure, each module will have its own directory within the acceptance-criteria folder. We plan to
adopt a Behavior-Driven Development (BDD) testing approach, using the GoLang programming language for our
acceptance-criteria functional tests. We will also be using the [GoDog](https://github.com/cucumber/godog) framework,
which functions similarly to Cucumber in Java projects.

The functional tests will then be containerized, allowing us to execute them as a Job in-cluster. After the job has run,
we would have pushed some metrics to our pushgateway, depending on the success or failure of each step.

See the [acceptance criteria for p2p-fast-feedback](p2p-fast-feedback) for working examples.