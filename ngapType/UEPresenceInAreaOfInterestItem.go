package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type UEPresenceInAreaOfInterestItem struct {
	LocationReportingReferenceID LocationReportingReferenceID
	UEPresence                   UEPresence
	IEExtensions                 *ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs `aper:"optional"`
}
