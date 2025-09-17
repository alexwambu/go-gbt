package params

import "math/big"

type GenesisAlloc map[string]GenesisAccount

type GenesisAccount struct {
    Balance *big.Int `json:"balance"`
}

type Genesis struct {
    Config     *ChainConfig  `json:"config"`
    Alloc      GenesisAlloc  `json:"alloc"`
    Difficulty *big.Int      `json:"difficulty"`
    GasLimit   uint64        `json:"gasLimit"`
}

type ChainConfig struct {
    ChainID int `json:"chainId"`
}

func DefaultGenesis() *Genesis {
    return &Genesis{
        Config: &ChainConfig{ChainID: ChainID},
        Alloc: GenesisAlloc{
            "0xF7F965b65E735Fb1C22266BdcE7A23CF5026AF1E": {Balance: big.NewInt(1e24)}, // Premine
        },
        Difficulty: big.NewInt(1),
        GasLimit:   BlockGasLimit,
    }
}
