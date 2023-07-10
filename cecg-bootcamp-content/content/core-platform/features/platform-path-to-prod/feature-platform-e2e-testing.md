+++
title = "Platform E2E Testing"
date = 2022-01-15T11:50:10+02:00
weight = 10
chapter = false
pre = "<b></b>"
+++

### Motivation

Well-tested platform with tests for all platform features to avoid any regressions

### Requirements

* Each platform feature is captured in an end to end acceptance test.
* Platform team notified of any test failures
* Ways of working updated to say that a red test should be the top priority

## Additional Information

For Monolithic deployments the tests can run after the deployment. Decoupled deployments are a bit tricker, likely there needs to be:

* Per components tests
* E2E platform tests that cover features that span platform components

What will be delivered? Report on all platform features functioning, updated every time a deployment takes place. 
