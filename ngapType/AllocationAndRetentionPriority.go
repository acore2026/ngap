package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type AllocationAndRetentionPriority struct {
	PriorityLevelARP        PriorityLevelARP
	PreEmptionCapability    PreEmptionCapability
	PreEmptionVulnerability PreEmptionVulnerability
	IEExtensions            *ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs `aper:"optional"`
}
