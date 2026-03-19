package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type CompletedCellsInTAIEUTRAItem struct {
	EUTRACGI     EUTRACGI                                                      `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs `aper:"optional"`
}
