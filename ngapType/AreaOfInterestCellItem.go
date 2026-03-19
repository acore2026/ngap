package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type AreaOfInterestCellItem struct {
	NGRANCGI     NGRANCGI                                                `aper:"valueLB:0,valueUB:2"`
	IEExtensions *ProtocolExtensionContainerAreaOfInterestCellItemExtIEs `aper:"optional"`
}
