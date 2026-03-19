package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type QosFlowSetupRequestItem struct {
	QosFlowIdentifier         QosFlowIdentifier
	QosFlowLevelQosParameters QosFlowLevelQosParameters                                `aper:"valueExt"`
	ERABID                    *ERABID                                                  `aper:"optional"`
	IEExtensions              *ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs `aper:"optional"`
}
