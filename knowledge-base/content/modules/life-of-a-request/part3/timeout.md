+++
title = "What is a timeout?"
weight = 4
chapter = false
+++

#### Overview

A network timeout happens when a network request takes too long to complete, causing it to fail. This can happen in various scenarios, such as when trying to load a webpage, send a message, or access an online resource. Network timeouts can occur for several reasons:

- **Slow Network Connection**: If the network is too slow, data cannot be transmitted within the expected timeframe.
- **Server Overload**: If the server handling the request is overloaded, it might not respond in time.
- **High Latency**: Long-distance connections or networks with high latency (delay) can also lead to timeouts.
- **Configuration Issues**: Sometimes, systems have low timeout settings, meaning they wait only a short time before assuming the request has failed.

When a timeout happens, the requesting application (like a browser or an app) will typically display an error message, prompting the user to try again or check their connection. Adjusting timeout settings, using a more reliable connection, or accessing the network resource at a less busy time can sometimes help resolve the issue.

#### Types of timeouts

When sending an HTTP request to a server, different types of timeouts may occur depending on the stage of the request and the network conditions.

| Type | Stage | Description |
| :---- | :---- | :---- |
| Connection  | Establishing connection | Limits the time to establish a connection to the server. |
| Read | Data Receiving | Limits waiting time for data after a request is sent. |
| Write | Data Sending | Limits the time allowed for data to be sent to the server. |
| Idle / Keep-Alive | Connection Idle | Closes connection if there’s no activity. |
| Response | Full Response | Limits total time for receiving a complete response. |
| DNS | DNS Resolution | Limits time allowed for resolving domain name to IP address. |
| TLS | TLS Handshake | Limits the time to establish a TLS handshake. |

#### Troubleshooting

To test timeouts you can use **curl** with the following options:

- **--max-time**: Maximum time allowed for the whole operation (in seconds).  
- **--connect-timeout**: Time limit for the connection phase only (in seconds).

To test from within a Kubernetes pod, you can use an image that includes **curl**, like **busybox**. Start a temporary pod with:

```
$ kubectl run curlpod --rm -i --tty --image=busybox -- /bin/sh
```

This will give you a shell in a pod where you can run **curl** commands.

**Connection Timeout**: To test the connection timeout (e.g., 2 seconds):

```
$ curl --connect-timeout 2 https://cecg.io

curl: (28) Connection timeout after 2000 ms
```

**Total Timeout**: To set a maximum total time of 2 seconds for the request:

```
$ curl --max-time 2 https://httpbin.org/delay/5  

curl: (28) Failed to connect to httpbin.org port 443 after 1519 ms: Operation timed out
```

See **Appendix I** for more information about using **curl**.

#### Further Reading {#further-reading-6}

- [Timeouts - everything curl](https://everything.curl.dev/usingcurl/timeouts.html)