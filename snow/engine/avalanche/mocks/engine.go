// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	consensusavalanche "github.com/haowang0402/avalanchego/snow/consensus/avalanche"
	common "github.com/haowang0402/avalanchego/snow/engine/common"

	ids "github.com/haowang0402/avalanchego/ids"

	mock "github.com/stretchr/testify/mock"

	snow "github.com/haowang0402/avalanchego/snow"

	time "time"

	version "github.com/haowang0402/avalanchego/version"
)

// Engine is an autogenerated mock type for the Engine type
type Engine struct {
	mock.Mock
}

// Accepted provides a mock function with given fields: validatorID, requestID, containerIDs
func (_m *Engine) Accepted(validatorID ids.ShortID, requestID uint32, containerIDs []ids.ID) error {
	ret := _m.Called(validatorID, requestID, containerIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, []ids.ID) error); ok {
		r0 = rf(validatorID, requestID, containerIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AcceptedFrontier provides a mock function with given fields: validatorID, requestID, containerIDs
func (_m *Engine) AcceptedFrontier(validatorID ids.ShortID, requestID uint32, containerIDs []ids.ID) error {
	ret := _m.Called(validatorID, requestID, containerIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, []ids.ID) error); ok {
		r0 = rf(validatorID, requestID, containerIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ancestors provides a mock function with given fields: validatorID, requestID, containers
func (_m *Engine) Ancestors(validatorID ids.ShortID, requestID uint32, containers [][]byte) error {
	ret := _m.Called(validatorID, requestID, containers)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, [][]byte) error); ok {
		r0 = rf(validatorID, requestID, containers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppGossip provides a mock function with given fields: nodeID, msg
func (_m *Engine) AppGossip(nodeID ids.ShortID, msg []byte) error {
	ret := _m.Called(nodeID, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, []byte) error); ok {
		r0 = rf(nodeID, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppRequest provides a mock function with given fields: nodeID, requestID, deadline, request
func (_m *Engine) AppRequest(nodeID ids.ShortID, requestID uint32, deadline time.Time, request []byte) error {
	ret := _m.Called(nodeID, requestID, deadline, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, time.Time, []byte) error); ok {
		r0 = rf(nodeID, requestID, deadline, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppRequestFailed provides a mock function with given fields: nodeID, requestID
func (_m *Engine) AppRequestFailed(nodeID ids.ShortID, requestID uint32) error {
	ret := _m.Called(nodeID, requestID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32) error); ok {
		r0 = rf(nodeID, requestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppResponse provides a mock function with given fields: nodeID, requestID, response
func (_m *Engine) AppResponse(nodeID ids.ShortID, requestID uint32, response []byte) error {
	ret := _m.Called(nodeID, requestID, response)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, []byte) error); ok {
		r0 = rf(nodeID, requestID, response)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Chits provides a mock function with given fields: validatorID, requestID, containerIDs
func (_m *Engine) Chits(validatorID ids.ShortID, requestID uint32, containerIDs []ids.ID) error {
	ret := _m.Called(validatorID, requestID, containerIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, []ids.ID) error); ok {
		r0 = rf(validatorID, requestID, containerIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connected provides a mock function with given fields: id, nodeVersion
func (_m *Engine) Connected(id ids.ShortID, nodeVersion version.Application) error {
	ret := _m.Called(id, nodeVersion)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, version.Application) error); ok {
		r0 = rf(id, nodeVersion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Context provides a mock function with given fields:
func (_m *Engine) Context() *snow.ConsensusContext {
	ret := _m.Called()

	var r0 *snow.ConsensusContext
	if rf, ok := ret.Get(0).(func() *snow.ConsensusContext); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snow.ConsensusContext)
		}
	}

	return r0
}

// Disconnected provides a mock function with given fields: id
func (_m *Engine) Disconnected(id ids.ShortID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: validatorID, requestID, containerID
func (_m *Engine) Get(validatorID ids.ShortID, requestID uint32, containerID ids.ID) error {
	ret := _m.Called(validatorID, requestID, containerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, ids.ID) error); ok {
		r0 = rf(validatorID, requestID, containerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAccepted provides a mock function with given fields: validatorID, requestID, containerIDs
func (_m *Engine) GetAccepted(validatorID ids.ShortID, requestID uint32, containerIDs []ids.ID) error {
	ret := _m.Called(validatorID, requestID, containerIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, []ids.ID) error); ok {
		r0 = rf(validatorID, requestID, containerIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAcceptedFailed provides a mock function with given fields: validatorID, requestID
func (_m *Engine) GetAcceptedFailed(validatorID ids.ShortID, requestID uint32) error {
	ret := _m.Called(validatorID, requestID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32) error); ok {
		r0 = rf(validatorID, requestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAcceptedFrontier provides a mock function with given fields: validatorID, requestID
func (_m *Engine) GetAcceptedFrontier(validatorID ids.ShortID, requestID uint32) error {
	ret := _m.Called(validatorID, requestID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32) error); ok {
		r0 = rf(validatorID, requestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAcceptedFrontierFailed provides a mock function with given fields: validatorID, requestID
func (_m *Engine) GetAcceptedFrontierFailed(validatorID ids.ShortID, requestID uint32) error {
	ret := _m.Called(validatorID, requestID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32) error); ok {
		r0 = rf(validatorID, requestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAncestors provides a mock function with given fields: validatorID, requestID, containerID
func (_m *Engine) GetAncestors(validatorID ids.ShortID, requestID uint32, containerID ids.ID) error {
	ret := _m.Called(validatorID, requestID, containerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, ids.ID) error); ok {
		r0 = rf(validatorID, requestID, containerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAncestorsFailed provides a mock function with given fields: validatorID, requestID
func (_m *Engine) GetAncestorsFailed(validatorID ids.ShortID, requestID uint32) error {
	ret := _m.Called(validatorID, requestID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32) error); ok {
		r0 = rf(validatorID, requestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFailed provides a mock function with given fields: validatorID, requestID
func (_m *Engine) GetFailed(validatorID ids.ShortID, requestID uint32) error {
	ret := _m.Called(validatorID, requestID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32) error); ok {
		r0 = rf(validatorID, requestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetVM provides a mock function with given fields:
func (_m *Engine) GetVM() common.VM {
	ret := _m.Called()

	var r0 common.VM
	if rf, ok := ret.Get(0).(func() common.VM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.VM)
		}
	}

	return r0
}

// GetVtx provides a mock function with given fields: vtxID
func (_m *Engine) GetVtx(vtxID ids.ID) (consensusavalanche.Vertex, error) {
	ret := _m.Called(vtxID)

	var r0 consensusavalanche.Vertex
	if rf, ok := ret.Get(0).(func(ids.ID) consensusavalanche.Vertex); ok {
		r0 = rf(vtxID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(consensusavalanche.Vertex)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ids.ID) error); ok {
		r1 = rf(vtxID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Gossip provides a mock function with given fields:
func (_m *Engine) Gossip() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Halt provides a mock function with given fields:
func (_m *Engine) Halt() {
	_m.Called()
}

// HealthCheck provides a mock function with given fields:
func (_m *Engine) HealthCheck() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Notify provides a mock function with given fields: _a0
func (_m *Engine) Notify(_a0 common.Message) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Message) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PullQuery provides a mock function with given fields: validatorID, requestID, containerID
func (_m *Engine) PullQuery(validatorID ids.ShortID, requestID uint32, containerID ids.ID) error {
	ret := _m.Called(validatorID, requestID, containerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, ids.ID) error); ok {
		r0 = rf(validatorID, requestID, containerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PushQuery provides a mock function with given fields: validatorID, requestID, container
func (_m *Engine) PushQuery(validatorID ids.ShortID, requestID uint32, container []byte) error {
	ret := _m.Called(validatorID, requestID, container)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, []byte) error); ok {
		r0 = rf(validatorID, requestID, container)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Put provides a mock function with given fields: validatorID, requestID, container
func (_m *Engine) Put(validatorID ids.ShortID, requestID uint32, container []byte) error {
	ret := _m.Called(validatorID, requestID, container)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, []byte) error); ok {
		r0 = rf(validatorID, requestID, container)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryFailed provides a mock function with given fields: validatorID, requestID
func (_m *Engine) QueryFailed(validatorID ids.ShortID, requestID uint32) error {
	ret := _m.Called(validatorID, requestID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32) error); ok {
		r0 = rf(validatorID, requestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields:
func (_m *Engine) Shutdown() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: startReqID
func (_m *Engine) Start(startReqID uint32) error {
	ret := _m.Called(startReqID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint32) error); ok {
		r0 = rf(startReqID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Timeout provides a mock function with given fields:
func (_m *Engine) Timeout() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
