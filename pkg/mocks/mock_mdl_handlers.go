// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solrac97gr/go-jwt-auth/pkg/middleware/domain/ports (interfaces: MiddlewareHandlers)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockMiddlewareHandlers is a mock of MiddlewareHandlers interface.
type MockMiddlewareHandlers struct {
	ctrl     *gomock.Controller
	recorder *MockMiddlewareHandlersMockRecorder
}

// MockMiddlewareHandlersMockRecorder is the mock recorder for MockMiddlewareHandlers.
type MockMiddlewareHandlersMockRecorder struct {
	mock *MockMiddlewareHandlers
}

// NewMockMiddlewareHandlers creates a new mock instance.
func NewMockMiddlewareHandlers(ctrl *gomock.Controller) *MockMiddlewareHandlers {
	mock := &MockMiddlewareHandlers{ctrl: ctrl}
	mock.recorder = &MockMiddlewareHandlersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMiddlewareHandlers) EXPECT() *MockMiddlewareHandlersMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockMiddlewareHandlers) Authenticate() func(*fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate")
	ret0, _ := ret[0].(func(*fiber.Ctx) error)
	return ret0
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockMiddlewareHandlersMockRecorder) Authenticate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockMiddlewareHandlers)(nil).Authenticate))
}