package main

import (
	"fmt"
	"os"

	"github.com/alexwambu/go-gbt/node"
)

func main() {
	cfg := node.DefaultConfig()

	n, err := node.New(cfg)
	if err != nil {
		fmt.Println("Failed to create node:", err)
		os.Exit(1)
	}

	if err := n.Start(); err != nil {
		fmt.Println("Failed to start node:", err)
		os.Exit(1)
	}

	defer n.Stop()
	select {}
}
