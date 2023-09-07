package main

import (
	pb "GolandProjects/grpc-tutorial/user/proto"
	"context"
	"log"
)

func doUser(c pb.UserServiceClient) {
	log.Println("doUser was invoked")
	res, err := c.User(context.Background(), &pb.UserRequest{
		Name: "Gonzalo",
		Age:  27,
	})

	if err != nil {
		log.Fatalf("Could not user: %v\n", err)
	}

	log.Printf("User:\nName = %s\nAge = %d\n", res.Name, res.Age)
}
