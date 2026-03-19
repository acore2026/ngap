package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PDUSessionResourceSecondaryRATUsageItem struct {
	PDUSessionID                        PDUSessionID
	SecondaryRATDataUsageReportTransfer aper.OctetString
	IEExtensions                        *ProtocolExtensionContainerPDUSessionResourceSecondaryRATUsageItemExtIEs `aper:"optional"`
}
