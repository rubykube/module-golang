package orderbook

type Trade struct {
	Bid    *Order
	Ask    *Order
	Volume uint64
	Price  uint64
}
