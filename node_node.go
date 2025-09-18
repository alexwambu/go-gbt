package node

import (
    "fmt"
    "github.com/alexwambu/go-gbt/core"
    "github.com/alexwambu/go-gbt/rpc"
)

type Node struct {
    BC     *core.Blockchain
    TxPool *core.TxPool
}

func New() *Node {
    return &Node{
        BC:     core.NewBlockchain(),
        TxPool: core.NewTxPool(),
    }
}

func (n *Node) Start() error {
    fmt.Println("Node started. RPC at :9636")
    server := &rpc.RPCServer{TxPool: n.TxPool}
    go server.Serve(":9636")
    select {} // keep alive
}
