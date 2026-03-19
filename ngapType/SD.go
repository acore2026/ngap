package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type SD struct {
	Value aper.OctetString `aper:"sizeLB:3,sizeUB:3"`
}
