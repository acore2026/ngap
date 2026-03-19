package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type QoSFlowsUsageReportItem struct {
	QosFlowIdentifier       QosFlowIdentifier
	RATType                 aper.Enumerated
	QoSFlowsTimedReportList VolumeTimedReportList
	IEExtensions            *ProtocolExtensionContainerQoSFlowsUsageReportItemExtIEs `aper:"optional"`
}
