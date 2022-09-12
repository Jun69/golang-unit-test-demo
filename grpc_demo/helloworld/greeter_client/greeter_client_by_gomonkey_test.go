package main

import (
	"context"
	"fmt"
	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
	"time"
)

func TestGreeterClient_SayHello_ByGomonkey(t *testing.T) {
	cli := NewGreeterClient(nil, time.Minute)
	Convey("TestSayHello", t, func() {

		Convey("请求成功", func() {
			gomonkey.ApplyMethod(reflect.TypeOf(cli), "SayHello", func(_ *GreeterClient, _ context.Context, message string) (string, error) {
				return "hello bob", nil
			})
			res, err1 := cli.SayHello(context.Background(), "bob")
			So(res, should.Equal, "hello bob")
			So(err1, should.BeEmpty)
		})
		Convey("请求失败", func() {
			gomonkey.ApplyMethod(reflect.TypeOf(cli), "SayHello", func(_ *GreeterClient, _ context.Context, message string) (string, error) {
				return "", fmt.Errorf("error")
			})
			res, err1 := cli.SayHello(context.Background(), "b")
			So(res, should.Equal, "")
			So(err1, ShouldBeError)
		})
	})
}
