// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tsuna/gohbase (interfaces: Client)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	hrpc "github.com/tsuna/gohbase/hrpc"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Append mocks base method
func (m *MockClient) Append(arg0 *hrpc.Mutate) (*hrpc.Result, error) {
	ret := m.ctrl.Call(m, "Append", arg0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Append indicates an expected call of Append
func (mr *MockClientMockRecorder) Append(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockClient)(nil).Append), arg0)
}

// CheckAndPut mocks base method
func (m *MockClient) CheckAndPut(arg0 *hrpc.Mutate, arg1, arg2 string, arg3 []byte) (bool, error) {
	ret := m.ctrl.Call(m, "CheckAndPut", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAndPut indicates an expected call of CheckAndPut
func (mr *MockClientMockRecorder) CheckAndPut(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAndPut", reflect.TypeOf((*MockClient)(nil).CheckAndPut), arg0, arg1, arg2, arg3)
}

// Close mocks base method
func (m *MockClient) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// Delete mocks base method
func (m *MockClient) Delete(arg0 *hrpc.Mutate) (*hrpc.Result, error) {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockClientMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), arg0)
}

// Get mocks base method
func (m *MockClient) Get(arg0 *hrpc.Get) (*hrpc.Result, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClientMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), arg0)
}

// Increment mocks base method
func (m *MockClient) Increment(arg0 *hrpc.Mutate) (int64, error) {
	ret := m.ctrl.Call(m, "Increment", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Increment indicates an expected call of Increment
func (mr *MockClientMockRecorder) Increment(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Increment", reflect.TypeOf((*MockClient)(nil).Increment), arg0)
}

// Put mocks base method
func (m *MockClient) Put(arg0 *hrpc.Mutate) (*hrpc.Result, error) {
	ret := m.ctrl.Call(m, "Put", arg0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put
func (mr *MockClientMockRecorder) Put(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockClient)(nil).Put), arg0)
}

// Scan mocks base method
func (m *MockClient) Scan(arg0 *hrpc.Scan) hrpc.Scanner {
	ret := m.ctrl.Call(m, "Scan", arg0)
	ret0, _ := ret[0].(hrpc.Scanner)
	return ret0
}

// Scan indicates an expected call of Scan
func (mr *MockClientMockRecorder) Scan(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockClient)(nil).Scan), arg0)
}
