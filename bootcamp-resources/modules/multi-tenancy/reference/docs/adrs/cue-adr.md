# 1. ADR for Cue

Date: 2023-03-30

## Status

In review.

## Context

In order to template and generate k8 manifests based on user input we would ideally use a tool that also helps us validate
the manifests against the k8 APIs. 

## Decision

[Cuelang](https://cuelang.org/docs/) is tool of choice because of the following reasons:
- the syntax for `cue` files is readable
- it offers both client-side and server-side validation of the `cue` manifests files:
  - client side validation is useful when we validate user input - e.g: tenant configuration
  - server side validation means we can validate the manifests against the k8 api. The tool integrates with go api clients
    that can give us validation for custom resources like `hnc` or `tekton` related resources
- it integrates with the go ecosystem - this can be useful when you need both languages for specific tasks 
- it can produce `yaml` or `json` manifests based on the `cue` manifests which means that those can be easily 
  visualised and synced with let's say `flux` against the cluster
- it can generate `cue` files based on `yaml` manifests, so you don't need to manually write the manifests
- it offers the ability to use scripting by building user defined subcommands, so you can use `cue` to run your shell
  scripts, or apply resources with `kubectl` against your cluster

## Consequences

- using `cue` means that the learning curve can be a bit steep when working for the first time with it


