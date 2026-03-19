package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type DataForwardingResponseDRBItem struct {
	DRBID                        DRBID
	DLForwardingUPTNLInformation *UPTransportLayerInformation                                   `aper:"valueLB:0,valueUB:1,optional"`
	ULForwardingUPTNLInformation *UPTransportLayerInformation                                   `aper:"valueLB:0,valueUB:1,optional"`
	IEExtensions                 *ProtocolExtensionContainerDataForwardingResponseDRBItemExtIEs `aper:"optional"`
}
