+++
title = "Retrospective Policy Validation"
date = 2022-01-15T11:50:10+02:00
weight = 5
chapter = false
pre = "<b></b>"
+++

## Motivation
Be able to deploy new policies, knowing if any existing resources validate them.

Without this previously deployed applications re-deploying or even just moving between nodes may fail to re-deploy.

## Requirements
For every new policy, validate existing deployments conform to it. Out of the box solutions don't support this.

## Additional Information
This is not typically built into standard kubernetes policy controllers.



