+++
title = " "
date = 2022-01-15T11:50:10+02:00
weight = 13
chapter = false
pre = "<b>Platform Security</b>"
+++

### Security as a stakeholder
When building a container platform a company’s security department is a major stakeholder and should be consulted from the very beginning.

A container platform can offer many advantages to a security function:

* Enforced governance. Typically complaints from a security architect is that engineers don’t know about / care about security. Deploying to a central container platform can automate a lot of what typically is required from every development team. Some examples:
  * Vulnerability analysis in platform registries rather
* Shifted left security

## Build Time Protection
As much as possible should be verified pre-deployment. This can be more in P2P than Core Platform.

### Image Vulnerability Analysis 


**What:** Scan container images against industry standard vulnerability databases.

**Why:** To avoid vulnerable software being deployed to production.

**How:** 

Many open source, cloud and vendor products:

Open source:

* Clair: Tool parsing image contents and reporting vulnerabilities affecting the contents.
* Trivy: CLI for vulnerability scanning

Cloud:

* AWS: ECR offers:
  * Basic: Scan images with the Clair OS tool. Scan on push or manually.
  * Enhanced scanning with AWS Inspector. OS and Programming Language vulnerabilities.
* Google: Artefact registry and container registry

Vendor: 

* Qualys 
* Twistlock 

Caveat

* Ensure old images aren’t continuously scanned and produce false alerts

### Container Misconfiguration 
 
**What:** Detect common misconfiguration of Docker images

At build time detect misconfiguration in Docker files e.g.

* Last USER in the file should not be root (but there needs to be at least one USER statement)
* Tag the version of the FROM image explicitly (unless its scratch)
* Avoid using "latest" in the FROM statement
* Delete the apt-get lists after installing if not using scratch/distroless

For Kubernetes resources e.g.
* Ensure pods and controllers are not running as privileged
* Ensure pods images are hosted in a trusted ECR/GCR/ACR registry

**Why:** Avoid common security bad practices

**How:** 

* Trivy misconfiguration scanning: [Overview - Trivy ](https://aquasecurity.github.io/trivy/v0.24.1/misconfiguration/)
* [GitHub - hadolint/hadolint: Dockerfile linter, validate inline bash, written in Haskell](https://github.com/hadolint/hadolint)  or 
* [GitHub - goodwithtech/dockle: Container Image Linter for Security, Helping build the Best-Practice Docker Image, Easy to start  (covers more of the CIS bench mark)](https://github.com/goodwithtech/dockle)

### Kubernetes Deployment Misconfiguration 
**What:**

Enforce best practices for kubernetes deployments. Either at build time or prevention of a deployment.

**Why:**

Platform wide governance. 

**How:**

Trivy misconfiguration analysis

Kubebench

### Best Practice: Picking an image strategy
**What:**

Enforce docker security best practices by following these guidelines:
* Favour distroless container images, if not possible then use small container images like alpine
* Avoid unnecessary privileges, principle of least privilege (PoLP)
* Prevent sensitive information
* Sign images & verify them on runtime
* Run a container Image linter for Security, like Dockle

**Why:**

Minimise the types of containers and their attack surface running in the container platform making sure the images come from a trusted source and in the case the container is compromised an attacker will not have enough privileges

**How:**
* Defining and automating the process of building docker images
* Running Dockle in the CI pipeline and on a schedule basis
* Training application teams in the importance of docker security best practices
* Restricting the docker images application team can use from

### Library Vulnerability Analysis
**What:** More of an application pipeline concern but if the container platform includes a common P2P it should include a stage for tenant teams to run the correct vulnerability analysis tool for their tech stack.

**Why:** To find vulnerabilities in programming libraries which require patching or addressing CVEs

**How:**

* Running open source tools like 

[dependabot](https://github.com/dependabot) or similar

* [Mend (formerly WhiteSource) | Improving AppSec Outcomes ](https://www.whitesourcesoftware.com/)

 
### Code Vulnerability Analysis
**What:** scan the code written by an application team making sure the code is secured and free from vulnerabilities

**Why:** To find security vulnerabilities or hotspots in our own code which might be exploited by potential attackers

**How:**
* Running SAST (Static Application Security Testing) like:
* [SonarQube Code Security](https://www.sonarqube.org/features/security/)
* [Snyk Code](https://snyk.io/product/snyk-code/)
* [Sonatype Lift](https://lift.sonatype.com/getting-started)

 

### Centralised Artifact Repository 
**What:**
* A platform to centrally store any deployable components

**Why:**
* Single source of truth for every deployable
* Optimize build performance and storage costs by caching artifacts
* Avoid dependencies in other platforms (e.g. Dockerhub, NPM)

**How:**

By using a tool/platform like:
* [Artifactory - Universal Artifact Management](https://jfrog.com/artifactory/)
* [Nexus Repository Manager - Binary & Artifact Management | Sonatype](https://www.sonatype.com/products/nexus-repository?topnav=true) 

and making sure the artifacts are scanned for vulnerability setting minimum quality gates

### Signing of any deployable
**What:**
* Sign & verify deployables (docker images, helm charts, etc)

**Why:**
* To validate the authenticity of images and avoid supply chain attacks
* Prevent images from being built from dev work stations

**How:**
* Gatekeeper policy for image
* Restricted connectivity for pulling
* How with ECR/EKS for now
* From K8S 1.24 use sigstore as a new standard for signing, verifying and protecting software

## Runtime Protection
### Vulnerability Analysis
**What:** Continuously scan every image in-use 

**Why:** To catch vulnerabilities reported after deployment 

**How:**

Open source:
* Custom tools
* Trivy

Vendor: 
* Qualys 
* Crowdstrike 
* Twistlock (to confirm) 

### Misconfiguration Analysis 
**What:** 

Continuously validate k8s resources / prevent deployments

**Why:**

**How:**
* Everything enforced via policy manager e.g. OPA/jsPolicy
* If the policies are warn, continuous checking of the current state and prioritise fixing issues

### Malicious Activity

**What:** Suspicious activity

“Container Drift” 

**Why:** Detect breaches from inside containers

**How:**

Open Source:
* Falco (Falco ) 
* Prevention with image best practices / policy: no package managers, connectivity to install anything, etc, readonly file systems

Vendor:
* Crowdstrike

### Containers Privileged Access
**What:**
* Default that no tenants / pods can have privileged access
* Model for whose containers can do what

**Why:**

**How:**
* Default to non privileged access 
* Anything from platform team that requires privileged access should be built into the node image, not via a daemon set post installation 
* If it has to to run as a container in the cluster have a policy that explicitly whitelists those deployments 
* Implemented with policy manager

### Policy Manager 
**What:**
* Well defined what is and isn’t allowed in the cluster and who can do that

**Why:**
* Enforcing best practices/security practices

**How:**
* Trivy misconfiguration for Kubernetes
* OPA/Gatekeper or JSPolicy 
* Following Pod Security Standards: [Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/) 
* PodSecurityPolicy (deprecated): Use OPA/Gatekeeper to implement the same features 

### Pod Security Context Policy
**What:**
* Specific policy for what is expected/allowed for pod security contexts

**How:**
* OPA/Gatekepper or JSPolicy
* Different policy for tenants vs platform components 

## Kubernetes Node Protection
### Secure Host 
**What:**

Use standard hardening techniques e.g CIS for Linux on the Kubernetes nodes, if they are managed by the containers platform.

**How:**
* Regularly run linux bench
* Regularly

### Limited audited access to hosts
**What:**

Limit access to Kubernetes nodes and control plane nodes if not using a managed control plane.

**Why:**

…

**How:**

Depends on how nodes are provisioned

### Vulnerability Analysis
**What:**

Host vulnerability analysis.

**Why:**

**How:**
* There are many vendor and OS tools for this. Align with the client’s existing strategy.

### Network Logs
**What:**

Ability to turn on or continuous capturing of networking flow logs.

**Why:**

To debug / detect malicious activity.

**How:**

Depends on cloud provider / networking virtualisation. 

### Patch and updates 
**What:**

Kubernetes nodes should be continuously patched / updated. Including all Kubernetes components.

**Why:**

Keep up with the latest security updates.

**How:**

Environment dependent. 

## Tenant Access Control
### Authorization / Kubernetes RBAC - Christopher Batey
**What:**

A clearly defined, least privileged use of Kubernetes RBAC for developers and CI/CD agents.

**Why:**

**How:**
* AD Group / IAM role mapping to standard roles for:
  * Tenants
  * Applications
* RBAC should ensure:
  * Tenants have no access to namespaces of other tenants
  * Tenants have no access to platform namespaces
  * No cluster scoped access should be given to tenants apart from limited read only

### Engineer authentication: SSO Integration / 2FA

**What:**

Authentication to the Kubernetes API server should be via corporate SSO with 2FA.

**Why:**

Same access policies / management as other corporate RBAC.

Great user experience.

**How:**

* [How to Secure Your Kubernetes Cluster with OpenID Connect and RBAC ](https://developer.okta.com/blog/2021/11/08/k8s-api-server-oidc)

### CI/CD authentication
**What:**

How do deployments work from CI/CD. What Auth is used?

How does it work for manually kicked off jobs vs scheduled / triggered by commits 

**Why:**

 

**How:**
* Tech stack specific 

## Network Segregation
### Consider: Default deny or explicitly chose not to
**What:**

Consider implementing default deny, both within the cluster and outside.

**Why:**

Defence in depth

**How:**

This requires a CNI that supports NetworkPolicy and in most cases Cluster wide network policy

### Tenant managed network policy
**What:**

Give full control 

**Why:**

**How:**

* CNI that supports network policy e.g. CIlium

### Centrally managed network policy
**What:**

**Why:**

**How:**

* CNI that supports cluster wide network policy e.g. CIlium

### Egress firewalling
**What:** 

Likely implemented by cluster wide network policy. For multi tenant platforms different tenants need connectivity to different third parties. How is that managed?

**Why:**

Downstreams e.g. Persistence Tiers, typically open up connectivity fine grained e.g. this app can talk to this DB. We need to replicate that functionality in a multi tenant container platform.

**How:**
* Egress gateways with defined CIDR per client
* Open up full CIDR but have cluster wide network policies and agree with downstreams that the firewalling is managed by the platform not their firewall rules 

### Workload Segregation
Principle: no tenant should be able to adversely affect another tenant
* Enforced quotas?
* Enforced requests / limits?
* Testing of Auto scaling / Cost attribution 

## Corporate Integration
### Inventory Management
**What:**

Many enterprises have a central service catalog or inventory management. This is used to see who owns what service and for assigning incidents to.

**Why:**
* Assigning ownership to incidents
* Assigning vulberabilities

**How:**
* Service now is a common enterprise implementation

### SIEM
See [Container Platform SIEM](https://cecg-force.atlassian.net/wiki/spaces/CECG/pages/131629090)

## Secrets Management
Secrets should be passed / gathered at runtime, not built into images. Do not use environment variables for secrets. Ideally, centralise secrets to avoid security islands.

### Centralised Secrets Management
**What:**

* Centralise secrets / sensitive information so they are managed in a central place

**Why:**

The idea is to avoid different tools/platforms that come built-in with its own security components (that manage secrets, access control, audit, compliance, etc) but that does not facilitate interoperability with other tools and/or aggregation of security policies, management, and audit data.

**How:**

By consolidating in a single tool / platform, for example:
* [Vault by HashiCorp](https://www.vaultproject.io/)
* [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/)
* [CyberArk Secrets Manager](https://www.cyberark.com/resources/product-datasheets/cyberark-secrets-manager)

### Kubernetes Secrets Best Practices
**What:**

* Assuming that secrets end up in Kubernetes secrets follow these best practices

**Why:**

By definition, a secret must be kept secret preventing attackers from gaining access to sensitive information like passwords, auth tokens, api, keys, etc …

**How:**
* By following K8S secrets best practices:
* Enable encryption at rest for the data in etcd
* Limit access to etcd for admin users only
* Enable TLS/SSL between api server and the pods and api server and etcd 
* Avoid sharing or checking yaml/json files containing base64 secrets into a repo
* Limit which K8S users can read secrets
* Forbid exec into containers so secrets cannot be leaked by accessing environment variables or a secret mounted as a file
* Ideal secret management
  * Not available to everyone who has k8s access to the app i.e. not environment variable / file

### Secret Injection
**What:**

* How do secrets get into the application? 

**Why:**
* Understanding what/how and when secrets are injected is crucial to make sure sensitive information is not leaked/compromised

**How:**
* By making the secret available as a mounted file for the container which explicitly requires that information
* By adopting new patterns like:
  * Just-in-Time access to provide access only to specific components and for the time required
  * Secretless applications (not globally applicable): 
    * Applications don’t need secrets because they assume some service account that is mapoed to am IAM role that has permissions to get secrets do things
  * Whereby an application connects locally to a sidecar container which in turn authenticates the app, fetches the required credentials and establishes the connection (if required, e.g. database)
  * Kubernetes Secrets Integrations: An external secrets operator (e.g. External Secrets Operator ) whereby the secrets from an external secret management platform are automatically injected as K8S secrets

## Encryption in Transit
### Ingress TLS
**What:**

* TLS into the Container Platform

**How**
* L7 fronting LB terminates TLS
* Ingress controller terminates TLS

### Transparent inter service Encryption 
**What**
* Internal to the core platform, inter-service, encryption

**How**
* Service mesh
* Cilium ([Transparent Encryption (stable/beta) — Cilium 1.9.18 documentation](https://docs.cilium.io/en/v1.9/gettingstarted/encryption/) )

## Secure Kubernetes Setup
* Disable insecure access to the Kubernetes API
* Limit access to host + ability to use the container API

## Interesting articles:
### Docker Best Practices
* [Sysdig - Top 20 Dockerfile best practices](https://sysdig.com/blog/dockerfile-best-practices/)
* [Snyk - 10 Docker Security Best Practices](https://snyk.io/blog/10-docker-image-security-best-practices/)
* [Google Cloud - Best practices for building containers](https://cloud.google.com/architecture/best-practices-for-building-containers)
* [Infosec - Building container images using Dockerfile best practices](https://resources.infosecinstitute.com/topic/building-container-images-using-dockerfile-best-practices/)
* [Docker development best practices](https://docs.docker.com/develop/dev-best-practices/)
* [Medium - Dockerfile : Best practices for building an image](https://medium.com/swlh/dockerfile-best-practices-for-building-an-image-6120e512b1fa)

### Docker Image Strategy
* [Distroless Container Images](https://github.com/GoogleContainerTools/distroless)
* [Why distroless containers aren't the security solution you think they are](https://www.redhat.com/en/blog/why-distroless-containers-arent-security-solution-you-think-they-are)
* [The Bakery Model for Building Container Images and Microservices ](https://thenewstack.io/bakery-foundation-container-images-microservices/)

### Secrets Management
* [Security Islands](https://www.conjur.org/blog/security-islands/) 
* [Vault by HashiCorp](https://www.vaultproject.io/)
* [External Secrets Operator](https://external-secrets.io/v0.5.2/)
* [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/)
* [CyberArk Secrets Manager](https://www.cyberark.com/resources/product-datasheets/cyberark-secrets-manager)
