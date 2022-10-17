package compoundlike

import (
	"ethaddr-update/eth"
	"ethaddr-update/utils/common"
	"fmt"

	"github.com/0xVanfer/abigen/benqi/benqiComptroller"
	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/types"
)

func BenqiCTokens() {
	network := chainId.AvalancheChainName
	comptrollerAddress := ethaddr.BenqiComptrollerList[network]
	comptroller, _ := benqiComptroller.NewBenqiComptroller(types.ToAddress(comptrollerAddress), eth.GetConnector(network))
	allMarkets, _ := comptroller.GetAllMarkets(nil)
	for _, ctoken := range allMarkets {
		common.FindAndAlarm(types.ToString(ctoken), ethaddr.BenqiCTokenList[network], "Benqi CToken", network, types.ToLowerString(ctoken), "not found.")
	}
	fmt.Println("benqi c token checked.")
}

func TraderjoeCTokens() {
	network := chainId.AvalancheChainName
	comptrollerAddress := ethaddr.TraderjoeJoetrollerList[network]
	comptroller, _ := benqiComptroller.NewBenqiComptroller(types.ToAddress(comptrollerAddress), eth.GetConnector(network))
	allMarkets, _ := comptroller.GetAllMarkets(nil)
	for _, ctoken := range allMarkets {
		common.FindAndAlarm(types.ToString(ctoken), ethaddr.TraderjoeCTokenList[network], "Traderjoe CToken", network, types.ToLowerString(ctoken), "not found.")
	}
	fmt.Println("traderjoe c token checked.")
}
