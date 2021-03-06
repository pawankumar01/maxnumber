// Code generated by MockGen. DO NOT EDIT.
// Source: maxnumber/maxnumber (interfaces: NumberClient,Number_SendNumberClient)

// Package mock_maxnumber is a generated GoMock package.
package mock_maxnumber

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	maxnumber "maxnumber/maxnumber"
	reflect "reflect"
)

// MockNumberClient is a mock of NumberClient interface
type MockNumberClient struct {
	ctrl     *gomock.Controller
	recorder *MockNumberClientMockRecorder
}

// MockNumberClientMockRecorder is the mock recorder for MockNumberClient
type MockNumberClientMockRecorder struct {
	mock *MockNumberClient
}

// NewMockNumberClient creates a new mock instance
func NewMockNumberClient(ctrl *gomock.Controller) *MockNumberClient {
	mock := &MockNumberClient{ctrl: ctrl}
	mock.recorder = &MockNumberClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNumberClient) EXPECT() *MockNumberClientMockRecorder {
	return m.recorder
}

// SendNumber mocks base method
func (m *MockNumberClient) SendNumber(arg0 context.Context, arg1 ...grpc.CallOption) (maxnumber.Number_SendNumberClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendNumber", varargs...)
	ret0, _ := ret[0].(maxnumber.Number_SendNumberClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendNumber indicates an expected call of SendNumber
func (mr *MockNumberClientMockRecorder) SendNumber(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendNumber", reflect.TypeOf((*MockNumberClient)(nil).SendNumber), varargs...)
}

// MockNumber_SendNumberClient is a mock of Number_SendNumberClient interface
type MockNumber_SendNumberClient struct {
	ctrl     *gomock.Controller
	recorder *MockNumber_SendNumberClientMockRecorder
}

// MockNumber_SendNumberClientMockRecorder is the mock recorder for MockNumber_SendNumberClient
type MockNumber_SendNumberClientMockRecorder struct {
	mock *MockNumber_SendNumberClient
}

// NewMockNumber_SendNumberClient creates a new mock instance
func NewMockNumber_SendNumberClient(ctrl *gomock.Controller) *MockNumber_SendNumberClient {
	mock := &MockNumber_SendNumberClient{ctrl: ctrl}
	mock.recorder = &MockNumber_SendNumberClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNumber_SendNumberClient) EXPECT() *MockNumber_SendNumberClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockNumber_SendNumberClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockNumber_SendNumberClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockNumber_SendNumberClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockNumber_SendNumberClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockNumber_SendNumberClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNumber_SendNumberClient)(nil).Context))
}

// Header mocks base method
func (m *MockNumber_SendNumberClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockNumber_SendNumberClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockNumber_SendNumberClient)(nil).Header))
}

// Recv mocks base method
func (m *MockNumber_SendNumberClient) Recv() (*maxnumber.NumberMesg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*maxnumber.NumberMesg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockNumber_SendNumberClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockNumber_SendNumberClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockNumber_SendNumberClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockNumber_SendNumberClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNumber_SendNumberClient)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockNumber_SendNumberClient) Send(arg0 *maxnumber.NumberMesg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockNumber_SendNumberClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNumber_SendNumberClient)(nil).Send), arg0)
}

// SendMsg mocks base method
func (m *MockNumber_SendNumberClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockNumber_SendNumberClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNumber_SendNumberClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockNumber_SendNumberClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockNumber_SendNumberClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockNumber_SendNumberClient)(nil).Trailer))
}
