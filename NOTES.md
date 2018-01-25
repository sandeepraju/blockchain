# Notes

## [Design] Open Questions

- How do we get peer list?
- How to seed the network? -- can there be only one node?! who starts the genesis block and coinbase transaction.
- How do we persist the blockchain to disk? (periodic backups?)
- Proof of work?
- Signing & Checksum
- Conflict resolution
- Genesis block
- Coinbase for generation transaction
- Support for standard bitcoin address
- Key format?
- JSON based data serialization?
- Node wallet to save mined earnings (?)
- Should there be a limit for mempool (unconfirmed transactions)? Why is there a limit?
- How do we handle blocks that don't make it to the main chain (_'side chain'_ blocks)
- Lightweight vs Full nodes
- Storing unspents
- How to work out the difficulty & target for our toy implementation?
- Change in block creation fee
- How do we store UTXOs (unspent transaction outputs)

## [Implementation] Open Questions

- Go LevelDB (or [BuntDB](https://github.com/tidwall/buntdb))?
- YAML config file
- We need a simple client to act as a wallet for proof of concept
- Actor based design?
- HTTP vs gRPC for interfacing?

## Overview

### When the node starts up

- Connect to the _'peer discovery service'_ and get a list of peers to connect to.
    - __Q:__ How many peers can one node connect to?
- Load a local copy of blockchain from disk (if exists)
- Initiate a sync to make sure the local copy of the blockchain is up to date
- Start listening to the following requests
    - Request to process a transaction
    - Request to verify a new block
- In parallel, perform _'work'_. When the work ends, consolidate all the valid transactions and mine the block.
    - __Q:__ How to prioritize the transactions
    - __Q:__ Should there be a block size? (why / why not?)
    - __Q:__ Transaction age treshold (How long can a transaction be in queue to be added to a block?)
- Add the newly mined block to the blockchain
    - __Q:__ How to resolve conflicts in blockchain order?
    - __Q:__ How to break ties in mining?
- Broadcast the newly mined to other peers

### When a new transaction request comes in

-

### When a block verification request comes in

-