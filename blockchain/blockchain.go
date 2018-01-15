package blockchain

import (
	"math/rand"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

// DefaultBlockchain ...
var DefaultBlockchain *Blockchain

func init() {
	// Seed the rand
	rand.Seed(time.Now().Unix())
	DefaultBlockchain = NewWithGenesis()
}

// Block ...
type Block struct {
	Index        uint64
	Timestamp    int64
	Transactions []Transaction
	Proof        uint64
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
	lock         sync.RWMutex
	lastBlock    uint64
	chain        []Block
	transactions []Transaction
}

// AddTransaction ...
func (b *Blockchain) AddTransaction(transaction Transaction) (uint64, error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.transactions = append(b.transactions, transaction)
	return b.lastBlock + 1, nil
}

// New ...
func New(genesis Block) *Blockchain {
	bc := Blockchain{}
	return &bc
}

// NewWithGenesis ...
func NewWithGenesis() *Blockchain {
	// Generate a genesis block
	genesis := Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Transactions: make([]Transaction, 0),
		Proof:        rand.Uint64(),
		Parent:       "",
	}

	return New(genesis)
}
