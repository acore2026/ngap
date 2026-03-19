package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type DRBStatusDL12 struct {
	DLCOUNTValue COUNTValueForPDCPSN12                          `aper:"valueExt"`
	IEExtension  *ProtocolExtensionContainerDRBStatusDL12ExtIEs `aper:"optional"`
}
