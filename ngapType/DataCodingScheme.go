package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type DataCodingScheme struct {
	Value aper.BitString `aper:"sizeLB:8,sizeUB:8"`
}
