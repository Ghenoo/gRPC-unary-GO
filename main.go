package main

import (
	"golang-grpc-unary.com/example/client"
	"golang-grpc-unary.com/example/server"
)

func main() {
	go server.Run()
	client.Run()
}
