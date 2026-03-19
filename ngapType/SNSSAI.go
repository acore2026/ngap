package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type SNSSAI struct {
	SST          SST
	SD           *SD                                     `aper:"optional"`
	IEExtensions *ProtocolExtensionContainerSNSSAIExtIEs `aper:"optional"`
}
