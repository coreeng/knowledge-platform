+++
title = "Pipeline Generation"
date = 2022-01-15T11:50:10+02:00
weight = 3
chapter = false
pre = "<b></b>"
+++

## Motivation

Decouple the tenant from low level P2P tooling infra by generating the pipelines.

Better than templates as the shape can be updated centrally for all teams rather than teams having to upgrade to a new template.

## Requirements

Tenant pipelines are generated for the CI phase

Tenants aren’t coupled to the CI tool choice

## Additional Information

Take the pipeline shapes agreed in the definition phase and work out how they will be represented in the CI infrastructure tool of choice.

Pick the technique for how pipelines will be generated, templated, and how steps will be re-used.

Out typical approach is to have a config file in the repo of the tenant application that is tool agnostic.


