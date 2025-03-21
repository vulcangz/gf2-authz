// Code generated by MockGen. DO NOT EDIT.
// Source: internal/logic/manager/stats.go

// Package manager is a generated GoMock package.
package manager

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
)

// MockStats is a mock of Stats interface.
type MockStats struct {
	ctrl     *gomock.Controller
	recorder *MockStatsMockRecorder
}

// MockStatsMockRecorder is the mock recorder for MockStats.
type MockStatsMockRecorder struct {
	mock *MockStats
}

// NewMockStats creates a new mock instance.
func NewMockStats(ctrl *gomock.Controller) *MockStats {
	mock := &MockStats{ctrl: ctrl}
	mock.recorder = &MockStatsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStats) EXPECT() *MockStatsMockRecorder {
	return m.recorder
}

// BatchAddCheck mocks base method.
func (m *MockStats) BatchAddCheck(timestamp, allowed, denied int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchAddCheck", timestamp, allowed, denied)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchAddCheck indicates an expected call of BatchAddCheck.
func (mr *MockStatsMockRecorder) BatchAddCheck(timestamp, allowed, denied interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchAddCheck", reflect.TypeOf((*MockStats)(nil).BatchAddCheck), timestamp, allowed, denied)
}

// GetRepository mocks base method.
func (m *MockStats) GetRepository() orm.StatsRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepository")
	ret0, _ := ret[0].(orm.StatsRepository)
	return ret0
}

// GetRepository indicates an expected call of GetRepository.
func (mr *MockStatsMockRecorder) GetRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepository", reflect.TypeOf((*MockStats)(nil).GetRepository))
}
