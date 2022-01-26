// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/shape_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIShapeInterface is a mock of IShapeInterface interface.
type MockIShapeInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIShapeInterfaceMockRecorder
}

// MockIShapeInterfaceMockRecorder is the mock recorder for MockIShapeInterface.
type MockIShapeInterfaceMockRecorder struct {
	mock *MockIShapeInterface
}

// NewMockIShapeInterface creates a new mock instance.
func NewMockIShapeInterface(ctrl *gomock.Controller) *MockIShapeInterface {
	mock := &MockIShapeInterface{ctrl: ctrl}
	mock.recorder = &MockIShapeInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIShapeInterface) EXPECT() *MockIShapeInterfaceMockRecorder {
	return m.recorder
}

// Area mocks base method.
func (m *MockIShapeInterface) Area() float32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Area")
	ret0, _ := ret[0].(float32)
	return ret0
}

// Area indicates an expected call of Area.
func (mr *MockIShapeInterfaceMockRecorder) Area() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Area", reflect.TypeOf((*MockIShapeInterface)(nil).Area))
}
