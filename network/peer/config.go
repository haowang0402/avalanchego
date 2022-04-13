// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"time"

	"github.com/haowang0402/avalanchego/ids"
	"github.com/haowang0402/avalanchego/message"
	"github.com/haowang0402/avalanchego/network/throttling"
	"github.com/haowang0402/avalanchego/snow/networking/router"
	"github.com/haowang0402/avalanchego/snow/validators"
	"github.com/haowang0402/avalanchego/utils/logging"
	"github.com/haowang0402/avalanchego/utils/timer/mockable"
	"github.com/haowang0402/avalanchego/version"
)

type Config struct {
	// Size, in bytes, of the buffer this peer reads messages into
	ReadBufferSize int
	// Size, in bytes, of the buffer this peer writes messages into
	WriteBufferSize      int
	Clock                mockable.Clock
	Metrics              *Metrics
	MessageCreator       message.Creator
	Log                  logging.Logger
	InboundMsgThrottler  throttling.InboundMsgThrottler
	OutboundMsgThrottler throttling.OutboundMsgThrottler
	Network              Network
	Router               router.InboundHandler
	VersionCompatibility version.Compatibility
	VersionParser        version.ApplicationParser
	MySubnets            ids.Set
	Beacons              validators.Set
	NetworkID            uint32
	PingFrequency        time.Duration
	PongTimeout          time.Duration
	MaxClockDifference   time.Duration

	// Unix time of the last message sent and received respectively
	// Must only be accessed atomically
	LastSent, LastReceived int64
}
