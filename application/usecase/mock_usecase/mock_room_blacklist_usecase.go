// Code generated by MockGen. DO NOT EDIT.
// Source: room_blacklist_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	models "github.com/taniwhy/mochi-match-rest/domain/models"
	reflect "reflect"
)

// MockIRoomBlacklistUseCase is a mock of IRoomBlacklistUseCase interface
type MockIRoomBlacklistUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIRoomBlacklistUseCaseMockRecorder
}

// MockIRoomBlacklistUseCaseMockRecorder is the mock recorder for MockIRoomBlacklistUseCase
type MockIRoomBlacklistUseCaseMockRecorder struct {
	mock *MockIRoomBlacklistUseCase
}

// NewMockIRoomBlacklistUseCase creates a new mock instance
func NewMockIRoomBlacklistUseCase(ctrl *gomock.Controller) *MockIRoomBlacklistUseCase {
	mock := &MockIRoomBlacklistUseCase{ctrl: ctrl}
	mock.recorder = &MockIRoomBlacklistUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRoomBlacklistUseCase) EXPECT() *MockIRoomBlacklistUseCaseMockRecorder {
	return m.recorder
}

// GetByRoomID mocks base method
func (m *MockIRoomBlacklistUseCase) GetByRoomID(c *gin.Context) ([]*models.RoomBlacklist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByRoomID", c)
	ret0, _ := ret[0].([]*models.RoomBlacklist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByRoomID indicates an expected call of GetByRoomID
func (mr *MockIRoomBlacklistUseCaseMockRecorder) GetByRoomID(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByRoomID", reflect.TypeOf((*MockIRoomBlacklistUseCase)(nil).GetByRoomID), c)
}

// Insert mocks base method
func (m *MockIRoomBlacklistUseCase) Insert(c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockIRoomBlacklistUseCaseMockRecorder) Insert(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIRoomBlacklistUseCase)(nil).Insert), c)
}

// Delete mocks base method
func (m *MockIRoomBlacklistUseCase) Delete(c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockIRoomBlacklistUseCaseMockRecorder) Delete(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRoomBlacklistUseCase)(nil).Delete), c)
}
