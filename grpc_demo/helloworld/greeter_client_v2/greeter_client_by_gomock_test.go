package main

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	pb "grpc_demo/helloworld/helloworld"
	"grpc_demo/helloworld/mock"
	"testing"
)

func TestGetRespWithRpcCli(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock.NewMockGreeterClient(ctrl)
	gomock.InOrder(
		// 指定每次的替换值
		mockClient.EXPECT().SayHello(gomock.Any(), &pb.HelloRequest{Name: "boby"}).Return(&pb.HelloReply{Message: "Hello boby"}, nil),
		mockClient.EXPECT().SayHello(gomock.Any(), &pb.HelloRequest{Name: "b"}).Return(nil, fmt.Errorf("error")),
	)

	tests := []struct {
		name    string
		message string
		res     string
	}{
		{
			"正常请求",
			"boby",
			"Hello boby",
		},
		{
			"非法参数",
			"b",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convey.Convey(tt.name, t, func() {
				res, _ := GetRespWithRpcCli(mockClient, tt.message)
				convey.So(res, should.Equal, tt.res)
			})
		})
	}
}
