// Code generated by MockGen. DO NOT EDIT.
// Source: room_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/taniwhy/mochi-match-rest/domain/models"
	output "github.com/taniwhy/mochi-match-rest/domain/models/output"
	reflect "reflect"
)

// MockIRoomRepository is a mock of IRoomRepository interface
type MockIRoomRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRoomRepositoryMockRecorder
}

// MockIRoomRepositoryMockRecorder is the mock recorder for MockIRoomRepository
type MockIRoomRepositoryMockRecorder struct {
	mock *MockIRoomRepository
}

// NewMockIRoomRepository creates a new mock instance
func NewMockIRoomRepository(ctrl *gomock.Controller) *MockIRoomRepository {
	mock := &MockIRoomRepository{ctrl: ctrl}
	mock.recorder = &MockIRoomRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRoomRepository) EXPECT() *MockIRoomRepositoryMockRecorder {
	return m.recorder
}

// FindList mocks base method
func (m *MockIRoomRepository) FindList() ([]*output.RoomResBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList")
	ret0, _ := ret[0].([]*output.RoomResBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList
func (mr *MockIRoomRepositoryMockRecorder) FindList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockIRoomRepository)(nil).FindList))
}

// FindByLimitAndOffset mocks base method
func (m *MockIRoomRepository) FindByLimitAndOffset(limit, offset int) ([]*output.RoomResBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLimitAndOffset", limit, offset)
	ret0, _ := ret[0].([]*output.RoomResBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLimitAndOffset indicates an expected call of FindByLimitAndOffset
func (mr *MockIRoomRepositoryMockRecorder) FindByLimitAndOffset(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLimitAndOffset", reflect.TypeOf((*MockIRoomRepository)(nil).FindByLimitAndOffset), limit, offset)
}

// FindByLimitAndOffsetAndTitleAndHard mocks base method
func (m *MockIRoomRepository) FindByLimitAndOffsetAndTitleAndHard(limit, offset int, titles, hards []string) ([]*output.RoomResBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLimitAndOffsetAndTitleAndHard", limit, offset, titles, hards)
	ret0, _ := ret[0].([]*output.RoomResBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLimitAndOffsetAndTitleAndHard indicates an expected call of FindByLimitAndOffsetAndTitleAndHard
func (mr *MockIRoomRepositoryMockRecorder) FindByLimitAndOffsetAndTitleAndHard(limit, offset, titles, hards interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLimitAndOffsetAndTitleAndHard", reflect.TypeOf((*MockIRoomRepository)(nil).FindByLimitAndOffsetAndTitleAndHard), limit, offset, titles, hards)
}

// FindByID mocks base method
func (m *MockIRoomRepository) FindByID(roomID string) (*output.RoomResBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", roomID)
	ret0, _ := ret[0].(*output.RoomResBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockIRoomRepositoryMockRecorder) FindByID(roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockIRoomRepository)(nil).FindByID), roomID)
}

// FindByUserID mocks base method
func (m *MockIRoomRepository) FindByUserID(userID string) ([]*models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", userID)
	ret0, _ := ret[0].([]*models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID
func (mr *MockIRoomRepositoryMockRecorder) FindByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockIRoomRepository)(nil).FindByUserID), userID)
}

// FindUnlockByID mocks base method
func (m *MockIRoomRepository) FindUnlockByID(roomID string) (*models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUnlockByID", roomID)
	ret0, _ := ret[0].(*models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUnlockByID indicates an expected call of FindUnlockByID
func (mr *MockIRoomRepositoryMockRecorder) FindUnlockByID(roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUnlockByID", reflect.TypeOf((*MockIRoomRepository)(nil).FindUnlockByID), roomID)
}

// FindUnlockCountByID mocks base method
func (m *MockIRoomRepository) FindUnlockCountByID() (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUnlockCountByID")
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUnlockCountByID indicates an expected call of FindUnlockCountByID
func (mr *MockIRoomRepositoryMockRecorder) FindUnlockCountByID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUnlockCountByID", reflect.TypeOf((*MockIRoomRepository)(nil).FindUnlockCountByID))
}

// Insert mocks base method
func (m *MockIRoomRepository) Insert(room *models.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", room)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockIRoomRepositoryMockRecorder) Insert(room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIRoomRepository)(nil).Insert), room)
}

// Update mocks base method
func (m *MockIRoomRepository) Update(room *models.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", room)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockIRoomRepositoryMockRecorder) Update(room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRoomRepository)(nil).Update), room)
}

// Delete mocks base method
func (m *MockIRoomRepository) Delete(room *models.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", room)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockIRoomRepositoryMockRecorder) Delete(room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRoomRepository)(nil).Delete), room)
}

// LockFlg mocks base method
func (m *MockIRoomRepository) LockFlg(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockFlg", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockFlg indicates an expected call of LockFlg
func (mr *MockIRoomRepositoryMockRecorder) LockFlg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockFlg", reflect.TypeOf((*MockIRoomRepository)(nil).LockFlg), arg0, arg1)
}
