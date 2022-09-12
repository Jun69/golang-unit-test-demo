package main

import (
	"context"
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	pb "grpc_demo/helloworld/helloworld"
	"log"
	"net"
	"testing"
	"time"
)

type mockGreeterServer struct {
	pb.UnimplementedGreeterServer
}

// 自己实现一个简单的服务端
func (*mockGreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	if len(req.Name) <= 2 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
	}
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

// 构造一个闭包环境
func dialer(mockServer pb.GreeterServer) func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterGreeterServer(server, mockServer)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
func TestGreeterClient_SayHello_ByMock(t *testing.T) {
	tests := []struct {
		name string
		//response     *intel.IntelligenceResponse;
		message string
		res     string
	}{
		{
			"正常请求",
			"Boby",
			"Hello Boby",
		},
		{
			"非法参数",
			"d",
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(&mockGreeterServer{})))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convey.Convey(tt.name, t, func() {
				res, _ := NewGreeterClient(conn, time.Second).SayHello(context.Background(), tt.message)
				convey.So(res, should.Equal, tt.res)
				//convey.So(res, should.Resemble,tt.resp)
				//todo 1、为什么直接比较会卡死，下面的比较缺不会卡；2、 为什么 reflect.DeepEqual(res,tt.resp)也不相等
			})
		})
	}
}
