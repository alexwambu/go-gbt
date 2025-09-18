package core

import (
    "math/big"
)

type Block struct {
    Number     *big.Int
    Hash       []byte
    ParentHash []byte
    Txs        []*Transaction
}

type Blockchain struct {
    Blocks []*Block
}

func NewBlockchain() *Blockchain {
    return &Blockchain{Blocks: []*Block{}}
}

func (bc *Blockchain) AddBlock(b *Block) {
    bc.Blocks = append(bc.Blocks, b)
}
