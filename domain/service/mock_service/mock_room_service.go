// Code generated by MockGen. DO NOT EDIT.
// Source: room_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIRoomService is a mock of IRoomService interface
type MockIRoomService struct {
	ctrl     *gomock.Controller
	recorder *MockIRoomServiceMockRecorder
}

// MockIRoomServiceMockRecorder is the mock recorder for MockIRoomService
type MockIRoomServiceMockRecorder struct {
	mock *MockIRoomService
}

// NewMockIRoomService creates a new mock instance
func NewMockIRoomService(ctrl *gomock.Controller) *MockIRoomService {
	mock := &MockIRoomService{ctrl: ctrl}
	mock.recorder = &MockIRoomServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRoomService) EXPECT() *MockIRoomServiceMockRecorder {
	return m.recorder
}

// CanInsert mocks base method
func (m *MockIRoomService) CanInsert(id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanInsert", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanInsert indicates an expected call of CanInsert
func (mr *MockIRoomServiceMockRecorder) CanInsert(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanInsert", reflect.TypeOf((*MockIRoomService)(nil).CanInsert), id)
}

// IsLock mocks base method
func (m *MockIRoomService) IsLock(id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLock", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLock indicates an expected call of IsLock
func (mr *MockIRoomServiceMockRecorder) IsLock(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLock", reflect.TypeOf((*MockIRoomService)(nil).IsLock), id)
}

// IsOwner mocks base method
func (m *MockIRoomService) IsOwner(uid, rid string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOwner", uid, rid)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOwner indicates an expected call of IsOwner
func (mr *MockIRoomServiceMockRecorder) IsOwner(uid, rid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOwner", reflect.TypeOf((*MockIRoomService)(nil).IsOwner), uid, rid)
}
