// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tsuna/gohbase/hrpc (interfaces: Call)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
	hrpc "github.com/tsuna/gohbase/hrpc"
	reflect "reflect"
)

// MockCall is a mock of Call interface
type MockCall struct {
	ctrl     *gomock.Controller
	recorder *MockCallMockRecorder
}

// MockCallMockRecorder is the mock recorder for MockCall
type MockCallMockRecorder struct {
	mock *MockCall
}

// NewMockCall creates a new mock instance
func NewMockCall(ctrl *gomock.Controller) *MockCall {
	mock := &MockCall{ctrl: ctrl}
	mock.recorder = &MockCallMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCall) EXPECT() *MockCallMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockCall) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockCallMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockCall)(nil).Context))
}

// Key mocks base method
func (m *MockCall) Key() []byte {
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Key indicates an expected call of Key
func (mr *MockCallMockRecorder) Key() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockCall)(nil).Key))
}

// Name mocks base method
func (m *MockCall) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockCallMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockCall)(nil).Name))
}

// NewResponse mocks base method
func (m *MockCall) NewResponse() proto.Message {
	ret := m.ctrl.Call(m, "NewResponse")
	ret0, _ := ret[0].(proto.Message)
	return ret0
}

// NewResponse indicates an expected call of NewResponse
func (mr *MockCallMockRecorder) NewResponse() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewResponse", reflect.TypeOf((*MockCall)(nil).NewResponse))
}

// Region mocks base method
func (m *MockCall) Region() hrpc.RegionInfo {
	ret := m.ctrl.Call(m, "Region")
	ret0, _ := ret[0].(hrpc.RegionInfo)
	return ret0
}

// Region indicates an expected call of Region
func (mr *MockCallMockRecorder) Region() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Region", reflect.TypeOf((*MockCall)(nil).Region))
}

// ResultChan mocks base method
func (m *MockCall) ResultChan() chan hrpc.RPCResult {
	ret := m.ctrl.Call(m, "ResultChan")
	ret0, _ := ret[0].(chan hrpc.RPCResult)
	return ret0
}

// ResultChan indicates an expected call of ResultChan
func (mr *MockCallMockRecorder) ResultChan() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResultChan", reflect.TypeOf((*MockCall)(nil).ResultChan))
}

// SetRegion mocks base method
func (m *MockCall) SetRegion(arg0 hrpc.RegionInfo) {
	m.ctrl.Call(m, "SetRegion", arg0)
}

// SetRegion indicates an expected call of SetRegion
func (mr *MockCallMockRecorder) SetRegion(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRegion", reflect.TypeOf((*MockCall)(nil).SetRegion), arg0)
}

// Table mocks base method
func (m *MockCall) Table() []byte {
	ret := m.ctrl.Call(m, "Table")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Table indicates an expected call of Table
func (mr *MockCallMockRecorder) Table() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Table", reflect.TypeOf((*MockCall)(nil).Table))
}

// ToProto mocks base method
func (m *MockCall) ToProto() proto.Message {
	ret := m.ctrl.Call(m, "ToProto")
	ret0, _ := ret[0].(proto.Message)
	return ret0
}

// ToProto indicates an expected call of ToProto
func (mr *MockCallMockRecorder) ToProto() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToProto", reflect.TypeOf((*MockCall)(nil).ToProto))
}
