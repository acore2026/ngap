package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct UEPresenceInAreaOfInterestList */
/* UEPresenceInAreaOfInterestItem */
type UEPresenceInAreaOfInterestList struct {
	List []UEPresenceInAreaOfInterestItem `aper:"valueExt,sizeLB:1,sizeUB:64"`
}
