package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Simple JSON-RPC handler
func rpcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// A fake latest block number
	resp := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"result":  "999", // pretend block number
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	port := 9636
	http.HandleFunc("/", rpcHandler)
	fmt.Printf("🚀 GBTNetwork RPC running on http://localhost:%d\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
