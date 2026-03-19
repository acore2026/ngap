package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type TransportLayerAddress struct {
	Value aper.BitString `aper:"sizeExt,sizeLB:1,sizeUB:160"`
}
