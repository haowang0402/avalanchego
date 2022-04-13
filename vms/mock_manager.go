// Code generated by MockGen. DO NOT EDIT.
// Source: vms/manager.go

// Package vms is a generated GoMock package.
package vms

import (
	reflect "reflect"

	ids "github.com/haowang0402/avalanchego/ids"
	snow "github.com/haowang0402/avalanchego/snow"
	gomock "github.com/golang/mock/gomock"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockFactory) New(arg0 *snow.Context) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockFactoryMockRecorder) New(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockFactory)(nil).New), arg0)
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Alias mocks base method.
func (m *MockManager) Alias(id ids.ID, alias string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alias", id, alias)
	ret0, _ := ret[0].(error)
	return ret0
}

// Alias indicates an expected call of Alias.
func (mr *MockManagerMockRecorder) Alias(id, alias interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alias", reflect.TypeOf((*MockManager)(nil).Alias), id, alias)
}

// Aliases mocks base method.
func (m *MockManager) Aliases(id ids.ID) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Aliases", id)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Aliases indicates an expected call of Aliases.
func (mr *MockManagerMockRecorder) Aliases(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aliases", reflect.TypeOf((*MockManager)(nil).Aliases), id)
}

// GetFactory mocks base method.
func (m *MockManager) GetFactory(vmID ids.ID) (Factory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFactory", vmID)
	ret0, _ := ret[0].(Factory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFactory indicates an expected call of GetFactory.
func (mr *MockManagerMockRecorder) GetFactory(vmID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFactory", reflect.TypeOf((*MockManager)(nil).GetFactory), vmID)
}

// ListFactories mocks base method.
func (m *MockManager) ListFactories() ([]ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFactories")
	ret0, _ := ret[0].([]ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFactories indicates an expected call of ListFactories.
func (mr *MockManagerMockRecorder) ListFactories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFactories", reflect.TypeOf((*MockManager)(nil).ListFactories))
}

// Lookup mocks base method.
func (m *MockManager) Lookup(alias string) (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", alias)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup.
func (mr *MockManagerMockRecorder) Lookup(alias interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockManager)(nil).Lookup), alias)
}

// PrimaryAlias mocks base method.
func (m *MockManager) PrimaryAlias(id ids.ID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrimaryAlias", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrimaryAlias indicates an expected call of PrimaryAlias.
func (mr *MockManagerMockRecorder) PrimaryAlias(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrimaryAlias", reflect.TypeOf((*MockManager)(nil).PrimaryAlias), id)
}

// PrimaryAliasOrDefault mocks base method.
func (m *MockManager) PrimaryAliasOrDefault(id ids.ID) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrimaryAliasOrDefault", id)
	ret0, _ := ret[0].(string)
	return ret0
}

// PrimaryAliasOrDefault indicates an expected call of PrimaryAliasOrDefault.
func (mr *MockManagerMockRecorder) PrimaryAliasOrDefault(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrimaryAliasOrDefault", reflect.TypeOf((*MockManager)(nil).PrimaryAliasOrDefault), id)
}

// RegisterFactory mocks base method.
func (m *MockManager) RegisterFactory(vmID ids.ID, factory Factory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterFactory", vmID, factory)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterFactory indicates an expected call of RegisterFactory.
func (mr *MockManagerMockRecorder) RegisterFactory(vmID, factory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterFactory", reflect.TypeOf((*MockManager)(nil).RegisterFactory), vmID, factory)
}

// RemoveAliases mocks base method.
func (m *MockManager) RemoveAliases(id ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveAliases", id)
}

// RemoveAliases indicates an expected call of RemoveAliases.
func (mr *MockManagerMockRecorder) RemoveAliases(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAliases", reflect.TypeOf((*MockManager)(nil).RemoveAliases), id)
}

// Versions mocks base method.
func (m *MockManager) Versions() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Versions")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Versions indicates an expected call of Versions.
func (mr *MockManagerMockRecorder) Versions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Versions", reflect.TypeOf((*MockManager)(nil).Versions))
}
