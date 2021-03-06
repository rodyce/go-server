// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rauljordan/go-server/internal/db (interfaces: Database)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "github.com/rauljordan/go-server/models"
	reflect "reflect"
)

// MockDatabase is a mock of Database interface
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockDatabase) CreateUser(arg0 context.Context, arg1 string, arg2 []byte) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockDatabaseMockRecorder) CreateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDatabase)(nil).CreateUser), arg0, arg1, arg2)
}

// User mocks base method
func (m *MockDatabase) User(arg0 context.Context, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// User indicates an expected call of User
func (mr *MockDatabaseMockRecorder) User(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockDatabase)(nil).User), arg0, arg1)
}
