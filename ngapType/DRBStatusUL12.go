package ngapType

import "github.com/acore2026/aper"

// Need to import "github.com/acore2026/aper" if it uses "aper"

type DRBStatusUL12 struct {
	ULCOUNTValue              COUNTValueForPDCPSN12                          `aper:"valueExt"`
	ReceiveStatusOfULPDCPSDUs *aper.BitString                                `aper:"sizeLB:1,sizeUB:2048,optional"`
	IEExtension               *ProtocolExtensionContainerDRBStatusUL12ExtIEs `aper:"optional"`
}
