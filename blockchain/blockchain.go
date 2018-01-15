package blockchain

import (
	"math/rand"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
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
	id           uuid.UUID
	chain        []Block
	transactions []Transaction
}

// AddTransaction adds the transaction into the current set of transactions
// and returns the block id into which the transactions will eventually be added to
func (b *Blockchain) AddTransaction(transaction Transaction) (uint64, error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	return b.unsafeAddTransaction(transaction)
}

func (b *Blockchain) unsafeAddTransaction(transaction Transaction) (uint64, error) {
	b.transactions = append(b.transactions, transaction)
	return uint64(len(b.chain)), nil
}

// Mine ...
func (b *Blockchain) Mine() (*Block, error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.unsafeMine()
}

func (b *Blockchain) unsafeMine() (*Block, error) {
	previousBlock := b.chain[len(b.chain)-1]
	nextProof := b.proofOfWork(previousBlock.Proof)

	transaction := Transaction{
		Sender:    "0", // signifies that this transaction is for mining
		Recipient: b.id.String(),
		Amount:    decimal.NewFromFloat(1.0), // TODO: how did we come up with this?
	}
	b.unsafeAddTransaction(transaction)

	previousHash := Hash(previousBlock)

	newBlock := Block{
		Index:        uint64(len(b.chain)),
		Timestamp:    time.Now().Unix(),
		Transactions: b.transactions,
		Proof:        nextProof,
		Parent:       previousHash,
	}

	// Empty the current transactions
	b.transactions = make([]Transaction, 0)
	b.chain = append(b.chain, newBlock)

	return &newBlock, nil
}

func (b *Blockchain) proofOfWork(previousProof uint64) uint64 {
	// TODO: implement this
	return 0
}

// New ...
func New(genesis Block) *Blockchain {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err.Error())
	}

	chain := make([]Block, 0)
	chain = append(chain, genesis)

	return &Blockchain{
		lock:         sync.RWMutex{},
		id:           id,
		chain:        chain,
		transactions: make([]Transaction, 0),
	}
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
