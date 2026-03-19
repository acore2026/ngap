package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type FiveGSTMSI struct {
	AMFSetID     AMFSetID
	AMFPointer   AMFPointer
	FiveGTMSI    FiveGTMSI
	IEExtensions *ProtocolExtensionContainerFiveGSTMSIExtIEs `aper:"optional"`
}
