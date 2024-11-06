+++
title = "H - Using telnet"
weight = 8
chapter = false
+++

Telnet is a simple and effective tool for testing TCP connectivity to specific ports on remote hosts.

The basic syntax:

```
telnet <hostname or IP address> <port>
```

To connect to a web server running on cecg.io at port 80:

```
$ telnet cecg.io 80

Trying 75.2.60.5...
Connected to cecg.io.
Escape character is '^]'.
Connection closed by foreign host.
```

If the port is closed or unreachable, you will receive an error message, such as “Connection refused” or “Unable to connect”.

```
$ telnet cecg.io 80

Trying 75.2.60.5 ...
telnet: Unable to connect to remote host: Connection refused
```

Check if the service you are trying to connect to is running on the target host. Make sure you’re trying to connect to the correct port.
