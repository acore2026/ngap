package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PDUSessionResourceNotifyTransfer struct {
	QosFlowNotifyList   *QosFlowNotifyList                                                `aper:"optional"`
	QosFlowReleasedList *QosFlowListWithCause                                             `aper:"optional"`
	IEExtensions        *ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs `aper:"optional"`
}
