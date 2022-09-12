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
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
func TestSayHello(t *testing.T) {
	tests := []struct {
		name string
		//response     *intel.IntelligenceResponse;
		req  *pb.HelloRequest
		resp *pb.HelloReply
		err  error
	}{
		{
			"正常请求",
			&pb.HelloRequest{
				Name: "Boby",
			},
			&pb.HelloReply{Message: "Hello Boby"},
			nil,
		},
		{
			"非法参数",
			&pb.HelloRequest{
				Name: "d",
			},
			nil,
			status.Errorf(codes.InvalidArgument, "Invalid Argument"),
		},
	}
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	cli := pb.NewGreeterClient(conn)
	defer conn.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, e := cli.SayHello(ctx, tt.req)
			convey.Convey(tt.name, t, func() {
				convey.So(res.GetMessage(), should.Equal, tt.resp.GetMessage())
				//convey.So(res, should.Resemble,tt.resp)
				//todo 1、为什么直接比较会卡死，下面的比较缺不会卡；2、 为什么 reflect.DeepEqual(res,tt.resp)也不相等
				convey.So(e, should.Resemble, tt.err)
			})
		})
	}
}
