package ngapType

// Need to import "github.com/acore2026/aper" if it uses "aper"

type PagingAttemptCount struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:16"`
}
