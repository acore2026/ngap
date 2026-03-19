package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PDUSessionResourceSetupItemCxtReq struct {
	PDUSessionID                           PDUSessionID
	NASPDU                                 *NASPDU `aper:"optional"`
	SNSSAI                                 SNSSAI  `aper:"valueExt"`
	PDUSessionResourceSetupRequestTransfer aper.OctetString
	IEExtensions                           *ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs `aper:"optional"`
}
