package main

import (
	"ethaddr-update/config"
	"ethaddr-update/eth"
	"ethaddr-update/protocols/aave"
	"ethaddr-update/protocols/compoundlike"
	"ethaddr-update/utils/telegram"

	"github.com/0xVanfer/chainId"
)

func init() {
	// init config
	config.Init()
	// cnf := config.GetConfig()

	// initiate eth connection
	usedChains := []string{chainId.AvalancheChainName, chainId.EthereumChainName}
	err := eth.InitEthConnectors(usedChains)
	if err != nil {
		panic(err)
	}

	err = telegram.InitTgBot()
	if err != nil {
		panic(err)
	}

}

func main() {
	aave.AaveV3AvalancheAVS(chainId.AvalancheChainName)
	compoundlike.BenqiCTokens()
	compoundlike.TraderjoeCTokens()
}
