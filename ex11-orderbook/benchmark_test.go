package orderbook

import (
	"math"
	"sync"
	"testing"
	"time"

	"github.com/valyala/fastrand"
)

func BenchmarkOrderbook(b *testing.B) {
	fabricAsk := new(fabricator)
	fabricBid := new(fabricator)
	fabricBench := new(fabricator)

	fabricAsk.chances.buy = 0
	fabricAsk.chances.market = 0
	fabricAsk.volume.min = 1000
	fabricAsk.volume.max = 9000
	fabricAsk.price.min = 6001
	fabricAsk.price.max = 7000

	fabricBid.chances.buy = 1
	fabricBid.chances.market = 0
	fabricBid.volume.min = 1000
	fabricBid.volume.max = 9000
	fabricBid.price.min = 5001
	fabricBid.price.max = 6000

	fabricBench.chances.buy = 0.5
	fabricBench.chances.market = 0.5
	fabricBench.volume.min = 1000
	fabricBench.volume.max = 9000
	fabricBench.price.min = 5000
	fabricBench.price.max = 7000

	book := New()

	bids := make([]Order, 80000)
	asks := make([]Order, 80000)

	fabricBid.fabricate(bids, 1000001)
	fabricAsk.fabricate(asks, 1000002)

	for _, order := range bids {
		book.Match(&order)
	}

	for _, order := range asks {
		book.Match(&order)
	}

	pipe := fabricBench.pipe(b.N, 1000000)
	started := time.Now()
	{
		b.ResetTimer()
		for batch := range pipe {
			for _, order := range batch {
				book.Match(&order)
			}
		}
		b.StopTimer()
	}

	finished := time.Now()
	elapsed := finished.Sub(started).Seconds()

	if elapsed > 1 || b.N > 1000000 {
		b.Logf(
			"%d: %.2f orders/sec",
			b.N,
			float64(b.N)/elapsed,
		)

	}
}

type fabricator struct {
	chances struct {
		buy    float64
		market float64
	}

	volume struct {
		min float64
		max float64
	}

	price struct {
		min float64
		max float64
	}

	rng fastrand.RNG
}

func (fabricator *fabricator) pipe(
	count int,
	chunk int,
) chan []Order {
	pipe := make(chan []Order)

	ready := sync.WaitGroup{}
	ready.Add(1)

	fabricated := false

	var buffer [2][]Order

	buffer[0] = make([]Order, chunk)
	buffer[1] = make([]Order, chunk)

	go func() {
		steps := int(math.Ceil(float64(count) / float64(chunk)))

		for generation := 0; generation < steps; generation++ {
			fabricator.fabricate(buffer[generation%2], generation)

			if !fabricated {
				ready.Done()

				fabricated = true
			}

			pipe <- buffer[generation%2]
		}

		close(pipe)
	}()

	ready.Wait()

	return pipe
}

func (fabricator *fabricator) fabricate(
	buffer []Order,
	generation int,
) {
	count := len(buffer)

	for i := 0; i < count; i++ {
		order := &buffer[i]

		order.ID = i + count*generation

		if fabricator.float64(0, 1) > fabricator.chances.buy {
			order.Side = SideBid
		} else {
			order.Side = SideAsk
		}

		if fabricator.float64(0, 1) > fabricator.chances.market {
			order.Kind = KindMarket
		} else {
			order.Kind = KindLimit
		}

		volume := fabricator.float64(
			fabricator.volume.min,
			fabricator.volume.max,
		)

		price := fabricator.float64(
			fabricator.price.min,
			fabricator.price.max,
		)

		order.Volume = uint64(volume * 10000)
		order.Price = uint64(price * 10000)
	}
}

func (fabricator *fabricator) float64(min, max float64) float64 {
	float := float64(fabricator.rng.Uint32()) / math.MaxUint32

	return float*(max-min) + min
}
