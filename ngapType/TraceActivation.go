package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type TraceActivation struct {
	NGRANTraceID                   NGRANTraceID
	InterfacesToTrace              InterfacesToTrace
	TraceDepth                     TraceDepth
	TraceCollectionEntityIPAddress TransportLayerAddress
	IEExtensions                   *ProtocolExtensionContainerTraceActivationExtIEs `aper:"optional"`
}
