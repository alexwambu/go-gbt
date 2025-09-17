package node

import (
	"log"
	"math/big"
	"sync"
	"time"
)

// Blockchain is a minimal in-memory blockchain
type Blockchain struct {
	mu          sync.RWMutex
	blockNumber *big.Int
	quit        chan struct{}
}

// NewBlockchain initializes blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{
		blockNumber: big.NewInt(0),
		quit:        make(chan struct{}),
	}
}

// Start simulates mining new blocks
func (bc *Blockchain) Start() {
	go func() {
		ticker := time.NewTicker(5 * time.Second) // mine every 5s
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				bc.mu.Lock()
				bc.blockNumber.Add(bc.blockNumber, big.NewInt(1))
				log.Printf("Mined new block #%s\n", bc.blockNumber.String())
				bc.mu.Unlock()
			case <-bc.quit:
				return
			}
		}
	}()
}

// Stop blockchain mining
func (bc *Blockchain) Stop() {
	close(bc.quit)
}

// CurrentBlock returns current block number
func (bc *Blockchain) CurrentBlock() *big.Int {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	return new(big.Int).Set(bc.blockNumber)
}
