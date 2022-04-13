// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/haowang0402/avalanchego/snow"
	"github.com/haowang0402/avalanchego/snow/consensus/snowball"
	"github.com/haowang0402/avalanchego/snow/consensus/snowman"
	"github.com/haowang0402/avalanchego/snow/engine/common"
	"github.com/haowang0402/avalanchego/snow/engine/snowman/block"
	"github.com/haowang0402/avalanchego/snow/validators"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	common.AllGetsServer

	Ctx        *snow.ConsensusContext
	VM         block.ChainVM
	Sender     common.Sender
	Validators validators.Set
	Params     snowball.Parameters
	Consensus  snowman.Consensus
}
