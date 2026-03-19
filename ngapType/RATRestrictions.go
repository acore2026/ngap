package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct RATRestrictions */
/* RATRestrictionsItem */
type RATRestrictions struct {
	List []RATRestrictionsItem `aper:"valueExt,sizeLB:1,sizeUB:16"`
}
