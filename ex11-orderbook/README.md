# Simple Orderbook

## Overview

Goal of this task is to write simple [orderbook][1] (i.e. orders matching
engine), benchmark it and provide some optimizations.

Typically, orderbook is an ordered list of buy (bids) and sell (asks) orders,
which are ordered by price and arrival time.

Orders that are located in orderbook are called resting orders.

Orderbook matches limit bid orders by scanning resting ask orders that have
sell price lower or equal to price from bid order. If matching order is found,
then orderbook produces trade which contains exchanged amount and ask price.
If bid order is completely fulfilled, then process interrupted, otherwise
process continues till bid order completely fulfilled or can't be fulfilled
because all ask prices for resting orders are greater than of bid. In that case
bid order became resting order and added to orderbook.

Market orders does not contain price, only amount. If incoming order is market
bid order, then it matches its amount against resting ask orders until:
a) incoming bid order is completely fulfilled or b) there are no more ask
orders left.

Exactly same algorithm applies for incoming ask orders, just mirrored.

When orderbook executes trade matching amount is removed from both bid and ask
orders and only three cases are possible: 1) either bid order is completely
fulfilled and process stops, 2) ask order is completely fulfilled and removed
from orderbook while process continues or 3) both bid and ask orders are
completely fulfilled, ask order removed from orderbook and process stops.

Orderbook always executes order using best price for incoming order.

For example, consider resting orders in orderbook:

```
#1 BID 1 $1000
#2 BID 2 $1500
#3 ASK 2 $2000
#4 ASK 1 $2500
```

Then, following incoming orders will be matched as:

* `LIMIT ASK 1 $900` will match order `#2` to produce `TRADE 1 $1500`;
* `LIMIT BID 1 $3000` will match order `#3` to produce `TRADE 1 $2000`;
* `MARKET ASK 3` will match first order `#2` and then `#1` to produce
  `TRADE 2 $1500`, `TRADE 1 $1000` in that order;
* `MARKET BID 3` will match first order `#3` and then `#4` to produce
  `TRADE 2 $2000`, `TRADE 1 $2500` in that order;

In other words, orderbook always prefer to match order which provides most
profitable price for incoming order.

## Implementation

Orderbook implementation should support two kind of orders: [market][2] and
[limit][3].

Orderbook should support partial fulfillment of orders.

Orderbook should operate completely in-memory and doesn't have any persistent
storage.

Orderbook should be provided as single go-lang package.

Orderbook implementation should provide following API methods:

* `New() -> *Orderbook`;

* `Match(*Order) -> ([]*Trade, *Order)` with following behavior:
  * it accepts order as single argument;
  * it returns two values: list of executed trades if any and rejected order if
    any;
  * it may reject order only and only if order is market order and there are no
    resting orders in orderbook that can fulfill it;
  * limit order can became resting if there are not enough other resting orders
    in orderbook that can fulfill it completely;
  * it should return slice of trades which represent executed trades between
    incoming order and resting orders in orderbook;
  * resting order should be kept in orderbook as long as it's not completely
    fulfilled; each partial fulfillment should decrease amount in order;

In this implementation orderbook can be non thread safe.

As a price and amount types you need to use single `uint64` value where last
4 decimal places represent fractional part of number, e.g.:

```
1      ->  0.0001
1000   ->  0.1000
10000  ->  1.0000
123456 -> 12.3456
```

## File structure

Project bootstrapped with `Trade` and `Order` structures, `Orderbook` structure
placeholders, tests and benchmarks.

You need to modify `orderbook.go` file.

To run tests use:

```bash
go test # OR
go test -failfast # to stop at first failing test
```

To run benchmarks use:

```bash
go test -bench . -benchmem # OR
go test -bench . -benchmem -run none # to skip tests
```

## Additional tasks

### Easy

Add order cancellation method `Cancel(ID) bool` to cancel resting order or
return false if no order found with specified ID.

### Hard

Optimize orderbook to be able to match more than ~6M orders per second (depends
on your hardware) in single thread.

### Nightmare

Make orderbook not only thread safe but matching orders faster when `Match()`
is called concurrently from several go-routines.

[1]: https://en.wikipedia.org/wiki/Order_book_(trading)
[2]: https://en.wikipedia.org/wiki/Order_(exchange)#Market_order
[3]: https://en.wikipedia.org/wiki/Order_(exchange)#Limit_order
