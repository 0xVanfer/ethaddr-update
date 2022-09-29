package platypus

import (
	"ethaddr-update/eth"
	"ethaddr-update/utils/common"
	"fmt"

	"github.com/0xVanfer/abigen/platypus/platypusRouterSavax"
	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/types"
)

func PlatypusAvalancheMainPools() {
	network := chainId.AvalancheChainName
	connector := eth.GetConnector(network)
	mainPoolRouterAddress := ethaddr.PlatypusRouterList[network][ethaddr.PlatypusMainPoolsName]
	router, _ := platypusRouterSavax.NewPlatypusRouterSavax(types.ToAddress(mainPoolRouterAddress), connector)
	addresses, _ := router.GetTokenAddresses(nil)
	for _, address := range addresses {
		fmt.Println(ethaddr.PlatypusLpList[network][ethaddr.PlatypusMainPoolsName][types.ToLowerString(address)])
		found := ethaddr.PlatypusLpList[network][ethaddr.PlatypusMainPoolsName][types.ToLowerString(address)] != ""
		common.ProcessBool(!found, "Platypus lack main pool:", types.ToLowerString(address))
	}
}
