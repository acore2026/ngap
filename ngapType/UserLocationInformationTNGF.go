package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type UserLocationInformationTNGF struct {
	TNAPID       TNAPID
	IPAddress    TransportLayerAddress
	PortNumber   *PortNumber                                                  `aper:"optional"`
	IEExtensions *ProtocolExtensionContainerUserLocationInformationTNGFExtIEs `aper:"optional"`
}
