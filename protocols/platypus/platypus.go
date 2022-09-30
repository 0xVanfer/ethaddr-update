package platypus

import (
	"ethaddr-update/eth"
	"ethaddr-update/utils/common"
	"strings"

	"github.com/0xVanfer/abigen/platypus/platypusMasterPlatypusV4"
	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/types"
)

func PlatypusAvalancheLps() {
	network := chainId.AvalancheChainName
	connector := eth.GetConnector(network)
	chef, _ := platypusMasterPlatypusV4.NewPlatypusMasterPlatypusV4(types.ToAddress(ethaddr.PlatypusMasterPlatypusV4List[chainId.AvalancheChainName]), connector)
	new, _ := chef.NewMasterPlatypus(nil)
	hasNewMaster := types.ToLowerString(new) != ethaddr.ZEROAddress
	common.ProcessBool(hasNewMaster, "Platypus has new master :", types.ToLowerString(new))

	poolLengthBig, _ := chef.PoolLength(nil)
	length := types.ToInt(poolLengthBig)
	for pid := 0; pid < length; pid++ {
		poolInfo, _ := chef.PoolInfo(nil, types.ToBigInt(pid))
		lp := types.ToLowerString(poolInfo.LpToken)
		found := false
		for _, mapp := range ethaddr.PlatypusLpList[chainId.AvalancheChainName] {
			for _, lpaddress := range mapp {
				if strings.EqualFold(lp, lpaddress) {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		common.ProcessBool(!found, "Platypus lack lp:", lp)
	}
}
