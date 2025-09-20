package blockchain

import (
	"math/big"
)

type Transaction struct {
	From  string
	To    string
	Value *big.Int
	Data  []byte
	Nonce uint64
}
