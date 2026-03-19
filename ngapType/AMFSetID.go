package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type AMFSetID struct {
	Value aper.BitString `aper:"sizeLB:10,sizeUB:10"`
}
