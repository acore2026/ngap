package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PDUSessionResourceItemCxtRelReq struct {
	PDUSessionID PDUSessionID
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs `aper:"optional"`
}
