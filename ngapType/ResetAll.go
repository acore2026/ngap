package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	ResetAllPresentResetAll aper.Enumerated = 0
)

type ResetAll struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
