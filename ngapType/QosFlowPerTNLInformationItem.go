package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type QosFlowPerTNLInformationItem struct {
	QosFlowPerTNLInformation QosFlowPerTNLInformation                                      `aper:"valueExt"`
	IEExtensions             *ProtocolExtensionContainerQosFlowPerTNLInformationItemExtIEs `aper:"optional"`
}
