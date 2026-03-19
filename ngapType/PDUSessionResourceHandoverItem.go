package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PDUSessionResourceHandoverItem struct {
	PDUSessionID            PDUSessionID
	HandoverCommandTransfer aper.OctetString
	IEExtensions            *ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs `aper:"optional"`
}
