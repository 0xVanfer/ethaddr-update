package aave

import (
	"ethaddr-update/eth"
	"ethaddr-update/utils/common"

	"github.com/0xVanfer/abigen/aave/aavePoolDataProvider"
	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/types"
)

// Aave v3 avalanche AVSTokens check.
func AaveV3AvalancheAVS(network string) {
	// eth connector
	connector := eth.GetConnector(network)
	// data provider
	dataProviderAddress := ethaddr.AavePoolDataProviderList[network]
	provider, _ := aavePoolDataProvider.NewAavePoolDataProvider(types.ToAddress(dataProviderAddress), connector)
	// aave v3 underlyings
	underlyings, _ := provider.GetAllReservesTokens(nil)
	for _, underlying := range underlyings {
		addrs, _ := provider.GetReserveTokensAddresses(nil, underlying.TokenAddress)
		// atokens
		common.FindAndAlarm(addrs.ATokenAddress, ethaddr.AaveATokenV3List[network], "Aave", network, underlying.Symbol, "AToken:", types.ToLowerString(addrs.ATokenAddress), "not found.")
		// vtokens
		common.FindAndAlarm(addrs.VariableDebtTokenAddress, ethaddr.AaveVTokenV3List[network], "Aave", network, underlying.Symbol, "VToken:", types.ToLowerString(addrs.VariableDebtTokenAddress), "not found.")
		// stokens
		common.FindAndAlarm(addrs.StableDebtTokenAddress, ethaddr.AaveSTokenV3List[network], "Aave", network, underlying.Symbol, "SToken:", types.ToLowerString(addrs.StableDebtTokenAddress), "not found.")
	}
}
