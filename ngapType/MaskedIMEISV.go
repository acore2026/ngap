package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type MaskedIMEISV struct {
	Value aper.BitString `aper:"sizeLB:64,sizeUB:64"`
}
