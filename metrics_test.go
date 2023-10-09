package wireguardreceiver

import (
	"testing"
	"time"

	"github.com/rogercoll/wireguardreceiver/internal/metadata"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func TestConvertPeerToMetrics(t *testing.T) {
	peer, err := getPeer()
	assert.Nil(t, err)

	mb := metadata.NewMetricsBuilder(metadata.DefaultMetricsBuilderConfig(), receivertest.NewNopCreateSettings())
	recordPeerMetrics(mb, pcommon.NewTimestampFromTime(time.Now()), "wg0", peer)
	assertPeerToMetrics(t, peer, mb.Emit())
}

func assertPeerToMetrics(t *testing.T, peer *wgtypes.Peer, md pmetric.Metrics) {
	assert.Equal(t, md.ResourceMetrics().Len(), 1)
	rsm := md.ResourceMetrics().At(0)

	resourceAttrs := map[string]string{
		"wireguard.peer.name":   "aPxGwq8zERHQ3Q1cOZFdJ+cvJX5Ka4mLN38AyYKYF10=",
		"wireguard.device.name": "wg0",
	}

	for k, v := range resourceAttrs {
		attr, exists := rsm.Resource().Attributes().Get(k)
		assert.True(t, exists)
		assert.Equal(t, v, attr.Str())
	}

	assert.Equal(t, rsm.ScopeMetrics().Len(), 1)

	metrics := rsm.ScopeMetrics().At(0).Metrics()
	assert.Equal(t, metrics.Len(), 2)
}

func getPeer() (*wgtypes.Peer, error) {
	key, err := wgtypes.ParseKey("aPxGwq8zERHQ3Q1cOZFdJ+cvJX5Ka4mLN38AyYKYF10=")
	if err != nil {
		return nil, err
	}

	return &wgtypes.Peer{
		PublicKey:     key,
		ReceiveBytes:  124,
		TransmitBytes: 92112,
	}, nil
}
