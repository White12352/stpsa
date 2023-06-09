### Structure

```json
{
  "detour": "upstream-out",
  "bind_interface": "en0",
  "inet4_bind_address": "0.0.0.0",
  "inet6_bind_address": "::",
  "routing_mark": 1234,
  "reuse_addr": false,
  "connect_timeout": "5s",
  "tcp_fast_open": false,
  "udp_fragment": false,
  "domain_strategy": "prefer_ipv6",
  "fallback_delay": "300ms"
}
```

### Fields

| Field                                                                                                                | Available Context |
|----------------------------------------------------------------------------------------------------------------------|-------------------|
| `bind_interface` /`*bind_address` /`routing_mark` /`reuse_addr` / `tcp_fast_open`/ `udp_fragment` /`connect_timeout` | `detour` not set  |

#### detour

The tag of the upstream outbound.

#### detour_redir

The upstream outbound redirection is allowed.
It is mainly used in conjunction with health checks, 
see the `detour_of` field of the health check.

#### bind_interface

The network interface to bind to.

#### inet4_bind_address

The IPv4 address to bind to.

#### inet6_bind_address

The IPv6 address to bind to.

#### routing_mark

!!! error ""

    Only supported on Linux.

Set netfilter routing mark.

#### reuse_addr

Reuse listener address.

#### tcp_fast_open

Enable TCP Fast Open.

#### udp_fragment

Enable UDP fragmentation.

#### connect_timeout

Connect timeout, in golang's Duration format.

A duration string is a possibly signed sequence of
decimal numbers, each with optional fraction and a unit suffix,
such as "300ms", "-1.5h" or "2h45m".
Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

#### domain_strategy

One of `prefer_ipv4` `prefer_ipv6` `ipv4_only` `ipv6_only`.

If set, the requested domain name will be resolved to IP before connect.

| Outbound | Effected domains         | Fallback Value                            |
|----------|--------------------------|-------------------------------------------|
| `direct` | Domain in request        | Take `inbound.domain_strategy` if not set | 
| others   | Domain in server address | /                                         |

#### fallback_delay

The length of time to wait before spawning a RFC 6555 Fast Fallback connection.
That is, is the amount of time to wait for connection to succeed before assuming
that IPv4/IPv6 is misconfigured and falling back to other type of addresses.
If zero, a default delay of 300ms is used.

Only take effect when `domain_strategy` is set.