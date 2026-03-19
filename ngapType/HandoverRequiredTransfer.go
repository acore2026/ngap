package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type HandoverRequiredTransfer struct {
	DirectForwardingPathAvailability *DirectForwardingPathAvailability                         `aper:"optional"`
	IEExtensions                     *ProtocolExtensionContainerHandoverRequiredTransferExtIEs `aper:"optional"`
}
