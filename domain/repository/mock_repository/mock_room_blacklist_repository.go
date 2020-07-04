// Code generated by MockGen. DO NOT EDIT.
// Source: room_blacklist_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/taniwhy/mochi-match-rest/domain/models"
	reflect "reflect"
)

// MockIRoomBlacklistRepository is a mock of IRoomBlacklistRepository interface
type MockIRoomBlacklistRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRoomBlacklistRepositoryMockRecorder
}

// MockIRoomBlacklistRepositoryMockRecorder is the mock recorder for MockIRoomBlacklistRepository
type MockIRoomBlacklistRepositoryMockRecorder struct {
	mock *MockIRoomBlacklistRepository
}

// NewMockIRoomBlacklistRepository creates a new mock instance
func NewMockIRoomBlacklistRepository(ctrl *gomock.Controller) *MockIRoomBlacklistRepository {
	mock := &MockIRoomBlacklistRepository{ctrl: ctrl}
	mock.recorder = &MockIRoomBlacklistRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRoomBlacklistRepository) EXPECT() *MockIRoomBlacklistRepositoryMockRecorder {
	return m.recorder
}

// FindByRoomID mocks base method
func (m *MockIRoomBlacklistRepository) FindByRoomID(arg0 string) ([]*models.RoomBlacklist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByRoomID", arg0)
	ret0, _ := ret[0].([]*models.RoomBlacklist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByRoomID indicates an expected call of FindByRoomID
func (mr *MockIRoomBlacklistRepositoryMockRecorder) FindByRoomID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByRoomID", reflect.TypeOf((*MockIRoomBlacklistRepository)(nil).FindByRoomID), arg0)
}

// Insert mocks base method
func (m *MockIRoomBlacklistRepository) Insert(arg0 *models.RoomBlacklist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockIRoomBlacklistRepositoryMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIRoomBlacklistRepository)(nil).Insert), arg0)
}

// Delete mocks base method
func (m *MockIRoomBlacklistRepository) Delete(roomID, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", roomID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockIRoomBlacklistRepositoryMockRecorder) Delete(roomID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRoomBlacklistRepository)(nil).Delete), roomID, userID)
}
