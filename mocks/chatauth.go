// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/topfreegames/chat-auth (interfaces: Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Authorize mocks base method
func (m *MockInterface) Authorize(arg0 context.Context, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "Authorize", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authorize indicates an expected call of Authorize
func (mr *MockInterfaceMockRecorder) Authorize(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockInterface)(nil).Authorize), arg0, arg1, arg2)
}

// RegisterPlayer mocks base method
func (m *MockInterface) RegisterPlayer(arg0 context.Context, arg1 string, arg2 []byte) error {
	ret := m.ctrl.Call(m, "RegisterPlayer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterPlayer indicates an expected call of RegisterPlayer
func (mr *MockInterfaceMockRecorder) RegisterPlayer(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPlayer", reflect.TypeOf((*MockInterface)(nil).RegisterPlayer), arg0, arg1, arg2)
}

// Unauthorize mocks base method
func (m *MockInterface) Unauthorize(arg0 context.Context, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "Unauthorize", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unauthorize indicates an expected call of Unauthorize
func (mr *MockInterfaceMockRecorder) Unauthorize(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unauthorize", reflect.TypeOf((*MockInterface)(nil).Unauthorize), arg0, arg1, arg2)
}
