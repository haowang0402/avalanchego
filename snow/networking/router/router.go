// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package router

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/haowang0402/avalanchego/api/health"
	"github.com/haowang0402/avalanchego/ids"
	"github.com/haowang0402/avalanchego/message"
	"github.com/haowang0402/avalanchego/snow/networking/benchlist"
	"github.com/haowang0402/avalanchego/snow/networking/handler"
	"github.com/haowang0402/avalanchego/snow/networking/timeout"
	"github.com/haowang0402/avalanchego/utils/logging"
)

// Router routes consensus messages to the Handler of the consensus
// engine that the messages are intended for
type Router interface {
	ExternalHandler
	InternalHandler

	Initialize(
		nodeID ids.ShortID,
		log logging.Logger,
		msgCreator message.Creator,
		timeouts *timeout.Manager,
		shutdownTimeout time.Duration,
		criticalChains ids.Set,
		onFatal func(exitCode int),
		healthConfig HealthConfig,
		metricsNamespace string,
		metricsRegisterer prometheus.Registerer,
	) error
	Shutdown()
	AddChain(chain handler.Handler)
	health.Checker
}

// InternalHandler deals with messages internal to this node
type InternalHandler interface {
	benchlist.Benchable

	RegisterRequest(
		nodeID ids.ShortID,
		chainID ids.ID,
		requestID uint32,
		op message.Op,
	)
}
