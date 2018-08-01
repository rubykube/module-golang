# Resolving Sleeping barber problem with go
The problem is simple and exposed there: http://en.wikipedia.org/wiki/Sleeping_barber_problem

## This is very simple:

- a barber manage a shop with a waiting room
- a waiting room has got some seat
- some clients enter the shop, try to find a free seat
- free seat: ok, wait for barber
- no seat: go out... (and retry later if you want)
- if the barber has no clients, he decide to sleep
- if client enter and see that the barber is sleeping, wake him

For the resolution, and to be more expressive, clients that don't find a seat to retry later.

In many langages, this problem can be solved by using mutex. But in Go the problem is a bit easier, we can use:

- channels
- select statement
