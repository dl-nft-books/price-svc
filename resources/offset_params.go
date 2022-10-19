package resources

type OffsetPageParams struct {
	Limit      uint64 `page:"limit" default:"15"`
	PageNumber uint64 `page:"number"`
}
