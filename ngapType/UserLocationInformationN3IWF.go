package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type UserLocationInformationN3IWF struct {
	IPAddress    TransportLayerAddress
	PortNumber   PortNumber
	IEExtensions *ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs `aper:"optional"`
}
