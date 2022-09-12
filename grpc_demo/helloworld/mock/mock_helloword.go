// Code generated by MockGen. DO NOT EDIT.
// Source: helloworld/helloworld_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	helloworld "grpc_demo/helloworld/helloworld"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockGreeterClient is a mock of GreeterClient interface.
type MockGreeterClient struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterClientMockRecorder
}

// MockGreeterClientMockRecorder is the mock recorder for MockGreeterClient.
type MockGreeterClientMockRecorder struct {
	mock *MockGreeterClient
}

// NewMockGreeterClient creates a new mock instance.
func NewMockGreeterClient(ctrl *gomock.Controller) *MockGreeterClient {
	mock := &MockGreeterClient{ctrl: ctrl}
	mock.recorder = &MockGreeterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGreeterClient) EXPECT() *MockGreeterClientMockRecorder {
	return m.recorder
}

// SayHello mocks base method.
func (m *MockGreeterClient) SayHello(ctx context.Context, in *helloworld.HelloRequest, opts ...grpc.CallOption) (*helloworld.HelloReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SayHello", varargs...)
	ret0, _ := ret[0].(*helloworld.HelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SayHello indicates an expected call of SayHello.
func (mr *MockGreeterClientMockRecorder) SayHello(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockGreeterClient)(nil).SayHello), varargs...)
}

// MockGreeterServer is a mock of GreeterServer interface.
type MockGreeterServer struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterServerMockRecorder
}

// MockGreeterServerMockRecorder is the mock recorder for MockGreeterServer.
type MockGreeterServerMockRecorder struct {
	mock *MockGreeterServer
}

// NewMockGreeterServer creates a new mock instance.
func NewMockGreeterServer(ctrl *gomock.Controller) *MockGreeterServer {
	mock := &MockGreeterServer{ctrl: ctrl}
	mock.recorder = &MockGreeterServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGreeterServer) EXPECT() *MockGreeterServerMockRecorder {
	return m.recorder
}

// SayHello mocks base method.
func (m *MockGreeterServer) SayHello(arg0 context.Context, arg1 *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SayHello", arg0, arg1)
	ret0, _ := ret[0].(*helloworld.HelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SayHello indicates an expected call of SayHello.
func (mr *MockGreeterServerMockRecorder) SayHello(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockGreeterServer)(nil).SayHello), arg0, arg1)
}

// mustEmbedUnimplementedGreeterServer mocks base method.
func (m *MockGreeterServer) mustEmbedUnimplementedGreeterServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedGreeterServer")
}

// mustEmbedUnimplementedGreeterServer indicates an expected call of mustEmbedUnimplementedGreeterServer.
func (mr *MockGreeterServerMockRecorder) mustEmbedUnimplementedGreeterServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedGreeterServer", reflect.TypeOf((*MockGreeterServer)(nil).mustEmbedUnimplementedGreeterServer))
}

// MockUnsafeGreeterServer is a mock of UnsafeGreeterServer interface.
type MockUnsafeGreeterServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeGreeterServerMockRecorder
}

// MockUnsafeGreeterServerMockRecorder is the mock recorder for MockUnsafeGreeterServer.
type MockUnsafeGreeterServerMockRecorder struct {
	mock *MockUnsafeGreeterServer
}

// NewMockUnsafeGreeterServer creates a new mock instance.
func NewMockUnsafeGreeterServer(ctrl *gomock.Controller) *MockUnsafeGreeterServer {
	mock := &MockUnsafeGreeterServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeGreeterServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeGreeterServer) EXPECT() *MockUnsafeGreeterServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedGreeterServer mocks base method.
func (m *MockUnsafeGreeterServer) mustEmbedUnimplementedGreeterServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedGreeterServer")
}

// mustEmbedUnimplementedGreeterServer indicates an expected call of mustEmbedUnimplementedGreeterServer.
func (mr *MockUnsafeGreeterServerMockRecorder) mustEmbedUnimplementedGreeterServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedGreeterServer", reflect.TypeOf((*MockUnsafeGreeterServer)(nil).mustEmbedUnimplementedGreeterServer))
}
