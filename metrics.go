package wireguardreceiver

import (
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
	resourceAttr.PutStr("device", deviceName)

	ms := rs.ScopeMetrics().AppendEmpty().Metrics()
	appendPeerMetrics(ms, stats, pbts)

	return md
}

func appendIOMetrics(ms pmetric.MetricSlice, peer *wgtypes.Peer, ts pcommon.Timestamp) {
	gaugeI(ms, "peer.usage.rx_bytes", "By", []point{{intVal: peer.ReceiveBytes}}, ts)
	gaugeI(ms, "peer.usage.tx_bytes", "By", []point{{intVal: peer.TransmitBytes}}, ts)
}

func gauge(ms pmetric.MetricSlice, metricName string, unit string) pmetric.NumberDataPointSlice {
	metric := initMetric(ms, metricName, unit)
	gauge := metric.SetEmptyGauge()
	return gauge.DataPoints()
}

func gaugeI(ms pmetric.MetricSlice, metricName string, unit string, points []point, ts pcommon.Timestamp) {
	dataPoints := gauge(ms, metricName, unit)
	for _, pt := range points {
		dataPoint := dataPoints.AppendEmpty()
		dataPoint.SetTimestamp(ts)
		dataPoint.SetIntValue(int64(pt.intVal))
	}
}
