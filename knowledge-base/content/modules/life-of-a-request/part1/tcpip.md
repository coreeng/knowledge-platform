+++
title = "What is TCP/IP?"
weight = 3
chapter = false
+++

#### Overview

Once we resolve the address of the domain, we can try to establish a connection.

The Internet Protocol (**IP**) is the primary protocol in the suite of protocols that manage communication across the Internet. It is responsible for addressing and routing packets of data from the source host (device) to the destination host, ensuring that data sent from one computer can reach another on the network, regardless of how many routers or networks lie between them.

Both **TCP** (Transmission Control Protocol) and **UDP** (User Datagram Protocol) are built on top of IP. These protocols are responsible for enabling data transmission between applications on different devices. Although they both work on top of IP, they have distinct differences and are suited for different use cases.

Every device on the Internet has an **IP address**, which is used as an identifier to route packets between devices. This address allows devices to identify each other and facilitates the routing of packets between them.

##### TCP

TCP is a connection-oriented protocol that provides reliable, ordered, and error-checked data transmission between devices.

All TCP connections begin with a three-way handshake. Before the client or the server can exchange any application data, they must agree on starting packet sequence numbers, as well as a number of other connection specific variables, from both sides. The sequence numbers are picked randomly from both sides for security reasons.

![Three-way TCP handshake](/images/loar/1-5.png)
_Figure 1-5. Three-way TCP handshake_


1. **SYN** \- Client picks a random sequence number and sends a SYN packet, which may also include additional TCP flags and options.

2. **SYN ACK** \- Server increments the client’s number by one, picks its own random sequence number, appends its own set of flags and options, and dispatches the response.

3. **ACK** \- Client increments both numbers by one and completes the handshake by dispatching the last ACK packet in the handshake.

Use TCP when reliability is more important than speed, such as when transmitting web pages, emails, or files.

##### UDP

UDP is a lightweight, connectionless protocol that provides fast, but less reliable, data transmission between devices over a network. It is designed for speed and efficiency, making it an ideal choice for applications where performance is critical and where the occasional loss of data packets can be tolerated.

Applications like live video, audio streaming, and gaming use UDP to ensure low latency. It’s acceptable to lose some data packets in these applications, as long as the stream remains continuous.

DNS uses UDP for name resolution queries because it involves simple request-response transactions where speed is more critical than guaranteed delivery. If a DNS query fails, the client can easily resend the request.

##### IP Address

The Internet address is also called an **IP address**. There are 2 versions: **IPv4** and **IPv6**.

IPv4 addresses are 32-bit numbers represented in a dotted decimal format. There are four octets (bytes), each ranging from 0 to 255, separated by periods. For example: **192.168.0.1**.

IPv6 addresses are 128-bit numbers written in hexadecimal format, separated by colons, and grouped into eight segments. Leading zeros in a segment can be omitted, and consecutive zero segments can be replaced by :: (only once per address). For example: **2001:0db8:85a3:0000:0000:8a2e:0370:7334**

**IPv4 addresses are widely used, but IPv6 is becoming more common due to the shortage of IPv4 addresses.**

##### Subnets

Subnets are smaller network segments within a larger network, allowing organisations to divide their IP address space for better traffic management, security, and resource allocation.

A subnet mask works alongside an IP address to define the subnet of the network. It masks the IP address to differentiate between the network and host portions. For instance, a subnet mask of **255.255.255.0** indicates that the first 24 bits (three octets) represent the network portion, leaving the last 8 bits for host addresses within the subnet.

##### CIDR

CIDR (Classless Inter-Domain Routing) notation represents IP addresses and their associated network masks. For example, **192.168.1.0/24** specifies an IP range where **/24** indicates that the first 24 bits are used for the network portion, leaving 8 bits for hosts.

In Kubernetes, CIDR is used primarily for managing IP address allocation within clusters and facilitating networking between pods and services.

##### Ports

IP address represents a host computer that can be running multiple services such as a web server and a mail server. To make a request to a specific service we need to specify its **port**.

In TCP/IP, ports are logical endpoints used to identify specific processes or services on a device. They are essential for network communication, enabling different applications and services to interact over the internet or local networks. Ports are represented by numbers ranging from 0 to 65535, and they are used by both TCP and UDP.

See [Appendix D](../../appendix/d) for more information about port numbers.  

![Connecting to a service by its port](/images/loar/1-6.png) 
_Figure 1-6. Connecting to a service by its port_

When we get a URL like `https://cecg.io` we assume that we will be making a connection to a web server using **https** protocol which is available on port **443** by default.

A service may be running on a different port than it is normally expected, in this case a port can be explicitly overridden in the URL, e.g. `https://cecg.io:8433`.

##### Sockets

A network socket is the representation of an IP address and a port through which a network connection can be activated.

When an application needs to send or receive data, it creates a socket, binds it to a specific port number and IP address, and then uses it to communicate with other sockets on remote machines. In simpler terms, a socket provides a way for two machines to talk to each other over a network.  

![Socket communication](/images/loar/1-7.png)
_Figure 1-7. Socket communication_

The sockets are used on both client and server sides. If a client wants to connect to a web server it sends a request from its own dynamic port (e.g., 50000\) to the server’s IP address using port 80 (HTTP). The server receives the request and knows how to respond to the correct client using the source IP and port (50000 in this example).

#### Troubleshooting

Troubleshooting TCP/IP connections involves identifying issues in network communication, which can be caused by problems in addressing, routing, or protocol configurations.

##### Ping the host

Use the **ping** command to check whether the target host is reachable.

```
$ ping cecg.io

PING cecg.io (75.2.60.5): 56 data bytes
64 bytes from 75.2.60.5: icmp_seq=0 ttl=245 time=48.675 ms
64 bytes from 75.2.60.5: icmp_seq=1 ttl=245 time=48.623 ms
64 bytes from 75.2.60.5: icmp_seq=2 ttl=245 time=44.944 ms
^C
--- cecg.io ping statistics ---
3 packets transmitted, 3 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 44.944/47.414/48.675/1.747 ms
```

If the host is unreachable, you’ll receive a timeout or unreachable error.

**Note:** **ping** uses **ICMP** protocol to communicate with the target host. Unfortunately, some providers block ICMP traffic due to security reasons. It may seem that the host is unreachable, but in fact it is only ICMP traffic that is blocked, but TCP/UDP traffic is allowed.

See [Appendix E](../../appendix/e) for more information about the **ping** command.

##### Traceroute

Use the **traceroute** command to show the path packets take to reach the target and help identify where the connection is breaking.

```
$ traceroute google.com

traceroute to google.com (142.250.201.14), 64 hops max, 40 byte packets
 1  192.168.0.1 (192.168.0.1)  4.561 ms  4.683 ms  4.256 ms
 2  * * *
 3  213.140.209.65 (213.140.209.65)  12.971 ms  13.228 ms  12.991 ms
 4  be4.core-r1.nic-south.cy.cablenet-as.net (213.140.198.56)  47.470 ms  54.168 ms  47.254 ms
 5  hu-0-5-0-0.edge-r.mar.fr.cablenet-as.net (213.140.198.140)  48.307 ms  48.077 ms  47.478 ms
 6  cnt-google-pni.edge-r.mar.fr.cablenet-as.net (91.184.192.134)  48.478 ms  46.185 ms  47.606 ms
 7  192.178.105.173 (192.178.105.173)  48.438 ms  49.361 ms
    192.178.105.211 (192.178.105.211)  50.252 ms
 8  216.239.42.99 (216.239.42.99)  52.241 ms
    216.239.42.135 (216.239.42.135)  52.126 ms
    216.239.42.99 (216.239.42.99)  51.892 ms
 9  mrs08s19-in-f14.1e100.net (142.250.201.14)  55.857 ms  49.231 ms  48.542 ms
```

The output shows the devices the packet passes through to reach the target host. The number of the devices is also called the number of hops.

**Note:** **traceroute** uses **ICMP** protocol to communicate with the devices. Unfortunately, some providers block ICMP traffic due to security reasons. Even though you see asterisks **\* \* \*** for certain hops, the **traceroute** can still reach the destination. The presence of asterisks at certain points simply means that those particular devices didn’t respond to the probes, but the overall route was successful.

```
$ traceroute cecg.io

traceroute to cecg.io (75.2.60.5), 6 hops max, 60 byte packets
 1  172.17.0.1 (172.17.0.1)  0.466 ms  0.013 ms  0.007 ms
 2  * * *
 3  * * *
 4  * * *
 5  * * *
 6  * * *
```

The target host **cecg.io** (75.2.60.5) is reachable with **ping**, but **traceroute** is not able to trace the path, returning only asterisks **\* \* \***.

In case when ICMP traffic is blocked we can use **TCP SYN** packets instead. To perform a traceroute using TCP packets instead of the default UDP or ICMP, you can use the **\-T** option with **traceroute**.

```
$ traceroute -T cecg.io

traceroute to cecg.io (75.2.60.5), 30 hops max, 60 byte packets
 1  172.17.0.1 (172.17.0.1)  0.280 ms  0.006 ms *
 2  acd89244c803f7181.awsglobalaccelerator.com (75.2.60.5)  61.390 ms  60.019 ms  51.133 ms
```

See [Appendix F](../../appendix/f) for more information about the **traceroute** command.

##### Check Ports

If you’re trying to connect to a specific service, verify whether the port is open on the target host.

Use **telnet** to test if a port is open:

```
$ telnet cecg.io 443

Trying 75.2.60.5...
Connected to cecg.io.
Escape character is '^]'.
```

We successfully connected to port 443 with **telnet**. Below is the example of a closed port:

```
$ telnet cecg.io 444

Trying 75.2.60.5...
telnet: Unable to connect to remote host: Connection refused
```

See [Appendix H](../../appendix/h) for more information about the **telnet** command.

Alternatively, use **nc** (Netcat):

```
$ nc -zv cecg.io 443

cecg.io [75.2.60.5] 443 (https) open
```

We successfully connected to port 443 with **nc**. Below is the example of a closed port:

```
$ nc -zv cecg.io 444

cecg.io [75.2.60.5] 444 (snpp) : Connection refused
```

See [Appendix G](../../appendix/g) for more information about the **nc** command.

Another option is **nmap** (Network Mapper), a powerful network scanning tool.

See [Appendix J](../../appendix/j) for more information about the **nmap** command.

##### Check Firewalls

Make sure no local firewall (like iptables on Linux) is blocking traffic to or from the target host.

Check with your network administrator if there are network firewalls or routers blocking traffic.

For cloud instances, verify that the correct security groups and firewall rules are set to allow inbound and outbound traffic.

Try to access the target from another host or network to determine whether the issue is specific to your machine or network.

#### Further Reading

- [RFC 791 \- Internet Protocol](https://datatracker.ietf.org/doc/html/rfc791)
- [RFC 9293 \- Transmission Control Protocol (TCP)](https://datatracker.ietf.org/doc/html/rfc9293)
- [RFC 792 \- Internet Control Message Protocol](https://datatracker.ietf.org/doc/html/rfc792)
