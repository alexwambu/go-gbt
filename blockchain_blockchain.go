package blockchain

type Blockchain struct {
	Transactions []*Transaction
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Transactions: []*Transaction{},
	}
}

func (bc *Blockchain) AddTransaction(tx *Transaction) {
	bc.Transactions = append(bc.Transactions, tx)
}
