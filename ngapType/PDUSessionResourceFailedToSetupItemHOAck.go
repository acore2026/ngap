package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PDUSessionResourceFailedToSetupItemHOAck struct {
	PDUSessionID                                   PDUSessionID
	HandoverResourceAllocationUnsuccessfulTransfer aper.OctetString
	IEExtensions                                   *ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs `aper:"optional"`
}
