package wireguardreceiver

import (
	"fmt"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var (
	errNoInterfacesFound          = fmt.Errorf("No Wireguard iterfaces found")
	errNoEndpointFound            = fmt.Errorf("No Wireguard endpoint found")
	errInvalidInterfaceAttributes = fmt.Errorf("Invalid interface attributes")
	localEndpointAttributes       = 5
	remoteEndpointAttributes      = 9
)

type clientFactory func() (wireguardClient, error)

type wireguardClient interface {
	Devices() ([]*wgtypes.Device, error)
}

func newWireguardClient() (wireguardClient, error) {
	return wgctrl.New()
}
