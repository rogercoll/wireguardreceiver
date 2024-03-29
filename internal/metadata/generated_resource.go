// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
// The ResourceBuilder is not thread-safe and must not to be used in multiple goroutines.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

// NewResourceBuilder creates a new ResourceBuilder. This method should be called on the start of the application.
func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetWireguardDeviceName sets provided value as "wireguard.device.name" attribute.
func (rb *ResourceBuilder) SetWireguardDeviceName(val string) {
	if rb.config.WireguardDeviceName.Enabled {
		rb.res.Attributes().PutStr("wireguard.device.name", val)
	}
}

// SetWireguardDeviceTypeLinuxKernel sets "wireguard.device.type=Linux kernel" attribute.
func (rb *ResourceBuilder) SetWireguardDeviceTypeLinuxKernel() {
	if rb.config.WireguardDeviceType.Enabled {
		rb.res.Attributes().PutStr("wireguard.device.type", "Linux kernel")
	}
}

// SetWireguardDeviceTypeOpenBSDKernel sets "wireguard.device.type=OpenBSD kernel" attribute.
func (rb *ResourceBuilder) SetWireguardDeviceTypeOpenBSDKernel() {
	if rb.config.WireguardDeviceType.Enabled {
		rb.res.Attributes().PutStr("wireguard.device.type", "OpenBSD kernel")
	}
}

// SetWireguardDeviceTypeFreeBSDKernel sets "wireguard.device.type=FreeBSD kernel" attribute.
func (rb *ResourceBuilder) SetWireguardDeviceTypeFreeBSDKernel() {
	if rb.config.WireguardDeviceType.Enabled {
		rb.res.Attributes().PutStr("wireguard.device.type", "FreeBSD kernel")
	}
}

// SetWireguardDeviceTypeWindowsKernel sets "wireguard.device.type=Windows kernel" attribute.
func (rb *ResourceBuilder) SetWireguardDeviceTypeWindowsKernel() {
	if rb.config.WireguardDeviceType.Enabled {
		rb.res.Attributes().PutStr("wireguard.device.type", "Windows kernel")
	}
}

// SetWireguardDeviceTypeUserspace sets "wireguard.device.type=userspace" attribute.
func (rb *ResourceBuilder) SetWireguardDeviceTypeUserspace() {
	if rb.config.WireguardDeviceType.Enabled {
		rb.res.Attributes().PutStr("wireguard.device.type", "userspace")
	}
}

// SetWireguardDeviceTypeUnknown sets "wireguard.device.type=unknown" attribute.
func (rb *ResourceBuilder) SetWireguardDeviceTypeUnknown() {
	if rb.config.WireguardDeviceType.Enabled {
		rb.res.Attributes().PutStr("wireguard.device.type", "unknown")
	}
}

// SetWireguardPeerName sets provided value as "wireguard.peer.name" attribute.
func (rb *ResourceBuilder) SetWireguardPeerName(val string) {
	if rb.config.WireguardPeerName.Enabled {
		rb.res.Attributes().PutStr("wireguard.peer.name", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
