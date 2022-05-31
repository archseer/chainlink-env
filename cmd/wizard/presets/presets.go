package presets

import (
	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-env/pkg/cdk8s/blockscout"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/chainlink"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/geth"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/mockserver"
	mockservercfg "github.com/smartcontractkit/chainlink-env/pkg/helm/mockserver-cfg"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/reorg"
)

// EVMOneNode local development Chainlink deployment
func EVMOneNode(config *environment.Config) error {
	return environment.New(config).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(geth.New(nil)).
		AddHelm(chainlink.New(nil)).
		Run()
}

// EVMMinimalLocalBS local development Chainlink deployment,
// 1 bootstrap + 4 oracles (minimal requirements for OCR) + Blockscout
func EVMMinimalLocalBS(config *environment.Config) error {
	return environment.New(config).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddChart(blockscout.New(&blockscout.Props{})).
		AddHelm(geth.New(nil)).
		AddHelm(chainlink.New(map[string]interface{}{
			"replicas": 5,
		})).
		Run()
}

// EVMMinimalLocal local development Chainlink deployment,
// 1 bootstrap + 4 oracles (minimal requirements for OCR)
func EVMMinimalLocal(config *environment.Config) error {
	return environment.New(config).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(geth.New(nil)).
		AddHelm(chainlink.New(map[string]interface{}{
			"replicas": 5,
		})).
		Run()
}

// EVMExternal local development Chainlink deployment for an external (testnet) network
func EVMExternal(config *environment.Config) error {
	return environment.New(config).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(chainlink.New(nil)).
		Run()
}

// EVMReorg deployment for two Ethereum networks re-org test
func EVMReorg(config *environment.Config) error {
	return environment.New(config).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(reorg.New(
			"geth-reorg",
			map[string]interface{}{
				"geth": map[string]interface{}{
					"genesis": map[string]interface{}{
						"networkId": "1337",
					},
				},
			})).
		AddHelm(reorg.New(
			"geth-reorg-2",
			map[string]interface{}{
				"geth": map[string]interface{}{
					"genesis": map[string]interface{}{
						"networkId": "2337",
					},
				},
			})).
		AddHelm(chainlink.New(map[string]interface{}{
			"replicas": 5,
			"env": map[string]interface{}{
				"eth_url": "ws://geth-reorg-ethereum-geth:8546",
			},
		})).
		Run()
}

// EVMSoak deployment for a long running soak tests
func EVMSoak(config *environment.Config) error {
	return environment.New(config).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(geth.New(map[string]interface{}{
			"resources": map[string]interface{}{
				"requests": map[string]interface{}{
					"cpu":    "1000m",
					"memory": "2048Mi",
				},
				"limits": map[string]interface{}{
					"cpu":    "1000m",
					"memory": "2048Mi",
				},
			},
		})).
		AddHelm(chainlink.New(map[string]interface{}{
			"replicas": 5,
			"db": map[string]interface{}{
				"stateful": true,
				"capacity": "30Gi",
				"resources": map[string]interface{}{
					"requests": map[string]interface{}{
						"cpu":    "250m",
						"memory": "256Mi",
					},
					"limits": map[string]interface{}{
						"cpu":    "250m",
						"memory": "256Mi",
					},
				},
			},
			"chainlink": map[string]interface{}{
				"resources": map[string]interface{}{
					"requests": map[string]interface{}{
						"cpu":    "1000m",
						"memory": "2048Mi",
					},
					"limits": map[string]interface{}{
						"cpu":    "1000m",
						"memory": "2048Mi",
					},
				},
			},
		})).
		Run()
}
