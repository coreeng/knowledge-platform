# Bootcamp - Canary Operator Reference App

## Pre-requisites

The local project setup uses `nix` with `direnv` so you need to install:
- [nix package manager](https://nixos.org/download.html)
- [direnv](https://direnv.net/)

After installing the above, you need to allow `direnv` in the project's directory:
- `direnv allow .`
- cd out of the current directory and back in for the `nix` setup to be installed
- [hook](https://direnv.net/docs/hook.html) `direnv` into your shell of choice and reload your shell

## How to test the canary operator on a local minikube cluster

For the canary operator task, a simplified CR called `canariedapp` is used. 
This would be the interface for the tenants:

```
apiVersion: canary.cecg.io/v1
kind: CanariedApp
metadata:
  name: canariedapp-sample
spec:
  replicas: 1
  image: bootcamp/minimal-ref-app:unstable-v1
  canary-spec:
    weight: 25
    replicas: 1
```

### Step 1 - Create a minikube cluster

`make bootstrap-cluster`

### Step 2 - Activate the minikube docker env

This is needed so that minikube can access all your built images directly, without the need to push 
to a local or remote docker repository. 

After the minikube cluster is up and running, run:

`eval "$(minikube -p minikube docker-env)"`

### Step 3 - Build the ref app images needed by the operator

You will need:
- an image with a "stable" minimal reference application
- an image with a "broken" minimal reference application

In order to build some test images for the minimal reference app go into the `minimal-reference-app` directory
in this repo, modify the test endpoint as required by returning for example a different http status code and 
re-build the image using `make docker-build`. 

In order 
To check that minikube has access to the above built images, run:
`minikube image list`

### Step 3 - Install your CRD and your operator into your cluster

```
make install
```

### Step 4 - create a CR 

There is a sample CR manifest already created for you that you can modify and apply:

```
kubectl apply -f config/samples/canary_v1_canariedapp.yaml
```

## How to develop on the canary operator

- When you modify the code for the operator run `make manifests`. There are more useful make tasks provided by kubebuilder
which you can find in this [Makefile](canary-operator/Makefile).
- To run the controller locally use `make run`. 

Note: When running the controller locally, because at the moment we are using the canary service to test the health,
you will need to port forward the service on your local for your controller to be able to access it. 

