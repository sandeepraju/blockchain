# Developer Guide

Ref: https://bitcoin.org/en/developer-guide#block-chain-overview

- The transactions are paired and hashed until one single transaction remains (this is the merkle root).
- Merkle root is stored in the block header.
- Each block also stores the hash of the previous block's header (and thus the chain)
- Transactions are chained to together (bitcoins actually move from transactions to transactions)
- Each transaction spends the satoshis previously received in one or more earlier transactions, so the input of one transaction is the output of a previous transaction.
- TX means a transaction and UTXO means unspent transaction output.
- A single transaction can create multiple outputs (esp. when spending to multiple addresses), but more importantly, each output of a particular transaction can only be used as an input once in the block chain -- this is what prevents the bitcoin from being double spent.
- Every output is tied to TXIDs (transaction identifiers), which are hashes of signed transactions.
- Because each output of a particular transaction can only be spent once, the outputs of all transactions included in the block chain can be categorized as either:
    - Unspent transactions (UTXOs)
    - Spent transaction outputs
- For any transaction (payment) to be valid and be accepted by the blockchain, it must only use UTXOs as the input.
- Ignoring the coinbase transactions, if the value of transactions output exceeds its input, the transaction will be rejected.
- If the inputs exceed the value of outputs, any difference in value maybe claimed as a transaction fee by the miner (who creates the block containing the transaction)