# How to run Tekton in a local Kubernetes Cluster

[Tekton](https://tekton.dev/) is used to run the Path to Production locally.

## Prerequisites

- `minikube`
- a local docker registry in `minikube`
- the service is deployed in `minikube`

Have a look at the [Readme Minikube section](../reference-application-go/README.md#how-to-run-the-service-in-minikube) on instructions on how to set these up.

## Install Tekton in minikube

`./tekton/install-tekton.sh`

Then install the [Tekton CLI](https://tekton.dev/docs/cli/).

## Deploy the event listener and PR pipeline

`./tekton/install-pipeline.sh`

This creates everything in a namespace called `reference-service-ci`

## Expose the Tekton dashboard

```
kubectl --namespace tekton-pipelines port-forward svc/tekton-dashboard 9097:9097
```

[Tekton Dashboard](http://localhost:9097)

## Expose the event listener to a GitHub WebHook

To allow GitHub webhooks to connect to a local cluster.

Expose the service with port forwarding:
```
kubectl port-forward -n reference-service-ci services/el-github-listener 8080
```

Expose your local workstation e.g. with [ngrok](https://formulae.brew.sh/cask/ngrok):

```
ngrok http 8080
```

In a real environment the event listener would need to be exposed via ingress or a service load balancer in GKE or EKS.

## GitHub SSH Key

Tekton requires a ssh key to checkout the repo.

You need to generate an `id_ed25519` ssh key by following the instructions [here](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent#generating-a-new-ssh-key)(Skip step #4)

**When prompted, please make sure to leave the passphrase of the key empty** 

Execute: 
```
$ pbcopy < ~/.ssh/id_ed25519.pub
# Copies the contents of the id_ed25519.pub file to your clipboard
```

Starting from step #2, add your newly created public key as a [Deploy Key](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/managing-deploy-keys#set-up-deploy-keys)

Convert the contents of the private key to base64:

```
cat ~/.ssh/id_ed25519 | base64
```

Put the base64 private key and known hosts file into a secret called `git-credentials`. You can use the template below

You can use the existing known hosts base64 within `git-credentials-template.yaml` which includes the GitHub IPs.
After creating the secret it should be applied to the Minikube cluster, e.g. `kubectl apply -f tekton/git-credentials.yaml`

```
apiVersion: v1
kind: Secret
metadata:
  name: git-credentials
  namespace: reference-service-ci
data:
  id_ed25519: LS0ta...
  known_hosts: Z2l0aHViLmNvbSwxNDAuODIuMTIxLjMgc3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBQkl3QUFBUUVBcTJBN2hSR21kbm05dFVEYk85SURTd0JLNlRiUWErUFhZUENQeTZyYlRyVHR3N1BIa2NjS3JwcDB5VmhwNUhkRUljS3I2cExsVkRCZk9MWDlRVXN5Q09WMHd6ZmpJSk5sR0VZc2RsTEppekhoYm4ybVVqdlNBSFFxWkVUWVA4MWVGekxRTm5QSHQ0RVZWVWg3VmZERVNVODRLZXptRDVRbFdwWExtdlUzMS95TWYrU2U4eGhIVHZLU0NaSUZJbVd3b0c2bWJVb1dmOW56cElvYVNqQit3ZXFxVVVtcGFhYXNYVmFsNzJKK1VYMkIrMlJQVzNSY1QwZU96UWdxbEpMM1JLclRKdmRzakUzSkVBdkdxM2xHSFNaWHkyOEczc2t1YTJTbVZpL3c0eUNFNmdiT0RxblRXbGc3K3dDNjA0eWRHWEE4VkppUzVhcDQzSlhpVUZGQWFRPT0KMTQwLjgyLjEyMS40IHNzaC1yc2EgQUFBQUIzTnphQzF5YzJFQUFBQUJJd0FBQVFFQXEyQTdoUkdtZG5tOXRVRGJPOUlEU3dCSzZUYlFhK1BYWVBDUHk2cmJUclR0dzdQSGtjY0tycHAweVZocDVIZEVJY0tyNnBMbFZEQmZPTFg5UVVzeUNPVjB3emZqSUpObEdFWXNkbExKaXpIaGJuMm1VanZTQUhRcVpFVFlQODFlRnpMUU5uUEh0NEVWVlVoN1ZmREVTVTg0S2V6bUQ1UWxXcFhMbXZVMzEveU1mK1NlOHhoSFR2S1NDWklGSW1Xd29HNm1iVW9XZjluenBJb2FTakIrd2VxcVVVbXBhYWFzWFZhbDcySitVWDJCKzJSUFczUmNUMGVPelFncWxKTDNSS3JUSnZkc2pFM0pFQXZHcTNsR0hTWlh5MjhHM3NrdWEyU21WaS93NHlDRTZnYk9EcW5UV2xnNyt3QzYwNHlkR1hBOFZKaVM1YXA0M0pYaVVGRkFhUT09Cg==
```

## Minikube Credentials

In order for Tekton to deploy the service for testing on Minikube, credentials are required.  On installation Minikube will create a client key and certificate in `~/.minikube/profiles/minikube/` and these can be copied into a Kubernetes secret for use by Tekton.

You can use the existing `minikube-credetials-template.yaml` file, replacing `MINIKUBE_CLIENT_KEY` and `MINIKUBE_CLIENT_CRT` with the Base64 encoded values found in `~/.minikube/profiles/minikube/client.key` and `~/.minikube/profiles/minikube/client.crt`.

```
Get the base64 client key:  cat ~/.minikube/profiles/minikube/client.key | base64
Get the base64 client crt:  cat ~/.minikube/profiles/minikube/client.crt | base64
```
After updating the secret it should be applied to the Minikube cluster, e.g. `kubectl apply -f tekton/minikube-credentials.yaml`

## Create a web hook on your GitHub fork repo

For GitHub to call out to Tekton for CI we need a webhook.

Go to your fork -> Settings -> Webhook -> Add WebHook

* URL: should be the ngrok URL. This expires every 2 hours unless you have an account so this will need updating.
* Content Type: `application/json`
* Secret: 1234567 (match up with the secret in `reference-event-listener.yaml`)

In a real scenario the secret could be from a system like Vault and the event listener would have some firewall rules
to only accept connections from the GitHub IPs.

## Testing it out!

Open a PR on your fork or push to main. That should trigger the webhook and a pipeline run in Tekton.

