// Code generated by MockGen. DO NOT EDIT.
// Source: game_hard_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	models "github.com/taniwhy/mochi-match-rest/domain/models"
	reflect "reflect"
)

// MockIGameHardUseCase is a mock of IGameHardUseCase interface
type MockIGameHardUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIGameHardUseCaseMockRecorder
}

// MockIGameHardUseCaseMockRecorder is the mock recorder for MockIGameHardUseCase
type MockIGameHardUseCaseMockRecorder struct {
	mock *MockIGameHardUseCase
}

// NewMockIGameHardUseCase creates a new mock instance
func NewMockIGameHardUseCase(ctrl *gomock.Controller) *MockIGameHardUseCase {
	mock := &MockIGameHardUseCase{ctrl: ctrl}
	mock.recorder = &MockIGameHardUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIGameHardUseCase) EXPECT() *MockIGameHardUseCaseMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockIGameHardUseCase) FindAll(c *gin.Context) ([]*models.GameHard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", c)
	ret0, _ := ret[0].([]*models.GameHard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockIGameHardUseCaseMockRecorder) FindAll(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIGameHardUseCase)(nil).FindAll), c)
}

// Insert mocks base method
func (m *MockIGameHardUseCase) Insert(c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockIGameHardUseCaseMockRecorder) Insert(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIGameHardUseCase)(nil).Insert), c)
}

// Update mocks base method
func (m *MockIGameHardUseCase) Update(c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockIGameHardUseCaseMockRecorder) Update(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIGameHardUseCase)(nil).Update), c)
}

// Delete mocks base method
func (m *MockIGameHardUseCase) Delete(c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockIGameHardUseCaseMockRecorder) Delete(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIGameHardUseCase)(nil).Delete), c)
}
