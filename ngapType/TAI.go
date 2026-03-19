package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type TAI struct {
	PLMNIdentity PLMNIdentity
	TAC          TAC
	IEExtensions *ProtocolExtensionContainerTAIExtIEs `aper:"optional"`
}
