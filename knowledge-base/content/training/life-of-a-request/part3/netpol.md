+++
title = "More about Network Policies"
weight = 2
chapter = false
+++

#### Overview

Similarly to Web App, we don’t want User API to be able to call any destinations except the allowed ones. We are applying the Default Deny network policy that forbids all **outgoing** traffic from the **backend** namespace with the exception of DNS port **53**.

![Network Policies for User API](/images/loar/3-2.png)
_Figure 3-2. Network Policies for User API_

The following network policy definition applies DefaultDeny for backend egress, but allows DNS resolution:

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: default-deny-with-dns-access
  namespace: backend
spec:
  podSelector: {}
  policyTypes:
    - Egress
  egress:
  - to:
    - namespaceSelector:
        matchLabels:
          kubernetes.io/metadata.name: kube-system
      podSelector:
        matchLabels:
          k8s-app: kube-dns
    ports:
    - protocol: UDP
      port: 53
```

To test the network policy let’s call curl from inside the User API pod:

```
$ kubectl exec -it user-api-54d897b874-vnpd6 -n backend -- curl -v --max-time 5 https://httpbin.org/uuid

*   Trying 52.87.99.32:443...
* ipv4 connect timeout after 2498ms, move on!
*   Trying 100.29.106.188:443...
* Connection timed out after 5000 milliseconds
* Closing connection 0
curl: (28) Connection timed out after 5000 milliseconds
```

We see that the IP address for **httpbin.org** was successfully resolved, but the connection is timed out.

#### Allowing external calls {#allowing-external-calls}

We need to explicitly allow calling the external addresses:

```
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-external-outbound
  namespace: backend
spec:
  podSelector:
    matchLabels:
      app: user-api
  policyTypes:
  - Egress
  egress:
  - to:
    - ipBlock:
        cidr: 0.0.0.0/0
```

Now let’s call **curl** again from inside the User API pod:

```
$ kubectl exec -it user-api-54d897b874-vnpd6 -n backend -- curl https://httpbin.org/uuid

{
  "uuid": "9185e170-9efe-4d2e-ac3f-694885fd0700"
}
```

We successfully called a third party API from **user-api** pod. Now we have all the data needed to generate the response to the user.
