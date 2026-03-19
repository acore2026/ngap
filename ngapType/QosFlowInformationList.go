package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct QosFlowInformationList */
/* QosFlowInformationItem */
type QosFlowInformationList struct {
	List []QosFlowInformationItem `aper:"valueExt,sizeLB:1,sizeUB:64"`
}
