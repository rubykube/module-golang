package orderbook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimit_Bid_NoMatch(t *testing.T) {
	new(testcase).
		OrderLimit(SideAsk, 10000, 60000000).
		OrderLimit(SideBid, 10000, 10000000).
		Assert(t)
}

func TestLimit_Ask_NoMatch(t *testing.T) {
	new(testcase).
		OrderLimit(SideBid, 10000, 10000000).
		OrderLimit(SideAsk, 10000, 60000000).
		Assert(t)
}

func TestLimit_Bid_ExactMatch(t *testing.T) {
	new(testcase).
		OrderLimit(SideAsk, 10000, 60000000).
		OrderLimit(SideBid, 10000, 60000000).
		Trade(10000, 60000000).
		Assert(t)
}

func TestLimit_Ask_ExactMatch(t *testing.T) {
	new(testcase).
		OrderLimit(SideBid, 10000, 60000000).
		OrderLimit(SideAsk, 10000, 60000000).
		Trade(10000, 60000000).
		Assert(t)
}

func TestLimit_Bid_SamePricePoint(t *testing.T) {
	new(testcase).
		OrderLimit(SideBid, 15000, 60000000).
		OrderLimit(SideAsk, 25000, 60001000).
		OrderLimit(SideBid, 15000, 60002000).
		OrderLimit(SideBid, 10000, 60003000).
		Trade(15000, 60001000).
		Trade(10000, 60001000).
		Assert(t)
}

func TestLimit_Bid_PartialMatch(t *testing.T) {
	new(testcase).
		OrderLimit(SideAsk, 20000, 60000000).
		OrderLimit(SideBid, 10000, 60000000).
		OrderLimit(SideBid, 10000, 60000000).
		Trade(10000, 60000000).
		Trade(10000, 60000000).
		Assert(t)
}

func TestLimit_Bid_PartialMatch_Rest(t *testing.T) {
	new(testcase).
		OrderLimit(SideAsk, 14000, 60000000).
		OrderLimit(SideAsk, 15000, 60000000).
		OrderLimit(SideBid, 30000, 60000000).
		Trade(14000, 60000000).
		Trade(15000, 60000000).
		Assert(t)
}

func TestLimit_Ask_PartialMatch_Rest(t *testing.T) {
	new(testcase).
		OrderLimit(SideBid, 14000, 60000000).
		OrderLimit(SideBid, 15000, 60000000).
		OrderLimit(SideAsk, 30000, 60000000).
		Trade(14000, 60000000).
		Trade(15000, 60000000).
		Assert(t)
}

func TestLimit_Bid_PartialMatch_DifferentPrices(t *testing.T) {
	new(testcase).
		OrderLimit(SideAsk, 15000, 60000000).
		OrderLimit(SideAsk, 15000, 61000000).
		OrderLimit(SideAsk, 15000, 61100000).
		OrderLimit(SideBid, 45000, 62000000).
		Trade(15000, 60000000).
		Trade(15000, 61000000).
		Trade(15000, 61100000).
		Assert(t)
}

func TestLimit_Ask_PartialMatch_DifferentPrices(t *testing.T) {
	new(testcase).
		OrderLimit(SideBid, 15000, 60000000).
		OrderLimit(SideBid, 15000, 61000000).
		OrderLimit(SideBid, 15000, 61100000).
		OrderLimit(SideAsk, 45000, 60000000).
		Trade(15000, 61100000).
		Trade(15000, 61000000).
		Trade(15000, 60000000).
		Assert(t)
}

func TestLimit_Bid_ExactMatch_Search(t *testing.T) {
	new(testcase).
		OrderLimit(SideAsk, 10000, 61000000).
		OrderLimit(SideAsk, 10000, 60000000).
		OrderLimit(SideBid, 10000, 60000000).
		Trade(10000, 60000000).
		Assert(t)
}

func TestMarket_Bid(t *testing.T) {
	new(testcase).
		OrderLimit(SideAsk, 10000, 61000000).
		OrderLimit(SideAsk, 10000, 60000000).
		OrderMarket(SideBid, 10000).
		Trade(10000, 60000000).
		Assert(t)
}

func TestMarket_Bid_Reject(t *testing.T) {
	new(testcase).
		OrderMarket(SideBid, 10000).
		Reject(1, 10000).
		Assert(t)
}

func TestMarket_Bid_Reject_PartialMatch(t *testing.T) {
	new(testcase).
		OrderLimit(SideBid, 5000, 60000000).
		OrderLimit(SideAsk, 10000, 60001000).
		OrderMarket(SideBid, 15000).
		Trade(10000, 60001000).
		Reject(3, 5000).
		Assert(t)
}

func TestMarket_Bid_AfterPreviousPositionClosed(t *testing.T) {
	new(testcase).
		OrderLimit(SideAsk, 8000, 55400000).
		OrderLimit(SideBid, 8000, 55400000).
		OrderLimit(SideAsk, 2000, 55630000).
		OrderMarket(SideBid, 8000).
		Trade(8000, 55400000).
		Trade(2000, 55630000).
		Reject(4, 6000).
		Assert(t)
}

func TestMarket_Ask_AfterPreviousPositionClosed(t *testing.T) {
	new(testcase).
		OrderLimit(SideBid, 8000, 55630000).
		OrderLimit(SideAsk, 8000, 55630000).
		OrderLimit(SideBid, 2000, 55400000).
		OrderMarket(SideAsk, 8000).
		Trade(8000, 55630000).
		Trade(2000, 55400000).
		Reject(4, 6000).
		Assert(t)
}

func TestMarket_Bid_LowestPriceFirst(t *testing.T) {
	new(testcase).
		OrderLimit(SideAsk, 1000, 50000000).
		OrderLimit(SideAsk, 5000, 51000000).
		OrderMarket(SideBid, 3000).
		Trade(1000, 50000000).
		Trade(2000, 51000000).
		Assert(t)
}

func TestMarket_Bid_LowestPriceFirst_SamePricePoint(t *testing.T) {
	new(testcase).
		OrderLimit(SideAsk, 1000, 50001000).
		OrderLimit(SideAsk, 5000, 50002000).
		OrderMarket(SideBid, 3000).
		Trade(1000, 50001000).
		Trade(2000, 50002000).
		Assert(t)
}

func TestMarket_Ask_HighestPriceFirst(t *testing.T) {
	new(testcase).
		OrderLimit(SideBid, 1000, 50000000).
		OrderLimit(SideBid, 5000, 51000000).
		OrderMarket(SideAsk, 3000).
		Trade(3000, 51000000).
		Assert(t)
}

func TestMarket_Ask_HighestPriceFirst_SamePricePoint(t *testing.T) {
	new(testcase).
		OrderLimit(SideBid, 1000, 50001000).
		OrderLimit(SideBid, 5000, 50002000).
		OrderMarket(SideAsk, 3000).
		Trade(3000, 50002000).
		Assert(t)
}

type testcase struct {
	Orders  []*Order
	Trades  []*Trade
	Rejects []*Order
}

func (testcase *testcase) OrderLimit(
	side Side,
	amount uint64,
	price uint64,
) *testcase {
	order := &Order{
		ID:     len(testcase.Orders) + 1,
		Side:   side,
		Kind:   KindLimit,
		Volume: amount,
		Price:  price,
	}

	testcase.Orders = append(testcase.Orders, order)

	return testcase
}

func (testcase *testcase) OrderMarket(
	side Side,
	amount uint64,
) *testcase {
	order := &Order{
		ID:     len(testcase.Orders) + 1,
		Side:   side,
		Kind:   KindMarket,
		Volume: amount,
	}

	testcase.Orders = append(testcase.Orders, order)

	return testcase
}

func (testcase *testcase) Trade(amount uint64, price uint64) *testcase {
	testcase.Trades = append(
		testcase.Trades,
		&Trade{
			Volume: amount,
			Price:  price,
		},
	)

	return testcase
}

func (testcase *testcase) Reject(uuid int, amount uint64) *testcase {
	testcase.Rejects = append(
		testcase.Rejects,
		&Order{
			ID:     uuid,
			Volume: amount,
		},
	)

	return testcase
}

func (testcase *testcase) Assert(t *testing.T) {
	test := assert.New(t)

	book := New()

	var (
		trades  []*Trade
		rejects []*Order
	)

	for _, order := range testcase.Orders {
		traded, reject := book.Match(order)

		trades = append(trades, traded...)

		if reject != nil {
			rejects = append(rejects, reject)
		}
	}

	test.Len(trades, len(testcase.Trades), "trades number mismatch")

	for i := 0; i < len(trades); i++ {
		test.Equal(
			int(testcase.Trades[i].Volume),
			int(trades[i].Volume),
			"trade %d volume mismatch", i,
		)

		test.Equal(
			int(testcase.Trades[i].Price),
			int(trades[i].Price),
			"trade %d price mismatch", i,
		)
	}

	test.Len(rejects, len(testcase.Rejects), "rejects number mismatch")

	for i := 0; i < len(rejects); i++ {
		test.Equal(
			testcase.Rejects[i].ID,
			rejects[i].ID,
			"reject %d ID mismatch", i,
		)

		test.Equal(
			int(testcase.Rejects[i].Volume),
			int(rejects[i].Volume),
			"reject %d volume mismatch", i,
		)
	}
}
