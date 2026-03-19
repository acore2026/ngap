package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type ServiceAreaInformationItem struct {
	PLMNIdentity   PLMNIdentity
	AllowedTACs    *AllowedTACs                                                `aper:"optional"`
	NotAllowedTACs *NotAllowedTACs                                             `aper:"optional"`
	IEExtensions   *ProtocolExtensionContainerServiceAreaInformationItemExtIEs `aper:"optional"`
}
