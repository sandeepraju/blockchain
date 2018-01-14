package main

import (
	"fmt"
	"time"

	"github.com/sandeepraju/blockchain/blockchain"
)

func main() {
	bc := blockchain.New(blockchain.Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Transactions: nil,
		Proof:        0,
		Parent:       "",
	})

	bc.Run()
	fmt.Println("Hello, Blockchain!")
}
