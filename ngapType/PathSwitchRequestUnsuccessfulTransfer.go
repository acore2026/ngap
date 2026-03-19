package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PathSwitchRequestUnsuccessfulTransfer struct {
	Cause        Cause                                                                  `aper:"valueLB:0,valueUB:5"`
	IEExtensions *ProtocolExtensionContainerPathSwitchRequestUnsuccessfulTransferExtIEs `aper:"optional"`
}
