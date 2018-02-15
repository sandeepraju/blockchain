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
	MerkleRoot string // Merkle hash tree root of transactions
	Time       time.Time
	Target     uint32
	Nonce      uint32
}

// Block ...
type Block struct {
	BlockHeader
	Transactions []Transaction
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
