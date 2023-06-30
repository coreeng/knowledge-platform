+++
title = "Workstation Setup"
weight = 2
chapter = false
+++

By the end of this stage you should have a development environment setup. 
These are all suggestions, if you have your preferred tools then go ahead and use those!

Tools to install:

* Brew: [Brew](https://brew.sh/)
    * Follow all setup instructions and add to your preferred shells setup scripts to have anything brew installs on the path
* GO
    * Run `brew install go`
* Java
    * Run `brew install openjdk@17`
    * When doing this read the output and make sure the symlink it suggests is created e.g.
        * `sudo ln -sfn /opt/homebrew/opt/openjdk@17/libexec/openjdk.jdk /Library/Java/JavaVirtualMachines/openjdk-17.jdk`
        * If you don’t do this jenv won’t work correctly
        * on Archlinux the `jenv` works without the symlink
* jEnv: [jEnv - Manage your Java environment](https://www.jenv.be/)
    * Install jEnv with the following command brew install jenv
    * Add the path in your `.bash_profile` and `init` jEnv by following the “Installation” instructions in the jEnv site. Remember to reload your bash profile (with `source .bash_profile` or opening a new terminal window.
    * Run `jenv add /opt/homebrew/Cellar/openjdk@17/17.0.5`  (might be a different minor version, and you may have to escape the @)
    * jEnv should output:
      ```text
      openjdk64-17.0.x added
      17.0.x added 
      ```
* Docker: [Install on Mac](https://docs.docker.com/desktop/install/mac-install/)
    * After installing, run the Docker application and give permissions. This will install `docker` and `docker-compose`
* IntelliJ: [IntelliJ IDEA](https://www.jetbrains.com/idea/download/#section=mac). Please check with your manager if you can access an IntelliJ IDEA Ultimate licence
    * If an Ultimate licence is available for you, download and install the IntelliJ IDEA Ultimate binaries
      * Make sure you have the Go plugin which can be installed from the IntelliJ interface under **Preferences -> Plugins**. Also make sure that the **Enable** go modules integration is checked :white_check_mark: under **Preferences -> Languages and Frameworks -> Go -> Go modules**
    * Alternatively, download and install IntelliJ IDEA Community Edition
      * Also download [VSCode](https://code.visualstudio.com/download)
      * Set up VSCode for Go development [Configure Visual Studio Code for Go development](https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code)
* Terminal
    * iTerm: [iTerm2 - macOS Terminal Replacement](https://iterm2.com/)  (or your preferred)
* Minikube & Kubernetes Tooling
    * Install with brew: [minikube](https://formulae.brew.sh/formula/minikube)
        * `brew install minikube`
    * Install with brew: [kubernetes-cli](https://formulae.brew.sh/formula/kubernetes-cli)
        * `brew install kubernetes-cli`
    * Install with brew: [helm](https://formulae.brew.sh/formula/helm)
        * `brew install helm`
* K6 (load testing)
    * Install with brew: [k6](https://formulae.brew.sh/formula/k6)
        * `brew install k6`
* Gradle
    * Install with brew: [gradle](https://formulae.brew.sh/formula/gradle)
        * `brew install gradle`





