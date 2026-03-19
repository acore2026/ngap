package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type UnavailableGUAMIItem struct {
	GUAMI                        GUAMI                                                 `aper:"valueExt"`
	TimerApproachForGUAMIRemoval *TimerApproachForGUAMIRemoval                         `aper:"optional"`
	BackupAMFName                *AMFName                                              `aper:"optional"`
	IEExtensions                 *ProtocolExtensionContainerUnavailableGUAMIItemExtIEs `aper:"optional"`
}
