package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type QosFlowModifyConfirmItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowModifyConfirmItemExtIEs `aper:"optional"`
}
