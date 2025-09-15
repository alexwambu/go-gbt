package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Default RPC URLs (can be overridden by ENV)
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		// fallback: check localhost and GBTNetwork
		rpcURL = "http://GBTNetwork:9636"
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("❌ Failed to connect to GBTNetwork RPC at %s: %v", rpcURL, err)
	}

	// Web endpoint for status
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		blockNumber, err := client.BlockNumber(r.Context())
		if err != nil {
			http.Error(w, "Failed to fetch block", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "✅ GBTNetwork RPC: %s | Latest Block: %d", rpcURL, blockNumber)
	})

	port := "9636"
	log.Printf("🚀 GBT node running on :%s using %s\n", port, rpcURL)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
