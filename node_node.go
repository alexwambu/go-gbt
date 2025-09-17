package node

import (
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/rpc"
)

type Node struct {
	config     *Config
	server     *rpc.Server
	blockchain *Blockchain
}

func New(cfg *Config) (*Node, error) {
	return &Node{
		config:     cfg,
		blockchain: NewBlockchain(),
	}, nil
}

func (n *Node) Start() error {
	// Start blockchain mining
	n.blockchain.Start()

	// Create RPC server
	srv := rpc.NewServer()
	n.server = srv

	// Register ETH API
	if err := srv.RegisterName("eth", &EthAPI{
		chainID:    n.config.ChainID,
		blockchain: n.blockchain,
	}); err != nil {
		return fmt.Errorf("failed to register eth API: %v", err)
	}

	addr := fmt.Sprintf("%s:%d", n.config.RPCAddr, n.config.RPCPort)
	log.Printf("GBTNetwork node running at %s\n", addr)

	go func() {
		if err := http.ListenAndServe(addr, srv); err != nil {
			log.Fatalf("RPC server failed: %v", err)
		}
	}()

	return nil
}

func (n *Node) Stop() {
	if n.blockchain != nil {
		n.blockchain.Stop()
	}
	if n.server != nil {
		n.server.Stop()
	}
}

type EthAPI struct {
	chainID    int
	blockchain *Blockchain
}

func (api *EthAPI) ChainId() (*big.Int, error) {
	return big.NewInt(int64(api.chainID)), nil
}

func (api *EthAPI) BlockNumber() (string, error) {
	block := api.blockchain.CurrentBlock()
	return fmt.Sprintf("0x%x", block), nil
}

func (api *EthAPI) GetBalance(address string, block string) (string, error) {
	balance := big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(1e18))
	return fmt.Sprintf("0x%x", balance), nil
}
