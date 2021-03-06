// Code generated by MockGen. DO NOT EDIT.
// Source: entry_history_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/taniwhy/mochi-match-rest/domain/models"
	output "github.com/taniwhy/mochi-match-rest/domain/models/output"
	reflect "reflect"
)

// MockIEntryHistoryRepository is a mock of IEntryHistoryRepository interface
type MockIEntryHistoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIEntryHistoryRepositoryMockRecorder
}

// MockIEntryHistoryRepositoryMockRecorder is the mock recorder for MockIEntryHistoryRepository
type MockIEntryHistoryRepositoryMockRecorder struct {
	mock *MockIEntryHistoryRepository
}

// NewMockIEntryHistoryRepository creates a new mock instance
func NewMockIEntryHistoryRepository(ctrl *gomock.Controller) *MockIEntryHistoryRepository {
	mock := &MockIEntryHistoryRepository{ctrl: ctrl}
	mock.recorder = &MockIEntryHistoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEntryHistoryRepository) EXPECT() *MockIEntryHistoryRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockIEntryHistoryRepository) FindAll() ([]*models.EntryHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*models.EntryHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockIEntryHistoryRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIEntryHistoryRepository)(nil).FindAll))
}

// FindNotLeave mocks base method
func (m *MockIEntryHistoryRepository) FindNotLeave(userID string) (*models.EntryHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindNotLeave", userID)
	ret0, _ := ret[0].(*models.EntryHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindNotLeave indicates an expected call of FindNotLeave
func (mr *MockIEntryHistoryRepositoryMockRecorder) FindNotLeave(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindNotLeave", reflect.TypeOf((*MockIEntryHistoryRepository)(nil).FindNotLeave), userID)
}

// FindNotLeaveByRoomID mocks base method
func (m *MockIEntryHistoryRepository) FindNotLeaveByRoomID(userID, roomID string) (*models.EntryHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindNotLeaveByRoomID", userID, roomID)
	ret0, _ := ret[0].(*models.EntryHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindNotLeaveByRoomID indicates an expected call of FindNotLeaveByRoomID
func (mr *MockIEntryHistoryRepositoryMockRecorder) FindNotLeaveByRoomID(userID, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindNotLeaveByRoomID", reflect.TypeOf((*MockIEntryHistoryRepository)(nil).FindNotLeaveByRoomID), userID, roomID)
}

// FindListByRoomID mocks base method
func (m *MockIEntryHistoryRepository) FindListByRoomID(roomID string) ([]*output.JoinUserRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByRoomID", roomID)
	ret0, _ := ret[0].([]*output.JoinUserRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByRoomID indicates an expected call of FindListByRoomID
func (mr *MockIEntryHistoryRepositoryMockRecorder) FindListByRoomID(roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByRoomID", reflect.TypeOf((*MockIEntryHistoryRepository)(nil).FindListByRoomID), roomID)
}

// FindListByUserID mocks base method
func (m *MockIEntryHistoryRepository) FindListByUserID(userID string) ([]*models.EntryHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByUserID", userID)
	ret0, _ := ret[0].([]*models.EntryHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByUserID indicates an expected call of FindListByUserID
func (mr *MockIEntryHistoryRepositoryMockRecorder) FindListByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByUserID", reflect.TypeOf((*MockIEntryHistoryRepository)(nil).FindListByUserID), userID)
}

// FindNotLeaveListByRoomID mocks base method
func (m *MockIEntryHistoryRepository) FindNotLeaveListByRoomID(roomID string) ([]*output.JoinUserRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindNotLeaveListByRoomID", roomID)
	ret0, _ := ret[0].([]*output.JoinUserRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindNotLeaveListByRoomID indicates an expected call of FindNotLeaveListByRoomID
func (mr *MockIEntryHistoryRepositoryMockRecorder) FindNotLeaveListByRoomID(roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindNotLeaveListByRoomID", reflect.TypeOf((*MockIEntryHistoryRepository)(nil).FindNotLeaveListByRoomID), roomID)
}

// Insert mocks base method
func (m *MockIEntryHistoryRepository) Insert(arg0 *models.EntryHistory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockIEntryHistoryRepositoryMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIEntryHistoryRepository)(nil).Insert), arg0)
}

// Update mocks base method
func (m *MockIEntryHistoryRepository) Update(arg0 *models.EntryHistory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockIEntryHistoryRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIEntryHistoryRepository)(nil).Update), arg0)
}

// Delete mocks base method
func (m *MockIEntryHistoryRepository) Delete(arg0 *models.EntryHistory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockIEntryHistoryRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIEntryHistoryRepository)(nil).Delete), arg0)
}

// CountEntryUser mocks base method
func (m *MockIEntryHistoryRepository) CountEntryUser(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountEntryUser", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountEntryUser indicates an expected call of CountEntryUser
func (mr *MockIEntryHistoryRepositoryMockRecorder) CountEntryUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountEntryUser", reflect.TypeOf((*MockIEntryHistoryRepository)(nil).CountEntryUser), arg0)
}

// LeaveFlg mocks base method
func (m *MockIEntryHistoryRepository) LeaveFlg(rid, uid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaveFlg", rid, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// LeaveFlg indicates an expected call of LeaveFlg
func (mr *MockIEntryHistoryRepositoryMockRecorder) LeaveFlg(rid, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveFlg", reflect.TypeOf((*MockIEntryHistoryRepository)(nil).LeaveFlg), rid, uid)
}
