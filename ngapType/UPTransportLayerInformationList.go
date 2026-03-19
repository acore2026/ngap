package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct UPTransportLayerInformationList */
/* UPTransportLayerInformationItem */
type UPTransportLayerInformationList struct {
	List []UPTransportLayerInformationItem `aper:"valueExt,sizeLB:1,sizeUB:3"`
}
