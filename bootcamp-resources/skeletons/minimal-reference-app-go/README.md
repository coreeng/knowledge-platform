## Minimal reference app Go

The purpose of this app is to serve as a minimal go app that exposes a metric against which we can validate the health of the app using the alerts in alertmanager. 

The app is intentionally designed to satisy the following:
- stateless
- two very basic endpoints for pass and failure so that we can pick one of them for validation when testing canary deployments