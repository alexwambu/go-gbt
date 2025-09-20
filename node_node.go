package node

import "gbt/blockchain"

type Node struct {
	Chain *blockchain.Blockchain
}

func NewNode(chain *blockchain.Blockchain) *Node {
	return &Node{Chain: chain}
}

func (n *Node) AddTransaction(tx *blockchain.Transaction) {
	n.Chain.AddTransaction(tx)
}
