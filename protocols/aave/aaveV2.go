package aave

import (
	"errors"
	"ethaddr-update/utils/common"
	"fmt"
	"strings"

	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/types"
	"github.com/imroc/req"
)

// Aave v2 avalanche AVSTokens check.
func AaveV2AVS(network string) {
	pools, _ := ReqAavePools(network)
	for _, pool := range pools {
		if pool.Symbol[:3] == "Amm" {
			continue
		}
		// atokens
		common.FindAndAlarm(pool.AToken.ID, ethaddr.AaveATokenV2List[network], "Aave", network, pool.Symbol, "AToken:", types.ToLowerString(pool.AToken.ID), "not found.")
		// vtokens
		common.FindAndAlarm(pool.VToken.ID, ethaddr.AaveVTokenV2List[network], "Aave", network, pool.Symbol, "VToken:", types.ToLowerString(pool.VToken.ID), "not found.")
		// stokens
		common.FindAndAlarm(pool.SToken.ID, ethaddr.AaveSTokenV2List[network], "Aave", network, pool.Symbol, "SToken:", types.ToLowerString(pool.SToken.ID), "not found.")
	}
	fmt.Println("aave v2", network, "checked.")
}

type aaveAllReservesReq struct {
	Data struct {
		Reserves []aaveAllReserves `json:"reserves"`
	} `json:"data"`
}
type aaveAToken struct {
	ID string `json:"id"`
}

type aaveAllReserves struct {
	Symbol                     string     `json:"symbol"`
	Decimals                   int        `json:"decimals"`
	UnderlyingAsset            string     `json:"underlyingAsset"`
	AToken                     aaveAToken `json:"aToken"`
	VToken                     aaveAToken `json:"vToken"`
	SToken                     aaveAToken `json:"sToken"`
	LiquidityRate              string     `json:"liquidityRate"`
	StableBorrowRate           string     `json:"stableBorrowRate"`
	VariableBorrowRate         string     `json:"variableBorrowRate"`
	AEmissionPerSecond         string     `json:"aEmissionPerSecond"`
	VEmissionPerSecond         string     `json:"vEmissionPerSecond"`
	SEmissionPerSecond         string     `json:"sEmissionPerSecond"`
	ReserveFactor              string     `json:"reserveFactor"`
	TotalDeposits              string     `json:"totalDeposits"`
	TotalLiquidity             string     `json:"totalLiquidity"`
	TotalScaledVariableDebt    string     `json:"totalScaledVariableDebt"`
	TotalPrincipalStableDebt   string     `json:"totalPrincipalStableDebt"`
	TotalCurrentVariableDebt   string     `json:"totalCurrentVariableDebt"`
	TotalLiquidityAsCollateral string     `json:"totalLiquidityAsCollateral"`
	UtilizationRate            string     `json:"utilizationRate"`
}

func ReqAavePools(network string) (reserves []aaveAllReserves, err error) {
	var url string
	switch network {
	case chainId.AvalancheChainName:
		url = "https://api.thegraph.com/subgraphs/name/aave/protocol-v2-avalanche"
	case chainId.EthereumChainName:
		url = "https://api.thegraph.com/subgraphs/name/aave/protocol-v2"
	default:
		errString := "req aave pools not supported network:" + network
		err = errors.New(errString)
		return
	}
	payload := strings.NewReader(`{"query":"{\n reserves {\n symbol\n decimals\n underlyingAsset\n aToken{id}\n vToken{id}\n sToken{id}\n liquidityRate\n stableBorrowRate\n variableBorrowRate\n aEmissionPerSecond\n vEmissionPerSecond\n sEmissionPerSecond\n totalDeposits\n totalLiquidity\n totalScaledVariableDebt\n totalPrincipalStableDebt\n reserveFactor\n totalCurrentVariableDebt\n totalLiquidityAsCollateral\n utilizationRate\n}\n}\n\n"}`)
	r, _ := req.Post(url, payload)
	var v aaveAllReservesReq
	err = r.ToJSON(&v)
	return v.Data.Reserves, err
}
