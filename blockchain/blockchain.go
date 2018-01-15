package blockchain

import (
	"fmt"
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
	nextProof, err := b.proofOfWork(previousBlock.Proof)
	if err != nil {
		return nil, err
	}

	transaction := Transaction{
		Sender:    "0", // signifies that this transaction is for mining
		Recipient: b.id.String(),
		Amount:    decimal.NewFromFloat(1.0), // TODO: how did we come up with this?
	}
	b.unsafeAddTransaction(transaction)

	previousHash, err := HashBlock(previousBlock)
	if err != nil {
		return nil, err
	}

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

func (b *Blockchain) proofOfWork(previousProof uint64) (uint64, error) {
	// Simple Proof of Work Algorithm:
	//   - Find a number p' such that hash(pp') contains leading 4 zeroes, where p is the previous p'
	//   - p is the previous proof, and p' is the new proof
	proof := uint64(0)
	for {
		ok, err := b.verifyProof(previousProof, proof)
		if err != nil {
			return 0, err
		}

		// TODO: refactor this; looks really odd!
		if ok {
			return proof, nil
		}

		proof++
	}
}

func (b *Blockchain) verifyProof(previousProof uint64, currentProof uint64) (bool, error) {
	hash, err := Hash([]byte(fmt.Sprintf("%d:%d", previousProof, currentProof)))
	if err != nil {
		return false, err
	}
	return (hash[:4] == "0000"), nil
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
