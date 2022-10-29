package wireguardreceiver

import (
	"fmt"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func peerToMetrics(ts time.Time, deviceName string, peer *wgtypes.Peer) pmetric.Metrics {
	pbts := pcommon.NewTimestampFromTime(ts)

	md := pmetric.NewMetrics()

	rs := md.ResourceMetrics().AppendEmpty()

	resourceAttr := rs.Resource().Attributes()
	resourceAttr.PutStr("peer.device.name", deviceName)
	resourceAttr.PutStr("peer.name", peer.PublicKey.String())
	ms := rs.ScopeMetrics().AppendEmpty().Metrics()
	appendPeerMetrics(ms, peer, pbts)

	return md
}

func appendPeerMetrics(ms pmetric.MetricSlice, peer *wgtypes.Peer, ts pcommon.Timestamp) {
	gaugeI(ms, "usage.rx_bytes", "By", peer.ReceiveBytes, ts)
	gaugeI(ms, "usage.tx_bytes", "By", peer.TransmitBytes, ts)
	gaugeI(ms, "last_handshake", "s", int64(peer.LastHandshakeTime.Second()), ts)
}

func initMetric(ms pmetric.MetricSlice, name, unit string) pmetric.Metric {
	m := ms.AppendEmpty()
	m.SetName(fmt.Sprintf("peer.%s", name))
	m.SetUnit(unit)
	return m
}

func gauge(ms pmetric.MetricSlice, metricName string, unit string) pmetric.NumberDataPointSlice {
	metric := initMetric(ms, metricName, unit)
	gauge := metric.SetEmptyGauge()
	return gauge.DataPoints()
}

func gaugeI(ms pmetric.MetricSlice, metricName string, unit string, value int64, ts pcommon.Timestamp) {
	dataPoints := gauge(ms, metricName, unit)
	dataPoint := dataPoints.AppendEmpty()
	dataPoint.SetTimestamp(ts)
	dataPoint.SetIntValue(value)
}
