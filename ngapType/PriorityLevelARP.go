package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PriorityLevelARP struct {
	Value int64 `aper:"valueLB:1,valueUB:15"`
}
