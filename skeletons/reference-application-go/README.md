# CECG Reference Application - GoLang

## Path to Production (P2P) Interface

The P2P interface is how the generated pipelines interact with the repo.
For the CECG reference this follows the [3 musketeers pattern](https://3musketeers.io/) of using:

* Make
* Docker
* Compose

These all need to be installed.

## Path to Production (P2P) Tooling

### GitHub Actions

ci.yaml in [.github/workflows](.github/workflows) shows how the same GitHub action can be used for many repos using different technologies as long as the same Make targets are defined.

### Tekton

[Instructions for how to setup](../../tekton/README.md) a WebHook that calls a locally running Tekton.

## Structure

### Service

Service source code, using Go.

### Functional

Stubbed Functional Tests using [Cucumber Godog](https://github.com/cucumber/godog)

### NFT

Load tests using [K6](https://k6.io/).

## Running the application locally

### Application

```
make run-local
```

This application is exposed locally on port 8080 as well as being available to the tests when run with make.
This is as they are in the same docker network.

### Functional Tests

```
make stubbed-functional
```

You should see:

```
4 scenarios (4 passed)
16 steps (16 passed)

```

### Non-Functional Tests

```
make stubbed-nft
```

You should see:

```
     ✓ status was 200
     
     checks.........................: 100.00% ✓ 6590       ✗ 0    
     data_received..................: 850 kB  14 kB/s
     data_sent......................: 560 kB  9.2 kB/s
     http_req_blocked...............: avg=22.7µs  min=0s    med=4µs    max=2.9ms   p(90)=10µs   p(95)=29µs   
     http_req_connecting............: avg=11.99µs min=0s    med=0s     max=1.87ms  p(90)=0s     p(95)=0s     
     http_req_duration..............: avg=2.86ms  min=649µs med=2.02ms max=25.18ms p(90)=4.91ms p(95)=7.27ms 
       { expected_response:true }...: avg=2.86ms  min=649µs med=2.02ms max=25.18ms p(90)=4.91ms p(95)=7.27ms 
     http_req_failed................: 0.00%   ✓ 0          ✗ 6590 
     http_req_receiving.............: avg=40.32µs min=5µs   med=23µs   max=8.7ms   p(90)=63µs   p(95)=93.54µs
     http_req_sending...............: avg=32.3µs  min=2µs   med=13µs   max=5.17ms  p(90)=41µs   p(95)=74.54µs
     http_req_tls_handshaking.......: avg=0s      min=0s    med=0s     max=0s      p(90)=0s     p(95)=0s     
     http_req_waiting...............: avg=2.79ms  min=631µs med=1.96ms max=25ms    p(90)=4.81ms p(95)=7.12ms 
     http_reqs......................: 6590    108.135663/s
     iteration_duration.............: avg=1s      min=1s    med=1s     max=1.02s   p(90)=1s     p(95)=1.01s  
     iterations.....................: 6590    108.135663/s
     vus............................: 52      min=10       max=199
     vus_max........................: 200     min=200      max=200

```

## How to run the service in minikube

### Install Prerequisites

* A minikube cluster i.e. you've run `minikube start`
    * You need to enable the ingress addon by executing: `minikube addons enable ingress` and then follow the instructions from the output e.g. if on mac run `minikube tunnel`
* `kubectl` or use `minikube kubectl`

### Set up a local registry

How to set up a local registry on minikube:

* Enable the registry addon by executing: `minikube addons enable registry`
* The registry will be exposed on port 5000
* As the registry operates over an insecure connection, the Docker flag `--insecure-registry` will need to be set, so you need to run: `minikube start --insecure-registry minikube:5000`

For Mac users: please follow the instruction here: https://minikube.sigs.k8s.io/docs/handbook/registry/#docker-on-macos .
You will need to redirect port 5000 on the docker virtual machine over to port 5000 on the minikube machine.
Then you will be able to access your local registry on `localhost:5000`.

#### Using minikube's docker

An alternative to setting up a local registry would be to execute: `eval $(minikube -p minikube docker-env)`, which means that any
docker commands you execute will interact with the docker daemon running inside Minikube, rather than your local Docker environment.
The downside to this method is that you will need to execute this on any new shell you open up, but it's easier to set up.

Keep in mind that if you use this method you should set an `imagePullPolicy: Never` to your `deployment-minikube.yml` for your `reference-service`,
which would signify that the image should be expected to exist locally.

#### Push the image

Prerequisites:

- The `registry` in the [Makefile](Makefile) should be updated with your newly created registry.

```
make docker-build
make docker-push (If using the second method(Minikube's docker) you can skip this. Building the image is enough
```

If using the second method described above([Using minikube's docker](README.md#using-minikubes-docker)), you can execute
`minikube image ls` to see whether your image is in minikube.

### Deploy the service

Prerequisites:

If you are using the [first method](#set-up-a-local-registry):

- if your local registry is not on `minikube:5000` you will need to update the [deployment yml](service/k8s-manifests/deployment-minikube.yml)
  to pull the image from your local repository. e.g: `localhost:5000` if you are using docker machine on Mac and you followed the registry setup steps from above.

Create the namespace, secrets and the deployments:

```
kubectl apply -f service/k8s-manifests/namespace.yml
kubectl apply -f service/k8s-manifests/pv-dbdata.yml
kubectl apply -f service/k8s-manifests/pvc-dbdata.yml
kubectl apply -f service/k8s-manifests/secret-db.yml
kubectl apply -f service/k8s-manifests/deployment-minikube.yml 
```

Deploy the ingress and services:

```
kubectl apply -f service/k8s-manifests/expose.yml
```

Check that the service is running:

```
kubectl get pods -n reference-service-showcase
NAME                                 READY   STATUS    RESTARTS   AGE
reference-service-7cff68d485-q8mw5   1/1     Running   0          142m
```

Check that the ingress is created:

```
kubectl get ingress -n reference-service-showcase
NAME                CLASS   HOSTS   ADDRESS        PORTS   AGE
reference-service   nginx   *       192.168.49.2   80      144m
```

If on Linux or MacOS you can now access the service on the IP address (which is the minikube IP).

```
curl localhost/service/hello
Hello World!%
```

If this doesn't work ensure you followed the instructions when enabling the minikube ingress addon.

### Run the functional tests against deployed application

Prerequisites:
* [Godog binary](https://github.com/cucumber/godog#step-2---install-godog) is installed and in your $PATH.

This shows how you can run the same tests locally and on a deployed version using the `SERVICE_ENDPOINT` environment variable.

E.g: For a local run: 

```
cd functional/godogs
SERVICE_ENDPOINT="http://localhost:8080" godog run
```

### Run the non-functional tests against deployed application

For a local run: 

```
SERVICE_ENDPOINT="http://localhost:8080" k6 run ./nft/ramp-up/test.js
```

# Trigger GH actions