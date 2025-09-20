package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
)

func main() {
	// Basic logging
	log.Root().SetHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))

	// Data directory (can be overridden by env or CLI later)
	dataDir := "./data"
	if env := os.Getenv("DATADIR"); env != "" {
		dataDir = env
	}
	if err := os.MkdirAll(dataDir, 0o700); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create datadir: %v\n", err)
		os.Exit(1)
	}

	// Node configuration
	cfg := &node.Config{
		Name:        "GBTNetwork",
		HTTPHost:    "0.0.0.0",
		HTTPPort:    9636,
		DataDir:     dataDir,
		HTTPModules: []string{"eth", "net", "web3", "personal"},
	}

	// Create node
	stack, err := node.New(cfg)
	if err != nil {
		utils.Fatalf("Failed to create node: %v", err)
	}

	// Eth service config
	ethConfig := eth.DefaultConfig
	ethConfig.NetworkId = 999

	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return eth.New(ctx, &ethConfig)
	}); err != nil {
		utils.Fatalf("Failed to register eth service: %v", err)
	}

	// Initialize chain from genesis.json on first run (mimic `geth init`)
	genesisPath := filepath.Join(".", "genesis.json")
	if _, err := os.Stat(filepath.Join(dataDir, "geth")); os.IsNotExist(err) {
		// Try init (run geth init behavior by calling the geth init helper)
		// When using the node package, a typical way is to call core.GenerateGenesis
		// But easiest/most robust: call the external CLI binary if available
		// (If you want purely in-process init, see go-ethereum's "core" package docs.)
		fmt.Printf("Data directory not initialized — please run: ./beth --datadir %s init %s\n", dataDir, genesisPath)
		// We still continue to start node; if not initialized eth service will create empty chain.
	}

	// Start node
	if err := stack.Start(); err != nil {
		utils.Fatalf("Failed to start node: %v", err)
	}
	defer stack.Stop()

	fmt.Println("GBTNetwork node started. HTTP RPC listening on port 9636")
	select {}
}
