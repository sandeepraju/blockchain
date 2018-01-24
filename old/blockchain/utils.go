package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

// HashBlock ...
func HashBlock(block Block) (string, error) {
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	err := encoder.Encode(block)
	if err != nil {
		return "", err
	}
	return Hash(b.Bytes())
}

// Hash ...
func Hash(data []byte) (string, error) {
	hash := sha256.New()
	hash.Write(data)
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
