# Wireguard Stats Receiver

The Wireguard stats recevier queries the stats of the connected
Wireguard peers via generic netlink (Linux).
[wgctrl](https://github.com/WireGuard/wgctrl-go) is used to fetch the
stats and for the objects definition.

This receiver can we integrated with the custom [OpenTelemetry Collector
Builder](https://github.com/open-telemetry/opentelemetry-collector/tree/main/cmd/builder):

```yaml
receivers:
  - gomod: github.com/rogercoll/wireguardreceiver v0.0.3
```

**Requires privileged execution**

## Metrics

```
peer.usage.rx_bytes
peer.usage.tx_bytes
```

## Attributes

```
peer.device.name	## Interface name
peer.name		## Peer's public key
```

### Todo

- [ ] Add latest handshake
