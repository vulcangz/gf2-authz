// Code generated by MockGen. DO NOT EDIT.
// Source: internal/logic/manager/action.go

// Package manager is a generated GoMock package.
package manager

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

// MockAction is a mock of Action interface.
type MockAction struct {
	ctrl     *gomock.Controller
	recorder *MockActionMockRecorder
}

// MockActionMockRecorder is the mock recorder for MockAction.
type MockActionMockRecorder struct {
	mock *MockAction
}

// NewMockAction creates a new mock instance.
func NewMockAction(ctrl *gomock.Controller) *MockAction {
	mock := &MockAction{ctrl: ctrl}
	mock.recorder = &MockActionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAction) EXPECT() *MockActionMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAction) Create(identifier string) (*entity.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", identifier)
	ret0, _ := ret[0].(*entity.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockActionMockRecorder) Create(identifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAction)(nil).Create), identifier)
}

// GetRepository mocks base method.
func (m *MockAction) GetRepository() orm.ActionRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepository")
	ret0, _ := ret[0].(orm.ActionRepository)
	return ret0
}

// GetRepository indicates an expected call of GetRepository.
func (mr *MockActionMockRecorder) GetRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepository", reflect.TypeOf((*MockAction)(nil).GetRepository))
}
