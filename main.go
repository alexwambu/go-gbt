package main

import (
	"fmt"
	"log"
	"net/http"

	"gbt/blockchain"
	"gbt/node"
	"gbt/rpc"
)

func main() {
	// Initialize blockchain
	chain := blockchain.NewBlockchain()

	// Start node
	n := node.NewNode(chain)

	// Start RPC server
	rpcServer := rpc.NewRPCServer(n)
	http.Handle("/rpc", rpcServer)

	fmt.Println("🚀 GBTNetwork running at http://localhost:9636")
	log.Fatal(http.ListenAndServe(":9636", nil))
}
