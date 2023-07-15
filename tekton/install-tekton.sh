#!/bin/bash -x

kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
kubectl apply --filename https://storage.googleapis.com/tekton-releases/triggers/latest/release.yaml
kubectl apply --filename https://storage.googleapis.com/tekton-releases/triggers/latest/interceptors.yaml
kubectl apply --filename https://storage.googleapis.com/tekton-releases/dashboard/latest/release.yaml

# Wait for tekton to be ready

while [[ $(kubectl get pods -l app=tekton-pipelines-controller -n tekton-pipelines -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do echo "waiting for tekton to be ready" && sleep 1; done
while [[ $(kubectl get pods -l app=tekton-pipelines-webhook -n tekton-pipelines -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do echo "waiting for tekton to be ready" && sleep 1; done
while [[ $(kubectl get pods -l app=tekton-triggers-webhook -n tekton-pipelines -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do echo "waiting for tekton to be ready" && sleep 1; done
