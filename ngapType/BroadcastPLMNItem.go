package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type BroadcastPLMNItem struct {
	PLMNIdentity        PLMNIdentity
	TAISliceSupportList SliceSupportList
	IEExtensions        *ProtocolExtensionContainerBroadcastPLMNItemExtIEs `aper:"optional"`
}
