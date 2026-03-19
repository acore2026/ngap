package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type CellType struct {
	CellSize     CellSize
	IEExtensions *ProtocolExtensionContainerCellTypeExtIEs `aper:"optional"`
}
