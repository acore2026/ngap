package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	CellSizePresentVerysmall aper.Enumerated = 0
	CellSizePresentSmall     aper.Enumerated = 1
	CellSizePresentMedium    aper.Enumerated = 2
	CellSizePresentLarge     aper.Enumerated = 3
)

type CellSize struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}
