package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type TargetNGRANNodeToSourceNGRANNodeTransparentContainer struct {
	RRCContainer RRCContainer
	IEExtensions *ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs `aper:"optional"`
}
