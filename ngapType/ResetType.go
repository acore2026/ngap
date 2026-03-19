package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

const (
	ResetTypePresentNothing int = iota /* No components present */
	ResetTypePresentNGInterface
	ResetTypePresentPartOfNGInterface
	ResetTypePresentChoiceExtensions
)

type ResetType struct {
	Present           int
	NGInterface       *ResetAll
	PartOfNGInterface *UEAssociatedLogicalNGConnectionList
	ChoiceExtensions  *ProtocolIESingleContainerResetTypeExtIEs
}
