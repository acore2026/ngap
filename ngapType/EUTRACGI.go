package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type EUTRACGI struct {
	PLMNIdentity      PLMNIdentity
	EUTRACellIdentity EUTRACellIdentity
	IEExtensions      *ProtocolExtensionContainerEUTRACGIExtIEs `aper:"optional"`
}
