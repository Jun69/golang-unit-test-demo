package main

import "log"

func main() {
	client, err := NewGreeterClient()
	if err != nil {
		log.Println(err.Error())
		return
	}
	res, err := GetRespWithRpcCli(client.Client, "boby")
	if err != nil {
		log.Println(err.Error())
		return
	}
	println(res)
}
