+++
title = "D - Ports"
weight = 4
chapter = false
+++

In TCP/IP networking, certain port numbers are reserved for well-known services or protocols. These reserved ports are defined by the Internet Assigned Numbers Authority (IANA) and are classified into three main ranges:

#### Well-Known Ports (0 to 1023)

These are reserved for system or well-known services and protocols. Only privileged users or system administrators can bind a service to these ports.

For example:
- Port **80**: HTTP (Hypertext Transfer Protocol)
- Port **443**: HTTPS (Secure HTTP)
- Port **21**: FTP (File Transfer Protocol)
- Port **25**: SMTP (Simple Mail Transfer Protocol)
- Port **53**: DNS (Domain Name Service)

#### Registered Ports (1024 to 49151)

These are used for applications that are not typically bundled with the operating system but are still common. These ports are assigned to user applications or processes.

For example:
- Port **3306**: MySQL Database
- Port **8080**: Alternative HTTP Port
- Port **1521**: Oracle Database Default Listener

#### Dynamic/Private Ports (49152 to 65535)

These are also known as ephemeral ports, used for temporary or client-side communications. When an application starts a communication session, the operating system assigns an ephemeral port from this range for the duration of the session.
