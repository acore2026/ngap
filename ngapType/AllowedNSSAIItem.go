package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type AllowedNSSAIItem struct {
	SNSSAI       SNSSAI                                            `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerAllowedNSSAIItemExtIEs `aper:"optional"`
}
