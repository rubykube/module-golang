package orderbook

type Side int8

const (
	SideBid Side = 1
	SideAsk Side = 2
)

func (side Side) String() string {
	switch side {
	case SideBid:
		return "BID"
	case SideAsk:
		return "ASK"
	}

	return "UNKNOWN"
}

type Kind int8

const (
	KindMarket Kind = 1
	KindLimit  Kind = 2
)

func (kind Kind) String() string {
	switch kind {
	case KindMarket:
		return "MARKET"
	case KindLimit:
		return "LIMIT"
	}

	return "UNKNOWN"
}

type Order struct {
	ID int

	Side Side
	Kind Kind

	Volume uint64
	Price  uint64
}
