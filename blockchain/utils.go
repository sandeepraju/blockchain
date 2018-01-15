package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

// Hash ...
func Hash(block Block) (string, error) {
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	err := encoder.Encode(block)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	hash.Write(b.Bytes())
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
