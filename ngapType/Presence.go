package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	PresencePresentOptional    aper.Enumerated = 0
	PresencePresentConditional aper.Enumerated = 1
	PresencePresentMandatory   aper.Enumerated = 2
)

type Presence struct {
	Value aper.Enumerated `aper:"valueLB:0,valueUB:2"`
}
