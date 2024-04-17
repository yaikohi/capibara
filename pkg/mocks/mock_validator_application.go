// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports (interfaces: ValidatorApplication)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	ports "github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports"
	gomock "go.uber.org/mock/gomock"
)

// MockValidatorApplication is a mock of ValidatorApplication interface.
type MockValidatorApplication struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorApplicationMockRecorder
}

// MockValidatorApplicationMockRecorder is the mock recorder for MockValidatorApplication.
type MockValidatorApplicationMockRecorder struct {
	mock *MockValidatorApplication
}

// NewMockValidatorApplication creates a new mock instance.
func NewMockValidatorApplication(ctrl *gomock.Controller) *MockValidatorApplication {
	mock := &MockValidatorApplication{ctrl: ctrl}
	mock.recorder = &MockValidatorApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorApplication) EXPECT() *MockValidatorApplicationMockRecorder {
	return m.recorder
}

// Struct mocks base method.
func (m *MockValidatorApplication) Struct(arg0 ports.Evaluable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Struct", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Struct indicates an expected call of Struct.
func (mr *MockValidatorApplicationMockRecorder) Struct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Struct", reflect.TypeOf((*MockValidatorApplication)(nil).Struct), arg0)
}
