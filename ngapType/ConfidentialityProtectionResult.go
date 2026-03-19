package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	ConfidentialityProtectionResultPresentPerformed    aper.Enumerated = 0
	ConfidentialityProtectionResultPresentNotPerformed aper.Enumerated = 1
)

type ConfidentialityProtectionResult struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
