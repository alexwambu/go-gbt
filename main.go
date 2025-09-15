package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Simple JSON-RPC style handler
func rpcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Example static block height
	resp := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"result":  "999", // fake block number for now
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	port := 9636
	http.HandleFunc("/", rpcHandler)
	fmt.Printf("🚀 GBTNetwork node running at http://0.0.0.0:%d\n", port)

	// Important: bind to 0.0.0.0 for Codespaces + Render
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
