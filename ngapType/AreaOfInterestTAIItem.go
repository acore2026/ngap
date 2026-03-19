package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type AreaOfInterestTAIItem struct {
	TAI          TAI                                                    `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs `aper:"optional"`
}
