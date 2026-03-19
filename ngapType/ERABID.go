package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type ERABID struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:15"`
}
