package pkg

import "time"

// Global constants

const (
	// TxVersion ...
	TxVersion = 0
)

// Enum declarations

const (
	// TxStatusUnconfirmed ...
	TxStatusUnconfirmed = iota
	// TxStatusConfirmed ...
	TxStatusConfirmed
)

const (
	// TxTypeSendToAddress ...
	TxTypeSendToAddress = iota
)

const (
	// TxLockTypeHeight ...
	TxLockTypeHeight = iota
	// TxLockTypeEpoch ...
	TxLockTypeEpoch
)

// TxLock ...
type TxLock struct {
	Type  uint8  // TxLockType*
	Value uint64 // Interpret this as block height or unix epoch based on the type
}

// TxInput ...
type TxInput struct {
	PrevTxOutput string // Hash of the previous transaction
	ScriptSig    string
}

// TxOutput ...
type TxOutput struct {
	Value        uint64
	ScriptPubKey string
}

// Transaction ...
type Transaction struct {
	Version uint8
	Hash    string
	Time    time.Time
	Type    uint8 // TxType*
	Status  uint8 // TxStatus*
	TxLock
	Inputs  []TxInput
	Outputs []TxOutput
}

// IsCoinbase ...
func (t *Transaction) IsCoinbase() bool {
	if len(t.Inputs) == 1 && t.Inputs[0].PrevTxOutput == "" {
		return true
	}

	return false
}
