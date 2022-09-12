package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	log.Println("Client running ...")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	response, err := NewGreeterClient(conn, time.Second).SayHello(context.Background(), "bob")
	log.Println(response)
	log.Println(err)
}
