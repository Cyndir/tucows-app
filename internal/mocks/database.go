// Code generated by MockGen. DO NOT EDIT.
// Source: database.go
//
// Generated by this command:
//
//	mockgen -destination=../mocks/database.go -package=mocks -source=database.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/Cyndir/tucows-app/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockDatabase) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockDatabaseMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockDatabase)(nil).Connect))
}

// Disconnect mocks base method.
func (m *MockDatabase) Disconnect() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disconnect")
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockDatabaseMockRecorder) Disconnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockDatabase)(nil).Disconnect))
}

// GetOrder mocks base method.
func (m *MockDatabase) GetOrder(id string) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", id)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockDatabaseMockRecorder) GetOrder(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockDatabase)(nil).GetOrder), id)
}

// InsertOrder mocks base method.
func (m *MockDatabase) InsertOrder(order *model.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOrder", order)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOrder indicates an expected call of InsertOrder.
func (mr *MockDatabaseMockRecorder) InsertOrder(order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOrder", reflect.TypeOf((*MockDatabase)(nil).InsertOrder), order)
}

// UpdateOrder mocks base method.
func (m *MockDatabase) UpdateOrder(orderID, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", orderID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockDatabaseMockRecorder) UpdateOrder(orderID, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockDatabase)(nil).UpdateOrder), orderID, status)
}
