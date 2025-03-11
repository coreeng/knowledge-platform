+++
title = "K - Using OpenSSL"
weight = 11
chapter = false
+++

Debugging TLS issues using OpenSSL can help identify problems with SSL/TLS certificates, handshakes, ciphers, and other security configurations. OpenSSL provides a variety of tools and commands to troubleshoot these issues.

You can use openssl s_client to test and debug a connection to a TLS server. This command establishes a TLS connection and outputs detailed information about the session.

```
$ openssl s_client -connect cecg.io:443

Connecting to 75.2.60.5
CONNECTED(00000005)
depth=2 C=US, O=Internet Security Research Group, CN=ISRG Root X1
verify return:1
depth=1 C=US, O=Let's Encrypt, CN=E5
verify return:1
depth=0 CN=cecg.io
verify return:1
---
Certificate chain
 0 s:CN=cecg.io
   i:C=US, O=Let's Encrypt, CN=E5
   a:PKEY: id-ecPublicKey, 256 (bit); sigalg: ecdsa-with-SHA384
   v:NotBefore: Oct  1 21:33:44 2024 GMT; NotAfter: Dec 30 21:33:43 2024 GMT
 1 s:C=US, O=Let's Encrypt, CN=E5
   i:C=US, O=Internet Security Research Group, CN=ISRG Root X1
   a:PKEY: id-ecPublicKey, 384 (bit); sigalg: RSA-SHA256
   v:NotBefore: Mar 13 00:00:00 2024 GMT; NotAfter: Mar 12 23:59:59 2027 GMT
---
```

If you want to test a specific TLS version (e.g., TLS 1.2 or TLS 1.3), you can specify the protocol using the **-tls1_2** or **-tls1_3** flag.

```
$ openssl s_client -connect cecg.io:443 -tls1_3

. . .
    Protocol  : TLSv1.3
    Cipher    : TLS_AES_128_GCM_SHA256
. . .
```

For very detailed debugging, you can enable verbose output with the **-debug** flag:

```
$ openssl s_client -connect cecg.io:443 -debug

Connecting to 75.2.60.5
CONNECTED(00000005)
write to 0x600001e52300 [0x138828800] (315 bytes => 315 (0x13B))
0000 - 16 03 01 01 36 01 00 01-32 03 03 f6 95 52 0a 7f   ....6...2....R..
0010 - 53 b0 e8 c1 ba 7e 8a 13-72 66 c9 90 b4 22 22 58   S....~..rf...""X
0020 - f3 71 c0 e7 b0 44 79 f5-3f 20 e0 20 66 4d 16 60   .q...Dy.? . fM.`
. . .
```
