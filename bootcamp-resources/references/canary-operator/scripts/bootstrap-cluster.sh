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

printf "\n\n Starting the minikube cluster \n\n"
minikube start

printf "\n\n Enable the ingress addon so that we can use the NGINX ingress controller \n\n"
minikube addons enable ingress

