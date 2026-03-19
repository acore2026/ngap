package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type CancelledCellsInTAINRItem struct {
	NRCGI              NRCGI `aper:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs `aper:"optional"`
}
