// Code generated by MockGen. DO NOT EDIT.
// Source: report_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/taniwhy/mochi-match-rest/domain/models"
	reflect "reflect"
)

// MockIReportRepository is a mock of IReportRepository interface
type MockIReportRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIReportRepositoryMockRecorder
}

// MockIReportRepositoryMockRecorder is the mock recorder for MockIReportRepository
type MockIReportRepositoryMockRecorder struct {
	mock *MockIReportRepository
}

// NewMockIReportRepository creates a new mock instance
func NewMockIReportRepository(ctrl *gomock.Controller) *MockIReportRepository {
	mock := &MockIReportRepository{ctrl: ctrl}
	mock.recorder = &MockIReportRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIReportRepository) EXPECT() *MockIReportRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockIReportRepository) FindAll() ([]*models.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*models.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockIReportRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIReportRepository)(nil).FindAll))
}

// Insert mocks base method
func (m *MockIReportRepository) Insert(report *models.Report) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", report)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockIReportRepositoryMockRecorder) Insert(report interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIReportRepository)(nil).Insert), report)
}

// Delete mocks base method
func (m *MockIReportRepository) Delete(report *models.Report) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", report)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockIReportRepositoryMockRecorder) Delete(report interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIReportRepository)(nil).Delete), report)
}
