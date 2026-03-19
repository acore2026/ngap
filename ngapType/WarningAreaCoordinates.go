package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type WarningAreaCoordinates struct {
	Value aper.OctetString `aper:"sizeLB:1,sizeUB:1024"`
}
