package rpc

import (
	"encoding/json"
	"net/http"

	"gbt/blockchain"
	"gbt/node"
)

type RPCServer struct {
	node *node.Node
}

func NewRPCServer(n *node.Node) *RPCServer {
	return &RPCServer{node: n}
}

func (r *RPCServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var body map[string]interface{}
	json.NewDecoder(req.Body).Decode(&body)

	method := body["method"].(string)

	switch method {
	case "eth_sendRawTransaction":
		txData := body["params"].([]interface{})[0].(string)
		tx := &blockchain.Transaction{From: "user", To: "miner", Value: nil, Data: []byte(txData)}
		r.node.AddTransaction(tx)
		json.NewEncoder(w).Encode(map[string]string{"result": "tx accepted"})
	default:
		json.NewEncoder(w).Encode(map[string]string{"error": "method not found"})
	}
}
