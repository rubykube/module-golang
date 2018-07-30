# Key-Value persistent storage with replication

Task: implement dead-simple key-value specific storage for managing users
balances (replacement to redis) where key is UUID (16 bytes) and value is float
number (base64).

Supported operations:
- GET - return value of specified key
- SETX - set value to specified key only if value doesn't exist
- INCX - increment value of specified key only if value exists and return new
    value
- DECX - same as INCX but decrement.

The data must be replicated using Raft consensus (just use [hashicorp's
implementation](https://github.com/hashicorp/raft)), the data must be flushed
to disk after applying to follower nodes.

The service must provide TCP API.

## TCP Protocol specification:

### Request packet

- operation: 1 byte
- key: 16 bytes
- value: 8 bytes

### Response packet
- code: 1 byte
- value: 8 bytes

* Don't use boltdb as backend, it's unmaintained and just slow. Use
    [raft-fastlog](https://github.com/tidwall/raft-fastlog) as backend LogStore
    and StableStore

## Hints & Links

* [Go By Example: Mutexes](https://gobyexample.com/mutexes)
* [sync.RWMutex](https://golang.org/pkg/sync/#RWMutex)
* [raft implementation](https://github.com/hashicorp/raft)

