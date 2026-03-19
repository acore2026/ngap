package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct QosFlowAddOrModifyRequestList */
/* QosFlowAddOrModifyRequestItem */
type QosFlowAddOrModifyRequestList struct {
	List []QosFlowAddOrModifyRequestItem `aper:"valueExt,sizeLB:1,sizeUB:64"`
}
