package misc

import "github.com/ethereum/go-ethereum/params"

func ZeroGasPriceChain(config *params.ChainConfig) bool {
	return config.Clique != nil && config.Clique.ZeroGasPrice
}