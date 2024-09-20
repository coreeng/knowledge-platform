# Knowledge Base

The knowledge base is made up of:

* Hands on [bootcamp](./content/bootcamp/)

## Running the Site Locally

You can run the site locally using the following methods:

### Docker

**Prerequisites:**  
Make sure Docker is installed on your system. If not, you can download it [here](https://www.docker.com/products/docker-desktop/).

```sh
make run-local
```

> **_NOTE:_** You will see your changes live as you make them.

### Hugo

**Prerequisites:**  
Install Hugo by following the instructions [here](https://gohugo.io/installation/).

Start the Hugo server with:

```sh
hugo serve
```

> **_NOTE:_** With Hugo, you can see your changes live as you make them.

## Versioning and releasing

We use semantic versioning: `v{major}.{minor}.{patch}`, e.g. `v0.1.0`.

Merging a pull request will generate a new **minor** version, and a **tag** will be automatically created in the repository.

Next, an image will be built and pushed to the Docker repository with the new version.
