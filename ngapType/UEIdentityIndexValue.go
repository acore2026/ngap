package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	UEIdentityIndexValuePresentNothing int = iota /* No components present */
	UEIdentityIndexValuePresentIndexLength10
	UEIdentityIndexValuePresentChoiceExtensions
)

type UEIdentityIndexValue struct {
	Present          int
	IndexLength10    *aper.BitString `aper:"sizeLB:10,sizeUB:10"`
	ChoiceExtensions *ProtocolIESingleContainerUEIdentityIndexValueExtIEs
}
