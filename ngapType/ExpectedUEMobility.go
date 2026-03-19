package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	ExpectedUEMobilityPresentStationary aper.Enumerated = 0
	ExpectedUEMobilityPresentMobile     aper.Enumerated = 1
)

type ExpectedUEMobility struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
