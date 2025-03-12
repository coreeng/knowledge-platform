+++
title = "B - DNS Server Types"
weight = 2
chapter = false
+++

| Type | Description |
| :---- | :---- |
| Forwarding | Acts as a middleman that forwards DNS queries to an external DNS server, usually a recursive DNS server. This type of server does not resolve queries itself but forwards the queries it receives to another DNS server. It’s often used within enterprise networks to centralise DNS requests. |
| Recursive | It receives queries and performs the lookup process, retrieving the requested information from other DNS servers. If the recursive server doesn’t have the answer cached, it queries other DNS servers (root, TLD, authoritative) on behalf of the client to find the IP address.   It caches the responses for a set time (TTL) to speed up future lookups.  |
| Root | The top-level DNS server that handles requests regarding TLDs (Top-Level Domains). When a recursive DNS server doesn’t know how to resolve a query, it first contacts a root DNS server. The root server directs the resolver to the appropriate TLD server (like .com, .org, etc.). There are 13 sets of root DNS servers, distributed globally, to ensure reliability and quick response times. |
| TLD (top-level domain) | Handles requests for domains within a specific top-level domain (like .com, .net, .org). The TLD DNS server directs the recursive server to the authoritative DNS server for the specific domain being queried. For example, the .com TLD DNS server would provide the authoritative DNS server details for example.com. |
| Authoritative | Provides the actual answer to DNS queries based on the domain names it is responsible for. It holds and provides the authoritative IP address for a domain. For example, the authoritative DNS server for **cecg.io** would provide the IP address for **cecg.io** when asked. |
| Caching | Primarily stores the results of previous queries but does not hold any authoritative data. It speeds up DNS resolution by caching the DNS responses for future requests. These servers improve performance and reduce the load on authoritative servers. |
