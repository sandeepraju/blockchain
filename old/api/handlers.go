package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"

	"github.com/sandeepraju/blockchain/blockchain"
)

// GetIndex ...
func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Blockchain!")
}

// PostMine ...
func PostMine(w http.ResponseWriter, r *http.Request) {
	block, err := blockchain.DefaultBlockchain.Mine()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(block)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// PostTransaction ...
func PostTransaction(w http.ResponseWriter, r *http.Request) {

	mandatoryParams := []string{"sender", "recipient", "amount"}
	params := r.URL.Query()

	// Validate the parameters
	err := MustParams(mandatoryParams, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a transaction
	amount, err := decimal.NewFromString(params.Get("amount"))
	if err != nil {
		http.Error(w, "Invalid value in 'amount' parameter", http.StatusBadRequest)
		return
	}
	transaction := blockchain.Transaction{
		Sender:    params.Get("sender"),
		Recipient: params.Get("recipient"),
		Amount:    amount,
	}

	blockIndex, err := blockchain.DefaultBlockchain.AddTransaction(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("Transaction will be added to block number %d", blockIndex)
	jsonResponse, err := json.Marshal(map[string]string{"message": response})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
