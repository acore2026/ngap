package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct CellIDCancelledEUTRA */
/* CellIDCancelledEUTRAItem */
type CellIDCancelledEUTRA struct {
	List []CellIDCancelledEUTRAItem `aper:"valueExt,sizeLB:1,sizeUB:65535"`
}
