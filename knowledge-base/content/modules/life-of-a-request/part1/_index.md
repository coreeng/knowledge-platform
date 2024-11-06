+++
title = "Part 1 - Sending a request from the outside"
weight = 1
chapter = false
+++

- [What is a URL?](#url)
- [What is a DNS?](#dns)
- [What is TCP/IP?](#tcpip)
- [What is HTTP(S)?](#https)
- [What is an Ingress?](#ingress)
- [What is a Service?](#service)
- [What is inside a Pod?](#pod)

---

## What is a URL? {#url}

#### Overview

A URL (Uniform Resource Locator) is a specific type of address used to access resources on the internet. It serves as the “web address” that identifies the location of a resource, such as a web page, image, video, or file, on a network, typically the World Wide Web.

Let’s assume that the Web App is available at `https://cecg.io/webapp` and we are going to open this URL in a web browser.

![The structure of a URL](/images/loar/1-1.png)
_Figure 1-1. The structure of a URL_

A URL consists of several components, such as protocol, domain name, path, parameters and fragment (anchor).

In our example we have:

- protocol \-  **https**
- domain name \- **cecg.io**
- path \- **/webapp**

So, we are going to establish a connection with **cecg.io** using **https** protocol.

#### Further Reading

- [RFC 1738 \- Uniform Resource Locators (URL)](https://datatracker.ietf.org/doc/html/rfc1738)
- [RFC 3986 \- Uniform Resource Identifier (URI): Generic Syntax](https://datatracker.ietf.org/doc/html/rfc3986) 

## What is a DNS? {#dns}

#### Overview

The first thing worth mentioning is that we are addressing the Web App by its human-readable domain name. This is what humans usually work with, but the computer system needs to translate the domain name into the Internet address and route the request accordingly.

DNS, or Domain Name System, is a crucial component of the Internet, translating human-readable addresses into machine-readable Internet addresses. There are DNS servers that you can query to look up the addresses, e.g. **8.8.8.8** is a public DNS server provided by Google.

![Example of the Internet address resolution for the domain](/images/loar/1-2.png)
_Figure 1-2. Example of the Internet address resolution for the domain_

Check the [Appendix B](../appendix/#b) section for the detailed description of DNS server types.

When a user enters **cecg.io** into the web browser, it first checks whether this domain was resolved recently and if the response is cached locally which it can reuse. If not, then the request is sent to a DNS server provided by an Internet service provider (ISP) or a public DNS server. DNS servers do recursive resolution by querying root and top-level domain (TLD) servers.

If the response is cached and returned by any intermediate DNS server then we call it a “non-authoritative DNS response”. Authoritative response comes from a DNS server that is responsible for the domain records.

DNS records have **TTL** (Time to Live) to control how long information about a DNS query is cached by DNS resolvers and other systems before they need to re-query the authoritative DNS server for updated information.

If a DNS record has a TTL of 3600 seconds (1 hour), any DNS resolver that queries the record can cache the response for up to 1 hour. After this period, the resolver will request the updated information from the authoritative DNS server.

TTL is defined on an authoritative DNS server, the choice of the value affects performance and flexibility.

DNS database contain several record types, e.g.:

- **A** \- contains an IPv4 address
- **AAAA** \- contains an IPv6 address
- **CNAME** \- defines an alias
- **MX** \- defines a mail server
- **TXT** \- stores text information, can be used for domain verification.

#### Debugging Tools

Normally, a human does not need to interact with DNS directly as Internet address translation is handled by the operating system.

If you need to debug the responses from the DNS servers you can use one of the following Unix commands: **nslookup**, **dig** or **whois**.

In the following example **nslookup** is used to resolve **cecg.io** domain into address **75.2.60.5**.

```
$ nslookup cecg.io

Server:	213.140.209.239
Address:	213.140.209.239#53

Non-authoritative answer:
Name:	cecg.io
Address: 75.2.60.5
```

Another command is **dig**, it is more capable than **nslookup**, it provides many different flags for DNS troubleshooting. For example, you can get **TXT** records for **cecg.io** domain.

```
$ dig cecg.io TXT

. . .
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 9495
;; flags: qr rd ra; QUERY: 1, ANSWER: 5, AUTHORITY: 0, ADDITIONAL: 1

;; ANSWER SECTION:
cecg.io.		600	IN	TXT	"v=spf1 include:mail.zendesk.com ?all"
cecg.io.		600	IN	TXT	"v=spf1 a mx include:websitewelcome.com ~all"
cecg.io.		600	IN	TXT	"google-site-verification=WLW7m6cd2CzreF2J6dueADV5wxMcwf91MjjlzErEiog"
cecg.io.		600	IN	TXT	"atlassian-domain-verification=oa6cS5/QTuwj4Nv17iINtkulDZMFINlEHOEU1U/ucaXsyRpR477qQD2I575ZQf5z"
cecg.io.		600	IN	TXT	"MS=ms68596058"
```

Trying to resolve **cecgfoobar.io** will return the empty result:

```
$ dig cecgfoobar.io TXT

. . .
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NXDOMAIN, id: 42806
;; flags: qr rd ra; QUERY: 1, ANSWER: 0, AUTHORITY: 1, ADDITIONAL: 1
```

Check [Appendix C](../appendix/#c) section for more details about **dig** output.

For a simple DNS lookup you can use **\+short** flag to ignore the details.

```
$ dig cecg.io +short

75.2.60.5

$ dig microsoft.com +short

20.112.250.133
20.231.239.246
20.236.44.162
20.70.246.20
20.76.201.171
```

Yes, there can be several records of the same type. Having several **A** records for the domain is typically used for round-robin load balancing.

The Internet address that you get from the DNS is the destination where the client application will be sending the request.

It is also possible to check the owner of the Internet address by doing DNS **reverse** lookup. We can use either **whois** or **dig \-x** command on Unix, e.g.:

```
$ whois 75.2.60.5

NetRange:       75.2.0.0 - 75.2.191.255
CIDR:           75.2.0.0/17, 75.2.128.0/18
NetName:        AMAZO-4
NetHandle:      NET-75-2-0-0-1
Parent:         NET75 (NET-75-0-0-0-0)
NetType:        Direct Allocation
OriginAS:       AS16509
Organization:   Amazon.com, Inc. (AMAZO-4)
RegDate:        2018-01-10
Updated:        2018-01-11
Ref:            https://rdap.arin.net/registry/ip/75.2.0.0
```

Based on the output we can assume that this address belongs to Amazon.

#### ExternalDNS

[ExternalDNS](https://github.com/kubernetes-sigs/external-dns) is a Kubernetes add-on that automates the management of DNS records for Kubernetes services. It allows you to dynamically create DNS records for services or Ingress resources in Kubernetes, making it easier to manage external access to applications running in your cluster.

![Example of external-dns annotation for Ingress](/images/loar/1-3.png)
_Figure 1-3. Example of external-dns annotation for Ingress_

ExternalDNS monitors your Kubernetes cluster for changes to Ingresses and Services (of **type=LoadBalancer**). When an Ingress or Service is created, updated, or deleted, ExternalDNS uses the respective DNS cloud provider’s API to create, update, or remove DNS records.

ExternalDNS can be configured using annotations on Kubernetes resources to specify how DNS records should be managed.   
In the below example **external-dns.alpha.kubernetes.io/hostname** specifies the hostname for a given service:

```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/hostname: myapp.cecg.io
    external-dns.alpha.kubernetes.io/target: cecg.io
  name: cecg-training-platform
spec:
  ingressClassName: traefik
  rules:
    - host: myapp.cecg.io
      http:
        paths:
          - backend:
              service:
                name: myapp
                port:
                  number: 80
            path: /
            pathType: ImplementationSpecific
```

ExternalDNS will request a DNS provider (e.g. Cloud DNS) to create a record for this domain.

![Example of a type A record created in CloudDNS for myapp.cecg.io](/images/loar/1-4.png)
_Figure 1-4. Example of a type A record created in CloudDNS for myapp.cecg.io_

ExternalDNS simplifies the management of DNS records for Kubernetes services, automating the creation and updating of DNS entries as your infrastructure evolves. It’s especially useful for dynamic environments where services and their endpoints may change frequently.

#### Troubleshooting

Working with DNS you can face several issues that may cause the application request to fail.

##### Missing or invalid records

A new application was deployed, but a DNS record was not created or updated.

Use **nslookup** or **dig** to debug DNS resolution and identify where the query is failing.

```
$ nslookup unknown.cecg.io

** server can't find unknown.cecg.io: NXDOMAIN

$ dig unknown.cecg.io +short

(empty output)
```

If the output is missing the record or it is invalid you need to check the configuration of the authoritative DNS server.   
If an authoritative DNS server does not have the correct records, you need to debug a service responsible for DNS change propagation. In the case of our reference cluster it can be an issue with ExternalDNS failing to create or update DNS records due to invalid Ingress configuration.

If the records are correct on authoritative DNS, then the outdated values may be cached on intermediate servers.

##### Outdated records

Propagating changes through the internet takes time. Depending on Time-to-Live (TTL) settings and caching behaviour of DNS servers the clients may get outdated responses.

You can use **dig** to check current values for the DNS records. For example, the following command will show **A** records containing Internet addresses for the domain:

The answer section of the **dig cecg.io A** command contains a TTL value:

```
;; ANSWER SECTION:
cecg.io.		3600	IN	A	75.2.60.5
```

If a DNS **A** record for **cecg.io** has a TTL of **3600** seconds (1 hour), it means that once a DNS resolver has cached the IP address, it will continue to use that cached information for 1 hour before checking the authoritative DNS server again.

By default, **dig** is sending queries to local DNS servers configured in your OS:

```
$ dig cecg.io

;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 54490
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; ANSWER SECTION:
cecg.io.		1473	IN	A	75.2.60.5

;; Query time: 12 msec
;; SERVER: 213.140.209.239#53(213.140.209.239)
;; WHEN: Mon Oct 21 16:33:03 EEST 2024
;; MSG SIZE  rcvd: 52
```

You can use **dig** to query a specific DNS server, not only those configured in your system.

The following command will make a request to the public DNS **8.8.8.8**:

```
$ dig @8.8.8.8 cecg.io

;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 19694
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; ANSWER SECTION:
cecg.io.		3600	IN	A	75.2.60.5

;; Query time: 303 msec
;; SERVER: 8.8.8.8#53(8.8.8.8)
;; WHEN: Mon Oct 21 16:33:43 EEST 2024
;; MSG SIZE  rcvd: 52
```

We can query the authoritative DNS server for **cecg.io** which is **ns74.domaincontrol.com**:

```
$ dig @ns74.domaincontrol.com cecg.io

;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 54473
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 2, ADDITIONAL: 1

;; ANSWER SECTION:
cecg.io.		3600	IN	A	75.2.60.5

;; AUTHORITY SECTION:
cecg.io.		3600	IN	NS	ns74.domaincontrol.com.
cecg.io.		3600	IN	NS	ns73.domaincontrol.com.

;; Query time: 67 msec
;; SERVER: 173.201.74.47#53(173.201.74.47)
;; WHEN: Mon Oct 21 16:36:00 EEST 2024
;; MSG SIZE  rcvd: 107
```

You can compare the outputs of different DNS servers to find discrepancies.  
If you update a DNS record (like changing an IP address), the new information won’t be seen by users until the TTL expires for the old record in their cache. If the TTL is low (e.g., 60 seconds), it means that the record is refreshed more frequently. This is useful if the IP address of a service is likely to change soon.

A higher TTL (e.g., 86400 seconds, which is 24 hours) means that DNS servers and clients can cache the record for a longer time. This reduces the number of lookups to authoritative DNS servers but may delay updates when changes are made.

##### Slow queries

A DNS query can be slow due to several reasons. Depending on the DNS server location, network latency can add significant delays if it is far from the client. If the network path is congested, packet delays or drops can increase the DNS lookup time.

The server can become slow if it is overloaded with lots of requests. DNS servers typically cache records to respond faster. If the record is not cached and needs to be fetched, this can slow down the response.

Higher TTLs reduce the load on DNS servers and increase performance for end-users by avoiding unnecessary DNS queries. Lower TTLs provide flexibility during DNS record changes, ensuring that updates propagate faster.

Another reason is recursive lookups. If a server needs to perform recursive lookups (querying other DNS servers on the internet), it can take additional time.

**dig \+trace** is a useful command in DNS troubleshooting as it allows you to see the path taken by a query as it traverses the DNS hierarchy, starting from the root servers.

```
$ dig cecg.io +trace

; <<>> DiG 9.18.28-0ubuntu0.24.04.1-Ubuntu <<>> cecg.io +trace
;; global options: +cmd

.			4502	IN	NS	a.root-servers.net.
.			4502	IN	NS	b.root-servers.net.
.			4502	IN	NS	c.root-servers.net.
.			4502	IN	NS	d.root-servers.net.
.			4502	IN	NS	e.root-servers.net.
.			4502	IN	NS	f.root-servers.net.
.			4502	IN	NS	g.root-servers.net.
.			4502	IN	NS	h.root-servers.net.
.			4502	IN	NS	i.root-servers.net.
.			4502	IN	NS	j.root-servers.net.
.			4502	IN	NS	k.root-servers.net.
.			4502	IN	NS	l.root-servers.net.
.			4502	IN	NS	m.root-servers.net.
;; Received 420 bytes from 192.168.65.7#53(192.168.65.7) in 2 ms

io.			172800	IN	NS	a0.nic.io.
io.			172800	IN	NS	a2.nic.io.
io.			172800	IN	NS	b0.nic.io.
io.			172800	IN	NS	c0.nic.io.
io.			86400		IN	DS	57355 8 2 
. . .
;; Received 619 bytes from 199.7.91.13#53(d.root-servers.net) in 49 ms

cecg.io.		3600	IN	NS	ns73.domaincontrol.com.
cecg.io.		3600	IN	NS	ns74.domaincontrol.com.
. . .
;; Received 576 bytes from 65.22.162.17#53(c0.nic.io) in 70 ms

cecg.io.		3600	IN	A	75.2.60.5
cecg.io.		3600	IN	NS	ns74.domaincontrol.com.
cecg.io.		3600	IN	NS	ns73.domaincontrol.com.
;; Received 107 bytes from 173.201.74.47#53(ns74.domaincontrol.com) in 65 ms
```

The output shows the recursive DNS resolution for the **cecg.io** domain if the record is not cached by the local DNS server **192.168.65.7**.

Firstly, **d.root-servers.net** was chosen from the list of **root** servers. The root server was asked for a list of servers responsible for **.io** **TLD**.

Then **c0.nic.io** was queried to provide a list of nameservers for **cecg.io** domain.

Finally, **ns74.domaincontrol.com** was asked to return the **A** record for the domain.

In total the query took **186 ms**.

The same query sent to a local caching DNS server takes only several milliseconds:

```
$ dig @192.168.65.7 cecg.io

;; Query time: 2 msec
;; SERVER: 192.168.65.7#53(192.168.65.7) (UDP)
;; WHEN: Mon Oct 21 13:57:08 UTC 2024
;; MSG SIZE  rcvd: 48
```

##### Server is unavailable

Another common issue is that clients can’t connect to their DNS server. When this happens you can still access the hosts by its Internet address. It can be either a configuration issue or your ISP can be experiencing an outage.

Check **/etc/resolv.conf** for the DNS servers configured in your system:

```
$ cat /etc/resolv.conf

nameserver 213.140.209.239
nameserver 213.140.213.232
```

In the example above, you can see 2 nameservers provided by an ISP.

You can use **ping** to check if a server is reachable:

```
$ ping 213.140.209.239

PING 213.140.209.239 (213.140.209.239): 56 data bytes
64 bytes from 213.140.209.239: icmp_seq=0 ttl=61 time=12.354 ms
64 bytes from 213.140.209.239: icmp_seq=1 ttl=61 time=15.396 ms
64 bytes from 213.140.209.239: icmp_seq=2 ttl=61 time=12.329 ms
64 bytes from 213.140.209.239: icmp_seq=3 ttl=61 time=14.652 ms
^C
--- 213.140.209.239 ping statistics ---
4 packets transmitted, 4 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 12.329/13.683/15.396/1.367 ms
```

All packets received, nothing lost, the ping was successful.

In the below example the server is not reachable:

```
$ ping 213.140.209.240

PING 213.140.209.240 (213.140.209.240): 56 data bytes
Request timeout for icmp_seq 0
Request timeout for icmp_seq 1
Request timeout for icmp_seq 2
```

If the server is unreachable, you need to debug further to understand the reason for packet loss.

See [Appendix E](../appendix/#e) for more information about the **ping** command.  
A typical DNS server is running on port **53**, your firewall rules should allow outgoing connections to port **53**. You can use **nc** to check connectivity:

```
$ nc -vz 8.8.8.8 53

Connection to 8.8.8.8 port 53 [tcp/domain] succeeded!
```

In case of a closed port, the command will hang and the request will time out:

```
$ nc -vz 8.8.8.8 54

nc: connectx to 8.8.8.8 port 54 (tcp) failed: Operation timed out
```


See [Appendix G](../appendix/#g) for more information about the **nc** command.

##### Local overrides

The **/etc/hosts** file is a plain text file used by operating systems to map hostnames to Internet addresses. It serves as a local DNS-like system, allowing the computer to quickly resolve hostnames without needing to query external DNS servers.

This is useful in testing websites before DNS entries are live, but you may experience unexpected behaviour if you are not aware of it.

```
$ cat /etc/hosts

127.0.0.1	localhost
127.0.0.1	cecg.io
255.255.255.255	broadcasthost
::1             localhost
```

You should be aware that **nslookup** and **dig** tools ignore the **/etc/hosts** file as they query DNS servers directly. We can use **ping** to see that **cecg.io** is being actually resolved as **127.0.0.1**.

```
$ ping cecg.io

PING cecg.io (127.0.0.1): 56 data bytes
64 bytes from 127.0.0.1: icmp_seq=0 ttl=64 time=0.376 ms
64 bytes from 127.0.0.1: icmp_seq=1 ttl=64 time=0.133 ms
^C--- cecg.io ping statistics ---
2 packets transmitted, 2 packets received, 0% packet loss
round-trip min/avg/max/stddev = 0.065/0.178/0.376/0.118 ms
```

Double check the records in **/etc/hosts** to make sure they are correct.

#### Further Reading

- [RFC 1034 \- Domain names \- concepts and facilities](https://datatracker.ietf.org/doc/html/rfc1034)
- [RFC 1035 \- Domain names \- implementation and specification](https://datatracker.ietf.org/doc/html/rfc1035)

### What is TCP/IP? {#tcpip}

#### Overview

Once we resolve the address of the domain, we can try to establish a connection.

The Internet Protocol (**IP**) is the primary protocol in the suite of protocols that manage communication across the Internet. It is responsible for addressing and routing packets of data from the source host (device) to the destination host, ensuring that data sent from one computer can reach another on the network, regardless of how many routers or networks lie between them.

Both **TCP** (Transmission Control Protocol) and **UDP** (User Datagram Protocol) are built on top of IP. These protocols are responsible for enabling data transmission between applications on different devices. Although they both work on top of IP, they have distinct differences and are suited for different use cases.

Every device on the Internet has an **IP address**, which is used as an identifier to route packets between devices. This address allows devices to identify each other and facilitates the routing of packets between them.

**TCP**

TCP is a connection-oriented protocol that provides reliable, ordered, and error-checked data transmission between devices.

All TCP connections begin with a three-way handshake. Before the client or the server can exchange any application data, they must agree on starting packet sequence numbers, as well as a number of other connection specific variables, from both sides. The sequence numbers are picked randomly from both sides for security reasons.

![Three-way TCP handshake](/images/loar/1-5.png)
_Figure 1-5. Three-way TCP handshake_


1. **SYN** \- Client picks a random sequence number and sends a SYN packet, which may also include additional TCP flags and options.

2. **SYN ACK** \- Server increments the client’s number by one, picks its own random sequence number, appends its own set of flags and options, and dispatches the response.

3. **ACK** \- Client increments both numbers by one and completes the handshake by dispatching the last ACK packet in the handshake.

Use TCP when reliability is more important than speed, such as when transmitting web pages, emails, or files.

**UDP**

UDP is a lightweight, connectionless protocol that provides fast, but less reliable, data transmission between devices over a network. It is designed for speed and efficiency, making it an ideal choice for applications where performance is critical and where the occasional loss of data packets can be tolerated.

Applications like live video, audio streaming, and gaming use UDP to ensure low latency. It’s acceptable to lose some data packets in these applications, as long as the stream remains continuous.

DNS uses UDP for name resolution queries because it involves simple request-response transactions where speed is more critical than guaranteed delivery. If a DNS query fails, the client can easily resend the request.

**IP Address**

The Internet address is also called an **IP address**. There are 2 versions: **IPv4** and **IPv6**.

IPv4 addresses are 32-bit numbers represented in a dotted decimal format. There are four octets (bytes), each ranging from 0 to 255, separated by periods. For example: **192.168.0.1**.

IPv6 addresses are 128-bit numbers written in hexadecimal format, separated by colons, and grouped into eight segments. Leading zeros in a segment can be omitted, and consecutive zero segments can be replaced by :: (only once per address). For example: **2001:0db8:85a3:0000:0000:8a2e:0370:7334**

**IPv4 addresses are widely used, but IPv6 is becoming more common due to the shortage of IPv4 addresses.**

**Subnets**

Subnets are smaller network segments within a larger network, allowing organisations to divide their IP address space for better traffic management, security, and resource allocation.

A subnet mask works alongside an IP address to define the subnet of the network. It masks the IP address to differentiate between the network and host portions. For instance, a subnet mask of **255.255.255.0** indicates that the first 24 bits (three octets) represent the network portion, leaving the last 8 bits for host addresses within the subnet.

**CIDR**

CIDR (Classless Inter-Domain Routing) notation represents IP addresses and their associated network masks. For example, **192.168.1.0/24** specifies an IP range where **/24** indicates that the first 24 bits are used for the network portion, leaving 8 bits for hosts.

In Kubernetes, CIDR is used primarily for managing IP address allocation within clusters and facilitating networking between pods and services.

**Ports**

IP address represents a host computer that can be running multiple services such as a web server and a mail server. To make a request to a specific service we need to specify its **port**.

In TCP/IP, ports are logical endpoints used to identify specific processes or services on a device. They are essential for network communication, enabling different applications and services to interact over the internet or local networks. Ports are represented by numbers ranging from 0 to 65535, and they are used by both TCP and UDP.

See [Appendix D](../appendix/#d) for more information about port numbers.  

![Connecting to a service by its port](/images/loar/1-6.png) 
_Figure 1-6. Connecting to a service by its port_

When we get a URL like `https://cecg.io` we assume that we will be making a connection to a web server using **https** protocol which is available on port **443** by default.

A service may be running on a different port than it is normally expected, in this case a port can be explicitly overridden in the URL, e.g. `https://cecg.io:8433`.

**Socket**

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

See [Appendix E](../appendix/#e) for more information about the **ping** command.

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

See [Appendix F](../appendix/#f) for more information about the **traceroute** command.

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

See **Appendix H** for more information about the **telnet** command.

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

See [Appendix G](../appendix/#g) for more information about the **nc** command.

Another option is **nmap** (Network Mapper), a powerful network scanning tool.

See [Appendix J](../appendix/#j) for more information about the **nmap** command.

##### Check Firewalls

Make sure no local firewall (like iptables on Linux) is blocking traffic to or from the target host.

Check with your network administrator if there are network firewalls or routers blocking traffic.

For cloud instances, verify that the correct security groups and firewall rules are set to allow inbound and outbound traffic.

Try to access the target from another host or network to determine whether the issue is specific to your machine or network.

#### Further Reading

- [RFC 791 \- Internet Protocol](https://datatracker.ietf.org/doc/html/rfc791)
- [RFC 9293 \- Transmission Control Protocol (TCP)](https://datatracker.ietf.org/doc/html/rfc9293)
- [RFC 792 \- Internet Control Message Protocol](https://datatracker.ietf.org/doc/html/rfc792)


## What is HTTP(S)? {#https}

#### Overview

A typical web server can accept connections using HTTP and/or HTTPS protocols.

HTTP stands for Hypertext Transfer Protocol. It is a protocol used for transmitting hypertext (e.g., web pages) over the internet. HTTP is the foundation of data communication on the World Wide Web, enabling web browsers and servers to exchange information and load resources such as HTML pages, images, videos, and more.






## What is an Ingress? {#ingress}



## What is a Service? {#service}



## What is inside a Pod? {#pod}


