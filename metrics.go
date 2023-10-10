package wireguardreceiver

import (
	"github.com/rogercoll/wireguardreceiver/internal/metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func recordPeerMetrics(mb *metadata.MetricsBuilder, now pcommon.Timestamp, deviceName string, peer *wgtypes.Peer) {
	// set metrics
	mb.RecordWireguardPeerNetworkIoUsageRxBytesDataPoint(now, peer.ReceiveBytes)
	mb.RecordWireguardPeerNetworkIoUsageTxBytesDataPoint(now, peer.TransmitBytes)
	// Always-present resource attrs
	rb := mb.NewResourceBuilder()

	rb.SetWireguardDeviceName(deviceName)
	rb.SetWireguardPeerName(peer.PublicKey.String())

	mb.EmitForResource(metadata.WithResource(rb.Emit()))
}
