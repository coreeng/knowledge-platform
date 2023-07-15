+++
title = "Continuous E2E Testing"
date = 2022-01-15T11:50:10+02:00
weight = 11
chapter = false
pre = "<b></b>"
+++

### Motivation

Find issues in the platform that occur not in response to a deployment.

E.g a test that tests external-dns creation of entries in Route53 might break as external-dns has lost connectivity to the AWS API due to something out side of the platform's control.

### Requirements

* Continuously run the platform E2E tests
* Results for N days visible for the platform team
* Platform team notified of failures

### Additional Information

After this feature, the Platform e2e tests will be continuously running. Detection of issues will be quicker. Temporary blips not caught by tests running after deployment will be detected.


