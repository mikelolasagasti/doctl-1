// Code generated by MockGen. DO NOT EDIT.
// Source: regions.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "github.com/digitalocean/doctl/do"
	gomock "go.uber.org/mock/gomock"
)

// MockRegionsService is a mock of RegionsService interface.
type MockRegionsService struct {
	ctrl     *gomock.Controller
	recorder *MockRegionsServiceMockRecorder
}

// MockRegionsServiceMockRecorder is the mock recorder for MockRegionsService.
type MockRegionsServiceMockRecorder struct {
	mock *MockRegionsService
}

// NewMockRegionsService creates a new mock instance.
func NewMockRegionsService(ctrl *gomock.Controller) *MockRegionsService {
	mock := &MockRegionsService{ctrl: ctrl}
	mock.recorder = &MockRegionsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegionsService) EXPECT() *MockRegionsServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockRegionsService) List() (do.Regions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(do.Regions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRegionsServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRegionsService)(nil).List))
}
