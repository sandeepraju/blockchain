# Notes

## [Design] Open Questions

- How do we get peer list?
- HTTP API (in lieu of Bitcoin RPC)?
- How do we persist the blockchain to disk?
- Proof of work?
- Signing & Checksum
- Conflict resolution
- Genesis block
- Coinbase for generation transaction
- Support for standard bitcoin address
- Key format?
- JSON based data serialization?

## [Implementation] Open Questions

- Go LevelDB (or [BuntDB](https://github.com/tidwall/buntdb))?
- YAML config file

## Other design considerations

- Lightweight vs Full nodes
- Storing unspents

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