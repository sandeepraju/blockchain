package blockchain

import "github.com/shopspring/decimal"

// Block ...
type Block struct {
	Index        int32
	Timestamp    int64
	Transactions []Transaction
	Proof        int32
	Parent       string
}

// Transaction ...
type Transaction struct {
	Sender    string
	Recipient string
	Amount    decimal.Decimal
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
