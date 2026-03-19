package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	DLForwardingPresentDlForwardingProposed aper.Enumerated = 0
)

type DLForwarding struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
