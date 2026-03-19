package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	TimerApproachForGUAMIRemovalPresentApplyTimer aper.Enumerated = 0
)

type TimerApproachForGUAMIRemoval struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
