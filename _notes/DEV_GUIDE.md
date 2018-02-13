# Developer Guide

- The transactions are paired and hashed until one single transaction remains (this is the merkle root).
- Merkle root is stored in the block header.
- Each block also stores the hash of the previous block's header (and thus the chain)
- 