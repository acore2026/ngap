package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type FiveGTMSI struct {
	Value aper.OctetString `aper:"sizeLB:4,sizeUB:4"`
}
