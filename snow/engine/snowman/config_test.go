// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/haowang0402/avalanchego/database/memdb"
	"github.com/haowang0402/avalanchego/snow/consensus/snowball"
	"github.com/haowang0402/avalanchego/snow/consensus/snowman"
	"github.com/haowang0402/avalanchego/snow/engine/common"
	"github.com/haowang0402/avalanchego/snow/engine/common/queue"
	"github.com/haowang0402/avalanchego/snow/engine/snowman/block"
	"github.com/haowang0402/avalanchego/snow/engine/snowman/bootstrap"
)

func DefaultConfigs() (bootstrap.Config, Config) {
	blocked, _ := queue.NewWithMissing(memdb.New(), "", prometheus.NewRegistry())

	bootstrapConfig := bootstrap.Config{
		Config:  common.DefaultConfigTest(),
		Blocked: blocked,
		VM:      &block.TestVM{},
	}

	engineConfig := Config{
		Ctx:        bootstrapConfig.Ctx,
		VM:         bootstrapConfig.VM,
		Sender:     bootstrapConfig.Sender,
		Validators: bootstrapConfig.Validators,
		Params: snowball.Parameters{
			K:                     1,
			Alpha:                 1,
			BetaVirtuous:          1,
			BetaRogue:             2,
			ConcurrentRepolls:     1,
			OptimalProcessing:     100,
			MaxOutstandingItems:   1,
			MaxItemProcessingTime: 1,
		},
		Consensus: &snowman.Topological{},
	}

	return bootstrapConfig, engineConfig
}
