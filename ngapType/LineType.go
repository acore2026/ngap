package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const ( /* Enum Type */
	LineTypePresentDsl aper.Enumerated = 0
	LineTypePresentPon aper.Enumerated = 1
)

type LineType struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
