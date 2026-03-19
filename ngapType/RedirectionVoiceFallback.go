package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	RedirectionVoiceFallbackPresentPossible    aper.Enumerated = 0
	RedirectionVoiceFallbackPresentNotPossible aper.Enumerated = 1
)

type RedirectionVoiceFallback struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
