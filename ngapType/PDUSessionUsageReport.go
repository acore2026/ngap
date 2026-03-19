package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PDUSessionUsageReport struct {
	RATType                   aper.Enumerated
	PDUSessionTimedReportList VolumeTimedReportList
	IEExtensions              *ProtocolExtensionContainerPDUSessionUsageReportExtIEs `aper:"optional"`
}
