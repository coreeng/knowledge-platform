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