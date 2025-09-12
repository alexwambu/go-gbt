package main

import (
	"os"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
)

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(
		log.LvlInfo,
		log.StreamHandler(os.Stderr, log.TerminalFormat(true)),
	))

	cfg := node.DefaultConfig
	cfg.Name = "GBTNetwork"
	cfg.HTTPHost = "0.0.0.0"
	cfg.HTTPPort = 9636

	stack, err := node.New(&cfg)
	if err != nil {
		utils.Fatalf("Failed to create GBTNetwork node: %v", err)
	}

	ethConfig := eth.DefaultConfig
	ethConfig.NetworkId = 999
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return eth.New(ctx, &ethConfig)
	}); err != nil {
		utils.Fatalf("Failed to register GBTNetwork service: %v", err)
	}

	if err := stack.Start(); err != nil {
		utils.Fatalf("Failed to start GBTNetwork node: %v", err)
	}

	stack.Wait()
}
