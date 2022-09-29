package eth

import (
	"ethaddr-update/config"
	"fmt"

	"github.com/0xVanfer/chainId"
	"github.com/ethereum/go-ethereum/ethclient"
)

var connectorAvax *ethclient.Client
var connectorEth *ethclient.Client

func InitEthConnectors(networks []string) (err error) {
	cnf := config.GetConfig()
	for _, network := range networks {
		switch network {
		case chainId.AvalancheChainName:
			connectorAvax, err = ethclient.Dial(cnf.Eth.AvalancheNode)
		case chainId.EthereumChainName:
			connectorEth, err = ethclient.Dial(cnf.Eth.EthereumNode)
		}
		if err != nil {
			fmt.Println(network, "eth provider connection failed:", err)
			return
		}
		fmt.Println(network, "eth provider connection success.")
	}
	return
}

func GetConnector(network string) (conn *ethclient.Client) {
	var ConnectorMap = map[string]*ethclient.Client{
		chainId.AvalancheChainName: connectorAvax,
		chainId.EthereumChainName:  connectorEth,
	}
	return ConnectorMap[network]
}
