[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# wireguard_stats

## Default Metrics

The following metrics are emitted by default. Each of them can be disabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: false
```

### wireguard.peer.network.io.usage.rx_bytes

Bytes received by the peer.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| By | Sum | Int | Cumulative | true |

### wireguard.peer.network.io.usage.tx_bytes

Bytes sent.

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| By | Sum | Int | Delta | false |

## Resource Attributes

| Name | Description | Values | Enabled |
| ---- | ----------- | ------ | ------- |
| wireguard.device.name | A Device is a WireGuard device. | Any Str | true |
| wireguard.device.type | A DeviceType specifies the underlying implementation of a WireGuard device. | Str: ``Linux kernel``, ``OpenBSD kernel``, ``FreeBSD kernel``, ``Windows kernel``, ``userspace``, ``unknown`` | true |
| wireguard.peer.name | A Device is a WireGuard device. | Any Str | true |
