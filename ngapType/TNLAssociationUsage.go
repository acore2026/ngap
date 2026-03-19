package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	TNLAssociationUsagePresentUe    aper.Enumerated = 0
	TNLAssociationUsagePresentNonUe aper.Enumerated = 1
	TNLAssociationUsagePresentBoth  aper.Enumerated = 2
)

type TNLAssociationUsage struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}
