+++
title = "F - Using traceroute"
weight = 6
chapter = false
+++

Another tool is **traceroute**. It is a network diagnostic tool used to track the path that packets take from your computer to a specific destination, usually an IP address or domain name. It helps identify routing issues, latency, and other network problems.

When you run a traceroute command, your computer sends a series of packets to the target host. Traceroute uses the Time-to-Live (TTL) field in the packet header as a way to trace the path through which data travels across the network. TTL limits the number of hops the packet can pass through before being discarded. This prevents packets from endlessly circulating in a network loop if there’s a routing problem.

When a traceroute command is initiated, the first packet sent to the destination has its TTL set to **1**. As the packet reaches the first device, the device decreases the TTL by 1, making it **0**. When TTL reaches 0, the packet is discarded by the device. The device then sends back an ICMP "Time Exceeded" message to the sender, indicating that the packet couldn’t go any further.

After receiving the response from the first hop, traceroute sends a second packet, but this time it sets the TTL to **2**. This process continues as traceroute increments the TTL by 1 for each subsequent packet (TTL \= 3, 4, 5, etc.), allowing the packet to go further into the network with each step, hop by hop. Eventually, a packet will reach its final destination. When the TTL is high enough to reach the destination, the destination host sends back an ICMP Echo Reply or other appropriate responses.

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

For each hop there is the round-trip time for packets. Typically, three time values are shown, representing the time it took for three packets to go to the hop and back.

You might see asterisks (**\* \* \***) instead of time values. This means that no response was received from that hop. The device at that hop may have firewall rules in place that block ICMP packets (which traceroute uses to determine hop times) but still allows other types of traffic (like TCP or UDP). This would result in no response being sent back to the source.

Despite receiving \* \* \*, the traceroute output is still considered successful if it reaches its final destination. This indicates that the destination is reachable even if intermediate hops are unresponsive.

Many firewalls and network security devices block ICMP traffic to prevent ping flooding, DoS attacks, or to secure the network against certain types of reconnaissance. This means that traditional ICMP-based tools like ping and traceroute may not work as expected because their ICMP Echo Request messages won’t get responses.

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

When ICMP  traffic is blocked, you can use TCP as an alternative for network diagnostics. Using TCP packets, you can perform network diagnostics and simulate similar functionalities as ICMP tools without being hindered by ICMP restrictions.

Instead of sending ICMP Echo Requests, TCP traceroute sends SYN packets to initiate a TCP connection on a specified port.

You can use the **-T** flag with **traceroute** to do tracing with TCP SYN packets.

```
$ traceroute -T cecg.io

traceroute to cecg.io (75.2.60.5), 30 hops max, 60 byte packets
 1  172.17.0.1 (172.17.0.1)  0.280 ms  0.006 ms *
 2  acd89244c803f7181.awsglobalaccelerator.com (75.2.60.5)  61.390 ms  60.019 ms  51.133 ms
```

In scenarios where ICMP is blocked, using TCP for traceroute and other diagnostic tools is a viable workaround.
