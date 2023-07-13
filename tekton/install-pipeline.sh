#!/bin/bash -x

kubectl apply -f tekton/reference-event-listener.yaml
kubectl apply -f tekton/ci-pipeline.yaml
kubectl apply -f https://raw.githubusercontent.com/tektoncd/catalog/main/task/git-clone/0.9/git-clone.yaml -n reference-service-ci
