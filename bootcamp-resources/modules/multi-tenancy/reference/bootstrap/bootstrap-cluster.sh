#!/bin/bash
MINIKUBE_STATUS=$(minikube status | grep host:)
if [ "$MINIKUBE_STATUS" == "host: Running" ]; then
  echo "Minikube is running, do you want to tear it down and restart? (y/n)"
  read RESTART_MINIKUBE
  if [ "$RESTART_MINIKUBE" == "y" ]; then
      echo "Will restart minikube.."
    else
      echo "Will not restart minikube, exiting"
      exit
  fi
fi

minikube delete
minikube start --network-plugin=cni

printf "\n\nReinstalling cilium.\n\n"
cilium uninstall
cilium install
cilium hubble enable --ui

# this is needed otherwise cilium status doesn't detect that cilium is not ready
sleep 10

printf "\n\nWait for cilium status to report ok.\n\n"
cilium status --wait

printf "\n\nInstalling HNC\n\n"
kubectl apply -f bootstrap/manifests/hnc-controller.yaml

printf "\n\nWaiting for HNC to be installed\n\n"

while [[ $(kubectl -n hnc-system get pods -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do
   sleep 1
done

printf "\n\nEnable cilium network policies HNC propagation"
kubectl hns config set-resource ciliumnetworkpolicies --mode Propagate

# make sure krew exists in the path before attempting to use it
export PATH="${PATH}:${HOME}/.krew/bin"
kubectl krew update && kubectl krew install hns
