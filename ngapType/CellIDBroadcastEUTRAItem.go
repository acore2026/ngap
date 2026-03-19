package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type CellIDBroadcastEUTRAItem struct {
	EUTRACGI     EUTRACGI                                                  `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs `aper:"optional"`
}
