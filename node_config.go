package node

type Config struct {
	RPCAddr string
	RPCPort int
	ChainID int
}

func DefaultConfig() *Config {
	return &Config{
		RPCAddr: "0.0.0.0",
		RPCPort: 9636,
		ChainID: 999,
	}
}
