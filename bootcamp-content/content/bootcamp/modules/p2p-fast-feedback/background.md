+++
title = "Background"
weight = 2
chapter = false
+++

The goal of the week is to deploy and modify a reference service. If any of the technologies are new to you, then spend time doing the following.

The modules assumes you have access to at least one of the following:
* [A Cloud Guru](https://acloudguru.com/)
* [Pluralsight](https://www.pluralsight.com/)
 
### Web and Networking

{{%expand "Materials on TCP and HTTP" %}}
* TCP
    * [What is the TCP/IP Model? Layers and Protocols Explained](https://www.freecodecamp.org/news/what-is-tcp-ip-layers-and-protocols-explained/)
    * [TCP vs. UDP — What's the Difference and Which Protocol is Faster?](https://www.freecodecamp.org/news/tcp-vs-udp/)
* HTTP
    * [An introduction to HTTP: everything you need to know](https://www.freecodecamp.org/news/http-and-everything-you-need-to-know-about-it/)
    * [HTTP vs HTTPS – What's the Difference?](https://www.freecodecamp.org/news/http-vs-https/)
* WebSockets
  {{% /expand%}}

### Golang programming

The primary language used in the platform engineering domain. During the bootcamp you'll use Golang to:
* Build an operator
* Modify an application deployed to Kubernetes

{{%expand "Materials on Golang and its ecosystem" %}}
* [Getting started](https://go.dev/doc/tutorial/getting-started)
* [Go by example](https://gobyexample.com/)

{{% /expand%}}

### Java programming

You don’t need to be an expert in Java but many of the applications deployed to the platform are written in Java. 
If you choose to use the Java reference application you can do tutorials online or pick it up following the 
specific tutorials like Spring Boot

{{%expand "Materials on Java and its ecosystem" %}}
* Learn Java
    * LeetCode
    * [Learn Java Programming (version 17)](https://www.freecodecamp.org/news/learn-java-programming/)
* Gradle: Learn by using in the project, key concepts linked below.
    * [The Gradle Wrapper](https://docs.gradle.org/current/userguide/gradle_wrapper.html)
* Sprint Boot: [Building an Application with Spring Boot](https://spring.io/guides/gs/spring-boot/)
* We’ll use Sprint Boot as our Java framework. It is the most popular. Some of our clients also use [Dropwizard](https://www.dropwizard.io/en/latest/)
   {{% /expand%}}


### Behaviour Driven Development

BDD style tests are used by applications that are deployed to the core platform as well as testing the core platform itself.

* [Introducing BDD](https://dannorth.net/introducing-bdd/)
* [Learn Cucumber BDD with Java -MasterClass Selenium Framework](https://www.udemy.com/course/cucumber-tutorial/)

### Containers

{{%expand "Materials on Docker and Containers" %}}
* New to Docker?
    * [Docker Quick Start on Cloud Guru](https://acloudguru.com/course/docker-quick-start)
    * [Docker intro on Plurasight](https://www.pluralsight.com/courses/getting-started-docker)
* New to Docker Compose?
    * Understand networking: [Networking in Compose](https://docs.docker.com/compose/networking/)
      {{% /expand%}}

### Local build setup
* Three amigos (Make, Docker, & Compose) [3 Musketeers](https://3musketeers.io/guide/)
    * We’ll use this to run commands with make and Docker



