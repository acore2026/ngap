package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PDUSessionResourceNotifyItem struct {
	PDUSessionID                     PDUSessionID
	PDUSessionResourceNotifyTransfer aper.OctetString
	IEExtensions                     *ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs `aper:"optional"`
}
