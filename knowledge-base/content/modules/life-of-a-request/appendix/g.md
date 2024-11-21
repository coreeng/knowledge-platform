+++
title = "G - Using netcat"
weight = 7
chapter = false
+++

Netcat, often referred to as **nc**, is a powerful networking utility that can read and write data across network connections using TCP or UDP. It’s commonly used for debugging and testing network connections, transferring files, and setting up network communication between systems.

The basic syntax: 

```
nc [flags] [hostname] [port]
```

There are several useful flags:

| Flag | Description |
| :---- | :---- |
| \-l | Listen mode. This makes nc listen for incoming connections on the specified port. Useful for setting up a server. nc \-l \-p 12345 |
| \-p | Specify the port to use. nc \-l \-p 12345 |
| \-u | Use UDP instead of TCP. By default, Netcat uses TCP. nc \-u \-l \-p 12345 |
| \-v | Verbose mode. This option provides more detailed output about the connection status. nc \-v cecg.io 80 |
| \-z | Zero-I/O mode. This flag is used for scanning. It does not send any data, just checks if the port is open. nc \-z \-v cecg.io 1-1000 |
| \-w | Specify a timeout for the connection (in seconds). This is useful to prevent hanging if the connection is not established.  nc \-w 5 cecg.io 80 |

The combination of **-z** and **-v** flags is used for port scanning.

```
$ nc -zv cecg.io 443

Connection to cecg.io port 443 [tcp/https] succeeded!
```
