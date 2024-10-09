package client

import (
	"context"
	"fmt"

	"golang-grpc-unary.com/example/pb"
	"google.golang.org/grpc"
)

func Run() {
	dial, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {panic(err)}

	defer dial.Close()

	userClient := pb.NewUserServiceClient(dial)

	user, err := userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id: "1",
		Name: "ghenoninho",
		Age: 20,
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("First user createad: %v\n", user)

	user, err = userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id: "2",
		Name: "epaminondas",
		Age: 20,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Second user createad: %v\n", user)

	getUserResponse, err := userClient.GetUser(context.Background(), &pb.GetUserRequest{
		Id: "1"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get user: %v\n", getUserResponse)
}