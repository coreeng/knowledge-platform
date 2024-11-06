+++
title = "What is a DNS?"
weight = 2
chapter = false
+++

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

Double-check the records in **/etc/hosts** to make sure they are correct.

#### Further Reading

- [RFC 1034 \- Domain names \- concepts and facilities](https://datatracker.ietf.org/doc/html/rfc1034)
- [RFC 1035 \- Domain names \- implementation and specification](https://datatracker.ietf.org/doc/html/rfc1035)
