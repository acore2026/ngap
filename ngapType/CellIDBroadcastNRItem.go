package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type CellIDBroadcastNRItem struct {
	NRCGI        NRCGI                                                  `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs `aper:"optional"`
}
