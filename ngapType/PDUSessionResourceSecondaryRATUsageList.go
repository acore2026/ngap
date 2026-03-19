package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct PDUSessionResourceSecondaryRATUsageList */
/* PDUSessionResourceSecondaryRATUsageItem */
type PDUSessionResourceSecondaryRATUsageList struct {
	List []PDUSessionResourceSecondaryRATUsageItem `aper:"valueExt,sizeLB:1,sizeUB:256"`
}
