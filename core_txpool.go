package core

import (
    "sync"
)

type Transaction struct {
    From  string
    To    string
    Value int64
    Data  []byte
}

type TxPool struct {
    mu   sync.Mutex
    pool []*Transaction
}

func NewTxPool() *TxPool {
    return &TxPool{}
}

func (p *TxPool) Add(tx *Transaction) {
    p.mu.Lock()
    defer p.mu.Unlock()
    p.pool = append(p.pool, tx)
}

func (p *TxPool) Pending() []*Transaction {
    return p.pool
}
