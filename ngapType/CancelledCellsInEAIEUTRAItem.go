package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type CancelledCellsInEAIEUTRAItem struct {
	EUTRACGI           EUTRACGI `aper:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs `aper:"optional"`
}
