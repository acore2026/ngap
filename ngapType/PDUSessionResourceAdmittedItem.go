package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PDUSessionResourceAdmittedItem struct {
	PDUSessionID                       PDUSessionID
	HandoverRequestAcknowledgeTransfer aper.OctetString
	IEExtensions                       *ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs `aper:"optional"`
}
