// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/haowang0402/avalanchego/database/memdb"
	"github.com/haowang0402/avalanchego/snow/consensus/avalanche"
	"github.com/haowang0402/avalanchego/snow/consensus/snowball"
	"github.com/haowang0402/avalanchego/snow/engine/avalanche/bootstrap"
	"github.com/haowang0402/avalanchego/snow/engine/avalanche/vertex"
	"github.com/haowang0402/avalanchego/snow/engine/common"
	"github.com/haowang0402/avalanchego/snow/engine/common/queue"
)

func DefaultConfig() (common.Config, bootstrap.Config, Config) {
	vtxBlocked, _ := queue.NewWithMissing(memdb.New(), "", prometheus.NewRegistry())
	txBlocked, _ := queue.New(memdb.New(), "", prometheus.NewRegistry())

	commonCfg := common.DefaultConfigTest()

	bootstrapConfig := bootstrap.Config{
		Config:     commonCfg,
		VtxBlocked: vtxBlocked,
		TxBlocked:  txBlocked,
		Manager:    &vertex.TestManager{},
		VM:         &vertex.TestVM{},
	}

	engineConfig := Config{
		Ctx:        bootstrapConfig.Ctx,
		VM:         bootstrapConfig.VM,
		Manager:    bootstrapConfig.Manager,
		Sender:     bootstrapConfig.Sender,
		Validators: bootstrapConfig.Validators,
		Params: avalanche.Parameters{
			Parameters: snowball.Parameters{
				K:                     1,
				Alpha:                 1,
				BetaVirtuous:          1,
				BetaRogue:             2,
				ConcurrentRepolls:     1,
				OptimalProcessing:     100,
				MaxOutstandingItems:   1,
				MaxItemProcessingTime: 1,
			},
			Parents:   2,
			BatchSize: 1,
		},
		Consensus: &avalanche.Topological{},
	}

	return commonCfg, bootstrapConfig, engineConfig
}
