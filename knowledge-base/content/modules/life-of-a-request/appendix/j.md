+++
title = "J - Using nmap"
weight = 10
chapter = false
+++

Nmap (Network Mapper) is a powerful open-source network scanning tool used for network discovery and security auditing. Nmap works by sending specially crafted packets to a target system or network and analysing the responses.

You can use **nmap** to scan for open ports in the target host:

```
$ nmap cecg.io

Starting Nmap 7.94SVN ( https://nmap.org ) at 2024-10-22 12:15 EEST
Nmap scan report for cecg.io (75.2.60.5)
Host is up (0.011s latency).
rDNS record for 75.2.60.5: acd89244c803f7181.awsglobalaccelerator.com
Not shown: 998 filtered tcp ports (no-response)
PORT    STATE SERVICE
80/tcp  open  http
443/tcp open  https

Nmap done: 1 IP address (1 host up) scanned in 4.38 seconds
```

The output shows that 2 ports are open on the host: **80** and **443**.
