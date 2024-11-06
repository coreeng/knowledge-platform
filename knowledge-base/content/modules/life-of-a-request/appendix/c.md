+++
title = "C - Using dig"
weight = 3
chapter = false
+++

```
$ dig cecg.io

; <<>> DiG 9.10.6 <<>> cecg.io
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 12195
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 1232
;; QUESTION SECTION:
;cecg.io.			IN	A

;; ANSWER SECTION:
cecg.io.		3523	IN	A	75.2.60.5

;; Query time: 18 msec
;; SERVER: 213.140.209.239#53(213.140.209.239)
;; WHEN: Mon Oct 21 13:39:06 EEST 2024
;; MSG SIZE  rcvd: 52
```

The response of the **dig** command contains several sections:

**HEADER** provides metadata about the query and the response:

- **Status**: Whether the query was successful (e.g., **NOERROR**) or if there was an error (e.g., **NXDOMAIN** for non-existent domain).
- **ID**: A unique identifier for the DNS query.
- **Flags**: Shows various flags related to the DNS query process such as **qr** (query response), **rd** (recursion desired), **ra** (recursion available), etc.
- **Query**: The number of queries issued.
- **Answer**, **Authority**, **Additional**: The count of entries in each of these sections.

**QUESTION SECTION** repeats the query itself, stating what DNS record you asked for. This section will show the domain (cecg.io) and the type of query (e.g., **A**, **MX**, **NS**).

**ANSWER SECTION** is returned only if the domain exists. It contains the actual DNS records (resource records) for the query. It shows the results from the DNS server. Common fields include:

- **Name**: The domain name.
- **TTL**: Time to Live, indicating how long the record is cached.
- **Class**: Almost always **IN** for the Internet.
- **Type**: The type of record (e.g., **A** for IP address, **MX** for mail servers).
- **Data**: The information associated with the record, such as an IP address for an **A** record.

**AUTHORITY SECTION** lists the authoritative name servers for the domain. This section appears if you are querying for a domain that is not in your DNS server’s cache, and it provides the DNS servers that are authoritative for that domain.

**ADDITIONAL SECTION** provides additional information, such as IP addresses for the name servers listed in the Authority Section.

**Query Time** returns the time taken to execute the query, in milliseconds.

**SERVER** is the DNS server that was used to resolve the query

**WHEN** is the date and time when the query was run.

**MSG SIZE** is the size of the DNS message that was sent and received.
