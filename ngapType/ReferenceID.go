package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type ReferenceID struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:64"`
}
