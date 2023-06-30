+++
title = "Monitoring Stack Setup"
weight = 4
chapter = false
+++

### Motivation

* Learn the most common monitoring stack in use today with Kubernetes: Grafana, Prometheus, & Alert Manager

### Requirements

* Grafana / Prometheus and Alert Manager setup on local Kubernetes Cluster
  * Everything to setup your environment should be in code 
  * Scripts required for installing it should be included in your forked reference application in a new folder called `monitoring` 
* Resource stats available for your reference application

### Additional Information

There is a handy helm chart that installs all three. Unfortunately it doesn’t quite work with the latest version of minikube and Kubernetes. 
We have a fork in CECG that does. GitHub - coreeng/helm-charts: [Prometheus community Helm charts](https://github.com/coreeng/helm-charts)

```bash
git clone git@github.com:coreeng/helm-charts.git
cd helm-charts/charts/kube-prometheus-stack
helm repo add grafana https://grafana.github.io/helm-charts
helm install prom ./ --set prometheus.prometheusSpec.enableRemoteWriteReceiver=true
```

We set `-prometheus.prometheusSpec.enableRemoteWriteReceiver` to allow k6 to write metrics directly rather than having prometheus scrape them during load testing.

{{% notice note %}}
Look at the commit history of the CECG fork and try and work out what changed in Kubernetes that stopped the original helm chart from working
{{% /notice %}}

Unless the helm chart has changed, and if it has please update these instructions, you should have three services. If they are different (i.e. the helm chart has changed) please update this page!

```
prom-grafana
prom-kube-prometheus-stack-alertmanager
prom-kube-prometheus-stack-prometheus
 ```

You can use kube [port-forward](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/) to access these as the helm chart doesn’t create ingress’s for them by default.  For example to access grafana:


```
kubectl port-forward service/prom-grafana 8000:80
kubectl port-forward service/prom-kube-prometheus-stack-prometheus 9090
kubectl port-forward service/prom-kube-prometheus-stack-alertmanager 9093
```

You should then be able to access

* Grafana
* Prometheus
* Alert Manager

At the time of writing the helm chart sets the grafana credentials to admin/prom-operator.

{{% notice note %}}
For an extra challenge, find out how this password is set, and upgrade your helm installation with a new password
{{% /notice %}}

### Questions / Defuzz / Decisions

### Deliverables (For Epic)

* Grafana deployed and accessible
* Prometheus deployed and accessible
* Alert manager deployed and accessible
* Resource stats available for the reference application
