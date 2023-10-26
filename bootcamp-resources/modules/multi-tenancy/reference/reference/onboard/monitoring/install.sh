#!/bin/bash

MONITORING_NAMESPACE="${MONITORING_NAMESPACE:=monitoring}"

echo "---Adding the prometheus community helm charts repo---"
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

NAMESPACES=$(kubectl get ns | grep monitoring | awk '{print $1}')

for namespace in $NAMESPACES; do

  NAMESPACE_MONITORING_HELM_CHART=${namespace}-prom

  #  firstly, check that no install exists for the namespace
  helm uninstall  --namespace ${namespace} ${NAMESPACE_MONITORING_HELM_CHART} || true

  echo "---Installing monitoring in the namespace: ${namespace}---"
  helm install --namespace ${namespace} ${NAMESPACE_MONITORING_HELM_CHART} prometheus-community/kube-prometheus-stack

done