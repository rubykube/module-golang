package orderbook

import (
	"math"
	"sync"
	"testing"
	"time"

	"github.com/valyala/fastrand"
)

func BenchmarkOrderbook(b *testing.B) {
	fabric := new(fabricator)

	fabric.chances.buy = 0.5
	fabric.chances.market = 0.5
	fabric.volume.min = 1000
	fabric.volume.max = 9000
	fabric.price.min = 5000
	fabric.price.max = 7000

	book := New()

	pipe := fabric.pipe(b.N, 1000000)

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
