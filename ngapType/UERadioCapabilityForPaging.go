package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type UERadioCapabilityForPaging struct {
	UERadioCapabilityForPagingOfNR    *UERadioCapabilityForPagingOfNR                             `aper:"optional"`
	UERadioCapabilityForPagingOfEUTRA *UERadioCapabilityForPagingOfEUTRA                          `aper:"optional"`
	IEExtensions                      *ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs `aper:"optional"`
}
