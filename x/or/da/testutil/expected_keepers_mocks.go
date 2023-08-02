// Code generated by MockGen. DO NOT EDIT.
// Source: x/or/da/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/Finschia/finschia-sdk/types"
	types0 "github.com/Finschia/finschia-sdk/x/auth/types"
	types1 "github.com/Finschia/finschia-sdk/x/or/da/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetParams mocks base method.
func (m *MockAccountKeeper) GetParams(ctx types.Context) types0.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types0.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockAccountKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockAccountKeeper)(nil).GetParams), ctx)
}

// MockRollupKeeper is a mock of RollupKeeper interface.
type MockRollupKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockRollupKeeperMockRecorder
}

// MockRollupKeeperMockRecorder is the mock recorder for MockRollupKeeper.
type MockRollupKeeperMockRecorder struct {
	mock *MockRollupKeeper
}

// NewMockRollupKeeper creates a new mock instance.
func NewMockRollupKeeper(ctrl *gomock.Controller) *MockRollupKeeper {
	mock := &MockRollupKeeper{ctrl: ctrl}
	mock.recorder = &MockRollupKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRollupKeeper) EXPECT() *MockRollupKeeperMockRecorder {
	return m.recorder
}

// GetRegisteredRollups mocks base method.
func (m *MockRollupKeeper) GetRegisteredRollups(ctx types.Context) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegisteredRollups", ctx)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetRegisteredRollups indicates an expected call of GetRegisteredRollups.
func (mr *MockRollupKeeperMockRecorder) GetRegisteredRollups(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegisteredRollups", reflect.TypeOf((*MockRollupKeeper)(nil).GetRegisteredRollups), ctx)
}

// GetRollup mocks base method.
func (m *MockRollupKeeper) GetRollup(ctx types.Context, rollupID string) (types1.RollupInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRollup", ctx, rollupID)
	ret0, _ := ret[0].(types1.RollupInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRollup indicates an expected call of GetRollup.
func (mr *MockRollupKeeperMockRecorder) GetRollup(ctx, rollupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRollup", reflect.TypeOf((*MockRollupKeeper)(nil).GetRollup), ctx, rollupID)
}
