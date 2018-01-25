package pkg

import "time"

// BlockHeader ...
type BlockHeader struct {
	Version    string
	PrevBlock  string // Hash of previous block
	MerkleRoot string // Hash of merkle root
	Time       time.Time
	Target     uint64
	Nonce      uint64
}

// Block ...
type Block struct {
	BlockHeader
	Transactions []Transaction // First transaction in this is the transaction to gift the miner
}
