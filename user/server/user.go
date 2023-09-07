package main

import (
	pb "GolandProjects/grpc-tutorial/user/proto"
	"context"
	"log"
)

func (s *Server) User(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("User function was invoked with %v\n", in)
	return &pb.UserResponse{
		Name: in.Name,
		Age:  in.Age,
	}, nil
}
