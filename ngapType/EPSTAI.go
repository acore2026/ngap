package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type EPSTAI struct {
	PLMNIdentity PLMNIdentity
	EPSTAC       EPSTAC
	IEExtensions *ProtocolExtensionContainerEPSTAIExtIEs `aper:"optional"`
}
