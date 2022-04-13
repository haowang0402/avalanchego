// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"math"

	"github.com/haowang0402/avalanchego/codec"
	"github.com/haowang0402/avalanchego/codec/linearcodec"
	"github.com/haowang0402/avalanchego/codec/reflectcodec"
)

const version = 0

var c codec.Manager

func init() {
	lc := linearcodec.New(reflectcodec.DefaultTagName, math.MaxUint32)
	c = codec.NewManager(math.MaxInt32)

	err := c.RegisterCodec(version, lc)
	if err != nil {
		panic(err)
	}
}
