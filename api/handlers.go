package api

import (
	"fmt"
	"net/http"
)

// GetIndex ...
func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Blockchain!")
}

// PostMine ...
func PostMine(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Mine")
}

// PostTransaction ...
func PostTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Transaction")
}
