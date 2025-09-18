package main

import (
    "fmt"
    "os"
    "github.com/alexwambu/go-gbt/node"
)

func main() {
    fmt.Println("🚀 Starting GBTNetwork Node...")
    n := node.New()
    if err := n.Start(); err != nil {
        fmt.Println("❌ Failed to start node:", err)
        os.Exit(1)
    }
}
