package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type QosFlowAcceptedItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowAcceptedItemExtIEs `aper:"optional"`
}
