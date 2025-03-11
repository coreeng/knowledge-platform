+++
title = "M - Getting Unix tools on Windows"
weight = 13
chapter = false
+++


## M - Getting Unix tools on Windows

The tools we cover in this module are mostly available on Unix machines. To install UNIX debugging tools like **traceroute** and **curl** on Windows, you can use several approaches, such as using native Windows installations, Windows Subsystem for Linux (WSL), or third-party tools.

WSL allows you to run a Linux environment directly on Windows without the overhead of a virtual machine. This is one of the easiest and most integrated ways to install and use Linux tools on Windows.

Open PowerShell as Administrator and run the following command to enable WSL:

```
$ wsl --install
```

This installs WSL and Ubuntu by default. You can install other distros if needed.

To install the tools run:

```
$ sudo apt update && sudo apt install traceroute
```
