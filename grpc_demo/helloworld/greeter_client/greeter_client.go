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

// Package main implements a Client for Greeter service.
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"time"

	"google.golang.org/grpc"
	pb "grpc_demo/helloworld/helloworld"
)

type GreeterClient struct {
	conn    *grpc.ClientConn
	timeout time.Duration
	Client  pb.GreeterClient
}

func NewGreeterClient(conn *grpc.ClientConn, timeout time.Duration) *GreeterClient {
	return &GreeterClient{
		conn:    conn,
		timeout: timeout,
		Client:  pb.NewGreeterClient(conn),
	}
}

func (d *GreeterClient) SayHello(ctx context.Context, message string) (string, error) {
	request := &pb.HelloRequest{Name: message}
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(d.timeout))
	defer cancel()

	response, err := d.Client.SayHello(ctx, request)
	if err != nil {
		if er, ok := status.FromError(err); ok {
			return "", fmt.Errorf("grpc: %s, %s", er.Code(), er.Message())
		}
		return "", fmt.Errorf("server: %s", err.Error())
	}
	return response.GetMessage(), nil
}
