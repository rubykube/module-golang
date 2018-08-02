## Implement simple Blockchain
A blockchain is a growing list of records, called blocks, which are linked
using cryptography. Blockchains which are readable by the public are widely
used by cryptocurrencies. Have fun implementing a blockchain using Golang.

In this excercise you need to implement **single node** blockchain daemon, which
persist blocks into sqlite database. **No network** communication is needed.

Start from simple blockchain implementation as described in the article in
Hints section and then proceed to saving blockchain into sqlite.

As a result you need to provide binary with following interface:

* `blockchain add <data>` — should add new block on top of blockchain and exit;
* `blockchain list` — should list blocks in blockchain and exit;
* `blockchain mine <difficulty>` (advanced) — start mining for new blocks with
  specified difficulty, mine blocks until program is interrupted. Print mined
  blocks to stdout as well as saving them to disk.

### Advanced

Implement mining algorithm with customized
[difficulty](https://en.bitcoin.it/wiki/Difficulty) setting.

### Hints
* [Basic prototype Blockchein (sha256)](https://jeiwan.cc/posts/building-blockchain-in-go-part-1/)

## Learning

* [How Does Blockchain Work](https://blockgeeks.com/guides/blockchain-consensus/)
* [Bitcoin - Beyond the Basics](https://www.youtube.com/watch?v=Dn6q9nveJbA)

### Consensus protocol For details please read the full [Raft
paper](https://ramcloud.stanford.edu/wiki/download/attachments/11370504/raft.pdf)
* [Blockchain Consensus Protocol](https://blockgeeks.com/guides/blockchain-consensus/)
* [The Raft Consensus Algorithm](https://raft.github.io/)
* [Raft](http://thesecretlivesofdata.com/raft/)

### Practical Byzantine Fault Tolerance One example of BFT in use is bitcoin,
a peer-to-peer digital cash system. The bitcoin network works in parallel to
generate a blockchain with proof-of-work allowing the system to overcome
Byzantine failures and reach a coherent global view of the system's state.
* [Byzantine fault tolerance](https://en.wikipedia.org/wiki/Byzantine_fault_tolerance)
* https://www.youtube.com/watch?v=_e4wNoTV3Gw
