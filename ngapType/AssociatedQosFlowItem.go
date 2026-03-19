package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type AssociatedQosFlowItem struct {
	QosFlowIdentifier        QosFlowIdentifier
	QosFlowMappingIndication *aper.Enumerated                                       `aper:"optional,valueExt,valueLB:0,valueUB:1"`
	IEExtensions             *ProtocolExtensionContainerAssociatedQosFlowItemExtIEs `aper:"optional"`
}
