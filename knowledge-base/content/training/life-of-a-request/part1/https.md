+++
title = "What is HTTP(S)?"
weight = 4
chapter = false
+++

#### Overview

A typical web server can accept connections using HTTP and/or HTTPS protocols.

HTTP stands for Hypertext Transfer Protocol. It is a protocol used for transmitting hypertext (e.g., web pages) over the internet. HTTP is the foundation of data communication on the World Wide Web, enabling web browsers and servers to exchange information and load resources such as HTML pages, images, videos, and more.

![Sending HTTP request to Web server](/images/loar/1-8.png)
_Figure 1-8. Sending HTTP request to Web server_

HTTP operates using a request-response model, where a client (typically a web browser) sends a request to a web server, and the server responds with the requested resource (e.g., a web page or file).

HTTP has evolved through various versions, each improving performance, security, and capabilities. Each version of HTTP is backward compatible with previous ones, meaning a server supporting HTTP/2 or HTTP/3 can still handle HTTP/1.x requests.

For more details about HTTP versions refer to:

- [RFC 1945 \- Hypertext Transfer Protocol \-- HTTP/1.0](https://datatracker.ietf.org/doc/html/rfc1945)
- [RFC 2616 \- Hypertext Transfer Protocol \-- HTTP/1.1](https://datatracker.ietf.org/doc/html/rfc2616)
- [RFC 7540 \- Hypertext Transfer Protocol Version 2 (HTTP/2)](https://datatracker.ietf.org/doc/html/rfc7540)
- [RFC 9000 \- QUIC: A UDP-Based Multiplexed and Secure Transport](https://datatracker.ietf.org/doc/html/rfc9000) 

##### HTTPS

HTTPS is a secured version of HTTP and is more preferable when exposing public endpoints.

HTTPS uses Transport Layer Security (TLS) protocol to encrypt the data exchanged between a web client and a web server. This encryption prevents attackers from intercepting sensitive information, such as login credentials, credit card details, or other personal data.

HTTPS ensures that the client (browser) is communicating with the intended server and not an impostor through a process called certificate validation. A digital certificate is issued to the website by a trusted certificate authority (CA), and the browser checks the validity of this certificate before establishing a connection.

#### SSL / TLS

SSL (Secure Sockets Layer) and TLS (Transport Layer Security) are both cryptographic protocols designed to provide secure communication over a computer network.

SSL was developed by Netscape in the mid-1990s to encrypt data transmitted over the internet, ensuring that sensitive information (like credit card numbers, usernames, and passwords) cannot be easily intercepted or read by unauthorised parties. SSL uses digital certificates to verify the identity of the parties involved in the communication. This helps ensure that users are communicating with the legitimate website and not an imposter. SSL ensures that the data sent and received during a session remains intact and has not been altered during transmission. It uses hashing algorithms to verify data integrity.

Due to the security flaws in SSL, it has been deprecated in favour of TLS. TLS is a more secure and efficient protocol that builds upon the concepts of SSL but with significant improvements in security and performance. TLS addresses vulnerabilities present in SSL, such as weak cipher suites, message authentication flaws, and other cryptographic weaknesses. SSL has been considered deprecated and insecure for many years. Modern systems no longer support SSL, favouring only the more secure TLS protocols.

Despite the differences, many people still refer to “SSL/TLS certificates” or simply “SSL” when discussing secure connections. This is mainly due to legacy terminology and familiarity, but technically, TLS is what is used in modern secure communication.

##### TLS Handshake

Before the client and the server can begin exchanging application data over Transport Layer Security (TLS), the encrypted tunnel must be negotiated: the client and the server must agree on the version of the TLS protocol, choose the ciphersuite, and verify certificates if necessary.

TLS runs over a reliable transport (TCP), which means that we must first complete the TCP three-way handshake, which takes one full round trip. See **Figure 1-9** for more details.

![TLS Handshake](/images/loar/1-9.png)
_Figure 1-9. TLS Handshake_

With the TCP connection in place, the client sends a number of specifications in plain text, such as the version of the TLS protocol it is running, the list of supported ciphersuites, and other TLS options it may want to use.

It is called a Handshake Phase when a client (e.g., a browser) connects to a server, they perform a handshake to agree on encryption algorithms and exchange keys. During this phase, the server provides its certificate, which contains a public key.  
Then there is a Key Exchange phase when a shared session key is generated using asymmetric cryptography (e.g., RSA, Diffie-Hellman, or Elliptic Curve Cryptography). This session key is used for symmetric encryption of the session data.

All subsequent communication is encrypted using the session key, providing confidentiality and integrity.

##### TLS Termination

TLS termination is a mechanism that terminates the encrypted connection at a designated server, typically a load balancer or reverse proxy, rather than having each backend server manage its own encryption.

![TLS Termination](/images/loar/1-10.png)
_Figure 1-10. TLS Termination_

A typical setup would involve a TLS-terminated load balancer that decrypts incoming HTTPS traffic and then forwards it to application servers using HTTP (unencrypted) or re-encrypts it using a separate internal certificate.

##### Server Name Indication (SNI)

SNI is an extension to the TLS protocol that allows a client to specify the hostname it is trying to connect to during the TLS handshake. This is particularly important in scenarios where multiple domains are hosted on a single IP address, as it helps the server present the correct TLS certificate for the requested domain.

When a user types `https://cecg.io` in their web browser, it initiates a TLS handshake and sends the SNI extension indicating it wants to connect to **cecg.io**. The server, upon receiving this request, presents the certificate for **cecg.io**, enabling the secure connection.

**SNI is sent in plaintext, which can reveal the hostname being accessed to eavesdroppers. This can be a privacy concern for users who wish to keep their browsing activities confidential.**

#### Public Key Infrastructure (PKI)

PKI is a framework used to manage digital certificates and public-key encryption. In the context of TLS, PKI is critical for establishing secure communication between clients and servers over a network, such as in HTTPS connections.

PKI uses digital certificates to authenticate identities. These certificates contain a public key and are typically issued by a trusted entity called a Certificate Authority (CA).

##### Chain of trust

PKI relies on a chain of trust, where the server certificate is trusted because it is signed by an intermediate CA, which is further signed by a root CA. 

![Chain of certificates](/images/loar/1-11.png)
_Figure 1-11. Chain of certificates_

The root CA is a trusted entity that issues digital certificates, e.g. companies like DigiCert, GlobalSign, and Let’s Encrypt. Root CAs are at the top of the certificate hierarchy and act as a trust anchor. Their certificates are embedded in web browsers and operating systems, allowing these systems to trust certificates issued by them.

Every browser and operating system provides a mechanism for you to manually import any certificate you trust. How you obtain the certificate and verify its integrity is completely up to you.

In practice, it would be impractical to store and manually verify each and every key for every website. The most common solution is to use certificate authorities (CAs) to do this job for us: the browser specifies which CAs to trust (root CAs), and the burden is then on the CAs to verify each site they sign, and to audit and verify that these certificates are not misused or compromised. If the security of any site with the CA’s certificate is breached, then it is also the responsibility of that CA to revoke the compromised certificate.

Verifying the chain of trust requires that the browser traverse the chain, starting from the site certificate, and recursively verify the certificate of the parent until it reaches a trusted root. Hence, it is critical that the provided chain includes all the intermediate certificates. If any are omitted, the browser will be forced to pause the verification process and fetch the missing certificates, adding additional DNS lookups, TCP handshakes, and HTTP requests into the process.

Currently, the majority of web browsers are shipped with pre-installed intermediate certificates issued and signed by a certificate authority, by public keys certified by so-called root certificates. This means browsers need to carry a large number of different certificate providers, increasing the risk of a key compromise.

When a key is known to be compromised, it could be fixed by revoking the certificate, but such a compromise is not easily detectable and can be a huge security breach. Browsers have to issue a security patch to revoke intermediary certificates issued by a compromised root certificate authority.

##### TLS Certificates

A TLS certificate is a digital certificate that authenticates the identity of a website and encrypts information sent to the server using TLS technology. This ensures that data transferred between the web server and the browser remains private and secure.

TLS operates on a PKI model, where each certificate contains a pair of keys: a public key and a private key. These keys are used to encrypt and decrypt data. The public key is included in the TLS certificate, while the private key is kept secret on the server. 

![Certificate viewer in Chrome](/images/loar/1-12.png)
_Figure 1-12. Certificate viewer in Chrome_

**Figure 1-12** shows a TLS certificate issued for **cecg.io** domain by **Let’s Encrypt** (CA). The certificate is valid until December 30th, 2024\.

TLS certificates have a finite lifespan and need to be renewed periodically to maintain secure communications. As of now, the maximum validity period for TLS certificates is 397 days, which is approximately 13 months. If a TLS certificate expires, the website will no longer be accessible via HTTPS, and users will see a warning message indicating that the site is not secure.

For more information about certificates refer to [RFC 5280 \- Internet X.509 Public Key Infrastructure Certificate and Certificate Revocation List (CRL) Profile](https://datatracker.ietf.org/doc/html/rfc5280).

##### Self-signed TLS Certificates

A self-signed TLS certificate is a digital certificate that is not issued by a trusted Certificate Authority (CA) but is instead created, signed, and validated by the organisation or individual who owns the website. Self-signed certificates are often used in development, testing, or internal environments where a CA-signed certificate is not required or practical.

OpenSSL can be used to generate a self-signed TLS certificate, check [Creating a Self-Signed Certificate With OpenSSL | Baeldung](https://www.baeldung.com/openssl-self-signed-cert) tutorial for more details

To use a self-signed certificate without warnings, you need to manually add it to the list of trusted certificates on the client device or browser, which essentially tells the client to trust the identity of the server.

See [Adding a Self-Signed Certificate to the Trusted List | Baeldung on Linux](https://www.baeldung.com/linux/add-self-signed-certificate-trusted-list) tutorial for instructions.

#### CertManager

[CertManager](https://cert-manager.io/) is an open-source Kubernetes add-on designed to automate the management, issuance, and renewal of TLS certificates within Kubernetes clusters. It simplifies the process of handling secure certificates, allowing Kubernetes resources to communicate securely without manual intervention. CertManager ensures that applications and services within the cluster have valid and updated certificates at all times, reducing the operational overhead of certificate management.

It supports multiple issuers, including: ACME (Let’s Encrypt and others), HashiCorp Vault, Self-signed certificates, Private CAs.

In our cluster we use CertManager to automatically obtain and renew HTTPS certificates for Kubernetes applications exposed via Ingress or Gateway.

![CertManager](/images/loar/1-13.png)
_Figure 1-13. CertManager_

CertManager greatly simplifies the process of managing TLS certificates, making it a valuable tool for securing communication in Kubernetes environments.

##### Let’s Encrypt

Let’s Encrypt is a free, automated, and open Certificate Authority (CA) that provides digital certificates to enable HTTPS for websites.

All digital certificates provided by Let’s Encrypt are issued at no cost.

Let’s Encrypt uses the ACME (Automated Certificate Management Environment) protocol to validate domain ownership, issue certificates, and automatically handle renewals.

Let’s Encrypt certificates are valid for 90 days and must be renewed regularly, though automation tools like CertManager make this process seamless.

#### Troubleshooting

OpenSSL is a cryptography toolkit implementing the Secure Sockets Layer (SSL v2/v3) and Transport Layer Security (TLS v1) network protocols and related cryptography standards required by them.

It also includes the **openssl** command, which provides a rich variety of commands You can use the same command to debug problems with SSL certificates.

See [Appendix K](../../appendix/k) for more details about **openssl** command.

To simulate certificate issues we can use [https://badssl.com/](https://badssl.com/) website.

##### Expired certificate

Perhaps the most common error you will encounter is that the certificate has expired.

```
$ curl https://expired.badssl.com/

curl: (60) SSL certificate problem: certificate has expired
```

The validity of the certificate can be checked in a web browser, or using **openssl**, for example:

```
$ openssl s_client -connect expired.badssl.com:443 -showcerts

Connecting to 104.154.89.105
CONNECTED(00000005)
. . . full details skipped . . .
Verification error: certificate has expired
```

This command connects to the specified server and displays the full certificate chain which is skipped in the example output.

If you want to see only the date when certificate expires:

```
$ </dev/null openssl s_client -connect expired.badssl.com:443 2>/dev/null | openssl x509 -noout -dates

notBefore=Apr  9 00:00:00 2015 GMT
notAfter=Apr 12 23:59:59 2015 GMT
```

##### Wrong domain

Perhaps the second most widely encountered error: the name on the certificate does not match the name by which you tried to access the service.

```
$ curl https://wrong.host.badssl.com
curl: (60) SSL: no alternative certificate subject name matches target host name 'wrong.host.badssl.com'
```

A common mistake you will encounter is the incorrect use of wildcards, whereby a service operator will assume that a wildcard cert for **\*.example.com** would be valid for **foo.bar.example.com**.

##### Self-signed certificate

When a client (e.g., a web browser) connects to a server using a self-signed certificate, it checks the certificate’s issuer. Because the certificate is not signed by a recognized CA, the browser will display a warning, prompting the user to decide whether to trust the certificate manually.

```
$ curl https://self-signed.badssl.com
curl: (60) SSL certificate problem: self signed certificate

$ openssl s_client -connect self-signed.badssl.com:443 -showcerts
Connecting to 104.154.89.105
. . .
Verify return code: 18 (self-signed certificate)
```

For debugging purposes there is  **\--insecure** flag for **curl** that ignores certificate issues:

```
$ curl -I --insecure  https://self-signed.badssl.com

HTTP/1.1 200 OK
```

#### Further Reading

- [RFC 2818 \- HTTP Over TLS](https://datatracker.ietf.org/doc/html/rfc2818)
- [RFC 5246 \- The Transport Layer Security (TLS) Protocol Version 1.2](https://datatracker.ietf.org/doc/html/rfc5246)
- [High Performance Browser Networking](https://hpbn.co/)
- [Debugging Certificate Errors](https://www.netmeister.org/blog/debugging-certificate-errors.html)
- [sysadvent: Day 3 \- Debugging SSL/TLS With openssl(1)](https://sysadvent.blogspot.com/2010/12/day-3-debugging-ssltls-with-openssl1.html)
- [Generating a self-signed certificate using OpenSSL](https://www.ibm.com/docs/en/api-connect/10.0.x?topic=profile-generating-self-signed-certificate-using-openssl) 
