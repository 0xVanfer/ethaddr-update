package main

import (
	"ethaddr-update/config"
	"ethaddr-update/eth"
	"ethaddr-update/protocols/aave"
	"ethaddr-update/protocols/compoundlike"
	"ethaddr-update/protocols/platypus"
	"ethaddr-update/utils/telegram"

	"github.com/0xVanfer/chainId"
)

func init() {
	// init config
	config.Init()
	// cnf := config.GetConfig()

	// initiate eth connection
	usedChains := []string{
		chainId.AvalancheChainName,
		chainId.EthereumChainName,
		chainId.PolygonChainName}
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
	aave.AaveV2AVS(chainId.AvalancheChainName)
	aave.AaveV2AVS(chainId.EthereumChainName)
	aave.AaveV3AVS(chainId.AvalancheChainName)
	aave.AaveV3AVS(chainId.PolygonChainName)
	compoundlike.BenqiCTokens()
	compoundlike.TraderjoeCTokens()
	platypus.PlatypusAvalancheLps()
}
