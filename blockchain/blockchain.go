package blockchain

import "github.com/shopspring/decimal"

// Block ...
type Block struct {
	index        int32
	timestamp    int64
	transactions []Transaction
	proof        int32
	parent       string
}

// Transaction ...
type Transaction struct {
	sender    string
	recipient string
	amount    decimal.Decimal
}

// Blockchain ...
type Blockchain struct {
	lastBlock    int32
	chain        []Block
	transactions []Transaction
}

func (b *Blockchain) addBlock() error {
	return nil
}

func (b *Blockchain) addTransaction(transaction Transaction) (int32, error) {
	b.transactions = append(b.transactions, transaction)
	return b.lastBlock + 1, nil
}

// New ...
func New(genesis Block) *Blockchain {
	bc := Blockchain{}
	return &bc
}
