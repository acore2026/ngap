package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PDUSessionResourceModifyItemModRes struct {
	PDUSessionID                             PDUSessionID
	PDUSessionResourceModifyResponseTransfer aper.OctetString
	IEExtensions                             *ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs `aper:"optional"`
}
