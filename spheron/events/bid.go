package events

type BidPlaced struct {
	ID       uint64
	BidPrice uint64
	Bidder   string
}
