+++
title = "I - Using curl"
weight = 9
chapter = false
+++

Using curl is a great way to debug connectivity issues, especially for checking if a machine can reach a specific endpoint, test server responses, or troubleshoot network connections. Below are some common curl commands and techniques that can help with connectivity debugging.

This is the most basic way to check if you can reach a website or endpoint:

```
$ curl https://cecg.io

Redirecting to https://www.cecg.io/
```

If there is no response, or it hangs, this can indicate network or DNS resolution issues.

Use the **-I** option to fetch only the headers of the response:

```
$ curl -I https://cecg.io

HTTP/2 301
content-type: text/plain; charset=utf-8
date: Thu, 24 Oct 2024 12:45:07 GMT
location: https://www.cecg.io/
server: Netlify
```

The **-v** option provides detailed information about the request and response, including DNS resolution, TLS handshake, and more:

```
$ curl -v https://cecg.io

* Host cecg.io:443 was resolved.
* IPv6: (none)
* IPv4: 75.2.60.5
*   Trying 75.2.60.5:443...
* Connected to cecg.io (75.2.60.5) port 443
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
*  CAfile: /etc/ssl/certs/ca-certificates.crt
*  CApath: /etc/ssl/certs
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_128_GCM_SHA256 / X25519 / id-ecPublicKey
* ALPN: server accepted h2
* Server certificate:
*  subject: CN=cecg.io
*  start date: Oct  1 21:33:44 2024 GMT
*  expire date: Dec 30 21:33:43 2024 GMT
*  subjectAltName: host "cecg.io" matched cert's "cecg.io"
*  issuer: C=US; O=Let's Encrypt; CN=E5
*  SSL certificate verify ok.
*   Certificate level 0: Public key type EC/prime256v1 (256/128 Bits/secBits), signed using ecdsa-with-SHA384
*   Certificate level 1: Public key type EC/secp384r1 (384/192 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type RSA (4096/152 Bits/secBits), signed using sha256WithRSAEncryption
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* using HTTP/2
* [HTTP/2] [1] OPENED stream for https://cecg.io/
* [HTTP/2] [1] [:method: GET]
* [HTTP/2] [1] [:scheme: https]
* [HTTP/2] [1] [:authority: cecg.io]
* [HTTP/2] [1] [:path: /]
* [HTTP/2] [1] [user-agent: curl/8.5.0]
* [HTTP/2] [1] [accept: */*]
> GET / HTTP/2
> Host: cecg.io
> User-Agent: curl/8.5.0
> Accept: */*
>
< HTTP/2 301
< content-type: text/plain; charset=utf-8
< date: Thu, 24 Oct 2024 12:45:54 GMT
< location: https://www.cecg.io/
< server: Netlify
< strict-transport-security: max-age=31536000
< x-nf-request-id: 01JAZ8ZS2QZWP29A321DZ36VD6
< content-length: 35
<
* Connection #0 to host cecg.io left intact
```

You can use curl to test connectivity to a specific port (such as testing an API endpoint):

```
$ curl http://cecg.io:443

Client sent an HTTP request to an HTTPS server.
```

To test if SSL/TLS is properly configured and if there are any certificate issues:

```
$ curl https://expired.badssl.com

curl: (60) SSL certificate problem: certificate has expired
More details here: https://curl.se/docs/sslcerts.html

curl failed to verify the legitimacy of the server and therefore could not establish a secure connection to it. To learn more about this situation and how to fix it, please visit the web page mentioned above.
```

The **--insecure** flag skips certificate validation. If you’re having TLS issues, it may show more details about the failure.

```
$ curl -v https://expired.badssl.com --insecure

* Host expired.badssl.com:443 was resolved.
* IPv6: (none)
* IPv4: 104.154.89.105
*   Trying 104.154.89.105:443...
* Connected to expired.badssl.com (104.154.89.105) port 443
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
* TLSv1.2 (IN), TLS handshake, Server finished (14):
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.2 (OUT), TLS handshake, Finished (20):
* TLSv1.2 (IN), TLS handshake, Finished (20):
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / prime256v1 / rsaEncryption
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: OU=Domain Control Validated; OU=PositiveSSL Wildcard; CN=*.badssl.com
*  start date: Apr  9 00:00:00 2015 GMT
*  expire date: Apr 12 23:59:59 2015 GMT
*  issuer: C=GB; ST=Greater Manchester; L=Salford; O=COMODO CA Limited; CN=COMODO RSA Domain Validation Secure Server CA
*  SSL certificate verify result: unable to get local issuer certificate (20), continuing anyway.
*   Certificate level 0: Public key type RSA (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type RSA (2048/112 Bits/secBits), signed using sha384WithRSAEncryption
*   Certificate level 2: Public key type RSA (4096/152 Bits/secBits), signed using sha384WithRSAEncryption
* using HTTP/1.x
> GET / HTTP/1.1
> Host: expired.badssl.com
> User-Agent: curl/8.5.0
> Accept: */*
```

Use timeouts to prevent curl from hanging indefinitely:

```
$ curl --connect-timeout 1 --max-time 5 http://httpbin.org/delay/10

curl: (28) Operation timed out after 5004 milliseconds with 0 bytes received
```

By using these curl techniques, you can get a lot of information that can help troubleshoot connectivity issues, network configuration, DNS problems, SSL certificate issues, and more.
