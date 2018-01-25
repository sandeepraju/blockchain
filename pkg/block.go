package pkg

import "time"

const (
	// BlockVersion ...
	BlockVersion = 0
)

// BlockHeader ...
type BlockHeader struct {
	Version    string
	PrevBlock  string // Hash of previous block
	MerkleRoot string // Hash of merkle root
	Time       time.Time
	Target     uint32
	Nonce      uint32
}

// Block ...
type Block struct {
	BlockHeader
	Transactions []Transaction // First transaction in this is the transaction to gift the miner
}

// IsGenesis ...
func (b *Block) IsGenesis() bool {
	if b.PrevBlock == "" {
		return true
	}

	return false
}

// NewBlock ...
func NewBlock() *Block {
	return &Block{}
}

// NewGenesisBlock ...
func NewGenesisBlock() *Block {
	return &Block{}
}
