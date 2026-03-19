package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type NRCGI struct {
	PLMNIdentity   PLMNIdentity
	NRCellIdentity NRCellIdentity
	IEExtensions   *ProtocolExtensionContainerNRCGIExtIEs `aper:"optional"`
}
