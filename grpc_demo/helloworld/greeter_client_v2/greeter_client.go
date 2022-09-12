/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	pb "grpc_demo/helloworld/helloworld"
	"time"
)

type GreeterClient struct {
	conn    *grpc.ClientConn
	timeout time.Duration
	Client  pb.GreeterClient
}

func NewGreeterClient() (*GreeterClient, error) {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defaultTimeout := time.Minute
	return &GreeterClient{
		conn:    conn,
		timeout: defaultTimeout,
		Client:  pb.NewGreeterClient(conn),
	}, nil
}

// GetRespWithRpcCli 必须采用这种依赖注入才能使用gomock
func GetRespWithRpcCli(client pb.GreeterClient, message string) (string, error) {
	res, err := client.SayHello(context.TODO(), &pb.HelloRequest{Name: message})
	if err != nil {
		if er, ok := status.FromError(err); ok {
			return "", fmt.Errorf("grpc: %s, %s", er.Code(), er.Message())
		}
		return "", fmt.Errorf("server: %s", err.Error())
	}
	return res.GetMessage(), nil
}
