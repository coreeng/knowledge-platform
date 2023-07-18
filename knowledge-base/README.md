# Knowledge Base

The knowledge base is made up of:

* Industry platform engineering modules on [Developer Platforms](./content/core-platform/) & [Consolidated P2P (paved paths)](./content/core-p2p/)
* Hands on [bootcamp](./content/bootcamp/)

## Running Locally

The site can be run locally with Docker:

Build the docker image:

```
docker build . -t knowledge-platform
```

Then run it locally:

```
docker run -d -p 8080:8080 knowledge-platform
```