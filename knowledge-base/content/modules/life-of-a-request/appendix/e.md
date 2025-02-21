+++
title = "E - Using ping"
weight = 5
chapter = false
+++

The **ping** command is a network utility used to test the reachability of a host on an IP network and measure the round-trip time for messages sent from the source to the destination and back.

When you run the **ping** command, your computer sends ICMP (Internet Control Message Protocol) Echo Request packets to the target host. ICMP is a network-layer protocol that reports errors and provides information related to the success or failure of data delivery. If the target host is reachable and configured to respond, it will send back an ICMP Echo Reply packet.

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

The reply includes information like:

- packet size, e.g.: **64 bytes**
- **icmp_seq** refers to the ICMP sequence number, e.g. **0, 1, 2**.  It represents the sequence number of the ICMP Echo Request packet sent to the target host.
- **ttl** (Time to Live), e.g. **245**. Represents the maximum number of hops (network devices) that a packet can pass through before being discarded.
- **time** (round-trip time): **48.675 ms**, which helps in determining the latency between the source and destination.

The command typically runs multiple times, sending several packets. It shows:

- The number of packets sent, received, and lost (if any).
- The minimum, maximum, and average round-trip times.
- Packet loss percentage (if some packets don’t get a reply).

```
$ ping 10.0.0.1

PING 10.0.0.1 (10.0.0.1): 56 data bytes
Request timeout for icmp_seq 0
Request timeout for icmp_seq 1
Request timeout for icmp_seq 2
^C
--- 10.0.0.1 ping statistics ---
4 packets transmitted, 0 packets received, 100.0% packet loss
```

If some packets are lost, it may indicate network issues, such as congestion or a problem with the destination host. 
