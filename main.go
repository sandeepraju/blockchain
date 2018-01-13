package main

import "fmt"
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
	chain        []Block
	transactions []Transaction
}

func (b *Blockchain) addBlock() error {
	return nil
}

func (b *Blockchain) addTransaction(transaction Transaction) error {

	return nil
}

// ** Utils **
func hash(block Block) string {
	return ""
}

func main() {
	fmt.Println("Hello, Blockchain!")
}
