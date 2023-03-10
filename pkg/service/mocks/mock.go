// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	entity "github.com/4halavandala/l0/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockTodoOrder is a mock of TodoOrder interface.
type MockTodoOrder struct {
	ctrl     *gomock.Controller
	recorder *MockTodoOrderMockRecorder
}

// MockTodoOrderMockRecorder is the mock recorder for MockTodoOrder.
type MockTodoOrderMockRecorder struct {
	mock *MockTodoOrder
}

// NewMockTodoOrder creates a new mock instance.
func NewMockTodoOrder(ctrl *gomock.Controller) *MockTodoOrder {
	mock := &MockTodoOrder{ctrl: ctrl}
	mock.recorder = &MockTodoOrderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoOrder) EXPECT() *MockTodoOrderMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodoOrder) Create(order entity.Model) (*entity.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", order)
	ret0, _ := ret[0].(*entity.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTodoOrderMockRecorder) Create(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoOrder)(nil).Create), order)
}

// GetAll mocks base method.
func (m *MockTodoOrder) GetAll() (*[]entity.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*[]entity.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTodoOrderMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTodoOrder)(nil).GetAll))
}

// GetAllFromCache mocks base method.
func (m *MockTodoOrder) GetAllFromCache() (*[]entity.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFromCache")
	ret0, _ := ret[0].(*[]entity.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFromCache indicates an expected call of GetAllFromCache.
func (mr *MockTodoOrderMockRecorder) GetAllFromCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFromCache", reflect.TypeOf((*MockTodoOrder)(nil).GetAllFromCache))
}

// GetById mocks base method.
func (m *MockTodoOrder) GetById(orderId string) (*entity.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", orderId)
	ret0, _ := ret[0].(*entity.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockTodoOrderMockRecorder) GetById(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockTodoOrder)(nil).GetById), orderId)
}

// GetByIdFromCache mocks base method.
func (m *MockTodoOrder) GetByIdFromCache(orderId string) (*entity.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIdFromCache", orderId)
	ret0, _ := ret[0].(*entity.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIdFromCache indicates an expected call of GetByIdFromCache.
func (mr *MockTodoOrderMockRecorder) GetByIdFromCache(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIdFromCache", reflect.TypeOf((*MockTodoOrder)(nil).GetByIdFromCache), orderId)
}

// InitCache mocks base method.
func (m *MockTodoOrder) InitCache() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitCache")
	ret0, _ := ret[0].(error)
	return ret0
}

// InitCache indicates an expected call of InitCache.
func (mr *MockTodoOrderMockRecorder) InitCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitCache", reflect.TypeOf((*MockTodoOrder)(nil).InitCache))
}
