package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type GUAMI struct {
	PLMNIdentity PLMNIdentity
	AMFRegionID  AMFRegionID
	AMFSetID     AMFSetID
	AMFPointer   AMFPointer
	IEExtensions *ProtocolExtensionContainerGUAMIExtIEs `aper:"optional"`
}
