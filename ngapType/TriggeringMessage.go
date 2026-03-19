package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	TriggeringMessagePresentInitiatingMessage    aper.Enumerated = 0
	TriggeringMessagePresentSuccessfulOutcome    aper.Enumerated = 1
	TriggeringMessagePresentUnsuccessfullOutcome aper.Enumerated = 2
)

type TriggeringMessage struct {
	Value aper.Enumerated `aper:"valueLB:0,valueUB:2"`
}
